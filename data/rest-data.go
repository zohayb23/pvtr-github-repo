package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/google/go-github/v71/github"
	"github.com/ossf/si-tooling/v2/si"
	"github.com/privateerproj/privateer-sdk/config"
)

type RestData struct {
	owner               string
	repo                string
	token               string
	Config              *config.Config
	WorkflowPermissions WorkflowPermissions
	Insights            si.SecurityInsights
	Releases            []ReleaseData
	Rulesets            []Ruleset
	contents            RepoContent
	ghClient            *github.Client
}

type RepoContent struct {
	Content    []*github.RepositoryContent
	SubContent map[string]RepoContent
}

type Ruleset struct {
	Type       string `json:"type"`
	Parameters struct {
		RequiredChecks []struct {
			Context string `json:"context"`
		} `json:"required_status_checks"`
	} `json:"parameters"`
}

type ReleaseData struct {
	Id      int            `json:"id"`
	Name    string         `json:"name"`
	TagName string         `json:"tag_name"`
	URL     string         `json:"url"`
	Assets  []ReleaseAsset `json:"assets"`
}

type ReleaseAsset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
}

type WorkflowPermissions struct {
	DefaultPermissions    string `json:"default_workflow_permissions"`
	CanApprovePullRequest bool   `json:"can_approve_pull_request_reviews"`
}

var APIBase = "https://api.github.com"

func (r *RestData) Setup() error {
	r.owner = r.Config.GetString("owner")
	r.repo = r.Config.GetString("repo")
	r.token = r.Config.GetString("token")

	r.getRepoContents()
	r.loadSecurityInsights()
	_ = r.getWorkflowPermissions()
	_ = r.getReleases()
	return nil
}

func (r *RestData) MakeApiCall(endpoint string, isGithub bool) (body []byte, err error) {
	r.Config.Logger.Trace(fmt.Sprintf("GET %s", endpoint))
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	if isGithub {
		request.Header.Set("Authorization", "Bearer "+r.token)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("error making http call: %s", err.Error())
		return nil, err
	}
	if response.StatusCode != 200 {
		err = fmt.Errorf("unexpected response: %s", response.Status)
		return nil, err
	}
	return io.ReadAll(response.Body)
}

func (r *RestData) getSourceFile(owner, repo, path string) (content *github.RepositoryContent, err error) {
	content, _, _, err = r.ghClient.Repositories.GetContents(context.Background(), owner, repo, path, nil)
	if err != nil {
		return
	}
	return content, nil
}

// checkFile accepts a filename like security-insights.yml or security.md and returns the path to that file
// if it exists in the root directory or forge directory of the repository or returns "" when the file is not found
func (r *RestData) checkFile(filename string) (filepath string) {
	filepath = ""
	for _, dirContents := range r.contents.Content {
		// top level directory contents
		if strings.EqualFold(*dirContents.Name, filename) {
			filepath = *dirContents.Path
			break
		}
	}
	// prefer files found in the root directory
	if filepath != "" {
		return filepath
	}
	for _, dirContents := range r.contents.SubContent[".github"].Content {
		// forge directory contents
		if dirContents.GetType() != "file" {
			continue
		}
		if strings.EqualFold(*dirContents.Name, filename) {
			filepath = *dirContents.Path
			break
		}
	}
	return filepath
}

func (r *RestData) GetDirectoryContent(path string) (dirContent []*github.RepositoryContent, err error) {
	workflowsDir, err := r.contents.GetSubdirContentByPath(r, path)
	if err != nil {
		return nil, fmt.Errorf("content not found at %s: %w", path, err)
	}

	for _, file := range workflowsDir.Content {
		if file.GetType() != "file" {
			continue
		}

		content, err := r.getSourceFile(r.owner, r.repo, file.GetPath())
		if err != nil {
			return nil, fmt.Errorf("failed to fetch workflow file %s: %s", file.GetPath(), err.Error())
		}
		dirContent = append(dirContent, content)
	}

	return dirContent, nil
}

func (r *RestData) GetFileContent(path string) (content *github.RepositoryContent, err error) {
	content, err = r.getSourceFile(r.owner, r.repo, path)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve file content for %s: %w", path, err)
	}
	if content == nil {
		return nil, fmt.Errorf("file not found at %s", path)
	}
	return content, nil
}

// returns true when a file with case insensitive name matching support.md is found in the root or forge directories or when the readme.md contains a heading named "Support"
func (r *RestData) HasSupportMarkdown() bool {
	if r.checkFile("support.md") != "" {
		return true
	}
	readmePath := r.checkFile("readme.md")
	if readmePath != "" {
		contents, err := r.getSourceFile(r.owner, r.repo, readmePath)
		if err != nil {
			r.Config.Logger.Error(fmt.Sprintf("failed to retrieve readme file data: %s", err.Error()))
			return false
		}
		content, err := contents.GetContent()
		if err != nil {
			r.Config.Logger.Error(fmt.Sprintf("failed to unpack readme contents: %s", err.Error()))
			return false
		}
		headings := parseMarkdownHeadings([]byte(content))
		for _, heading := range headings {
			if heading == "Support" {
				return true
			}
		}
	}
	return false
}

func parseMarkdownHeadings(content []byte) []string {
	var headings []string

	// Parse markdown into AST
	md := markdown.Parse(content, nil)

	// Walk the AST and collect headings
	ast.WalkFunc(md, func(node ast.Node, entering bool) ast.WalkStatus {
		if heading, ok := node.(*ast.Heading); ok && entering {
			// Get the text content of the heading
			if len(heading.Children) > 0 {
				if text, ok := heading.Children[0].(*ast.Text); ok {
					headings = append(headings, string(text.Literal))
				}
			}
		}
		return ast.GoToNext
	})

	return headings
}

func (r *RestData) loadSecurityInsights() {
	filepath := r.checkFile(si.SecurityInsightsFilename)
	if filepath != "" {
		insights, err := si.Read(r.owner, r.repo, filepath)
		r.Insights = insights
		if err != nil {
			r.Config.Logger.Error(fmt.Sprintf("failed to read security insights file: %s", err.Error()))
		}
		return
	}
}

func (r *RestData) getRepoContents() {
	_, content, _, err := r.ghClient.Repositories.GetContents(context.Background(), r.owner, r.repo, "", nil)
	if err != nil {
		r.Config.Logger.Error(fmt.Sprintf("failed to retrieve contents top level contents: %s", err.Error()))
		return
	}
	r.contents.Content = content
	if len(r.contents.Content) == 0 {
		r.Config.Logger.Error("no contents found at the top level of the repository")
		return
	}
	r.contents.SubContent = make(map[string]RepoContent)
	r.Config.Logger.Trace(fmt.Sprintf("retrieved %d top-level contents", len(r.contents.Content)))
}

func (c *RepoContent) GetSubdirContentByPath(r *RestData, path string) (RepoContent, error) {
	if c.SubContent == nil {
		return RepoContent{}, fmt.Errorf("no subdirectories found")
	}

	parts := strings.Split(path, "/")
	current := *c
	currentPath := ""

	for i, part := range parts {
		// Build the current path for this level
		if currentPath == "" {
			currentPath = part
		} else {
			currentPath = currentPath + "/" + part
		}

		// Check if we already have this subdirectory's content
		subdir, exists := current.SubContent[part]
		if !exists {
			// Find this directory in the current level's content
			var dirEntry *github.RepositoryContent
			for _, entry := range current.Content {
				if entry.GetType() == "dir" && entry.GetName() == part {
					dirEntry = entry
					break
				}
			}

			if dirEntry == nil {
				return RepoContent{}, fmt.Errorf("directory '%s' not found in path '%s'", part, path)
			}

			// Fetch the contents of this directory
			var err error
			subdir, err = r.getSubdirContents(dirEntry.GetPath())
			if err != nil {
				return RepoContent{}, fmt.Errorf("failed to retrieve contents for %s: %w", dirEntry.GetPath(), err)
			}

			// Cache the result
			current.SubContent[part] = subdir
		}

		// Move to the next level
		current = subdir

		// If this is the last part and we got here, we found the directory
		if i == len(parts)-1 {
			return current, nil
		}
	}

	return current, nil
}

// getSubdirContents fetches contents of a directory
func (r *RestData) getSubdirContents(path string) (RepoContent, error) {
	_, content, _, err := r.ghClient.Repositories.GetContents(context.Background(), r.owner, r.repo, path, nil)
	if err != nil {
		return RepoContent{}, err
	}

	return RepoContent{
		Content:    content,
		SubContent: make(map[string]RepoContent),
	}, nil
}

func (r *RestData) getReleases() error {
	endpoint := fmt.Sprintf("%s/repos/%s/%s/releases", APIBase, r.owner, r.repo)
	responseData, err := r.MakeApiCall(endpoint, true)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, &r.Releases)
}

func (r *RestData) getWorkflowPermissions() error {
	endpoint := fmt.Sprintf("%s/repos/%s/%s/actions/permissions/workflow", APIBase, r.owner, r.repo)
	responseData, err := r.MakeApiCall(endpoint, true)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(responseData, &r.WorkflowPermissions); err != nil {
		return fmt.Errorf("failed to parse permissions: %v", err)
	}
	return err
}

func (r *RestData) GetRulesets(branchName string) []Ruleset {
	endpoint := fmt.Sprintf("%s/repos/%s/%s/rules/branches/%s", APIBase, r.owner, r.repo, branchName)
	responseData, err := r.MakeApiCall(endpoint, true)
	if err != nil {
		r.Config.Logger.Error(fmt.Sprintf("error getting rulesets: %s", err.Error()))
	}

	_ = json.Unmarshal(responseData, &r.Rulesets)
	return r.Rulesets
}

func (r *RestData) IsCodeRepo() (bool, error) {
	endpoint := fmt.Sprintf("%s/repos/%s/%s/languages", APIBase, r.owner, r.repo)
	responseData, err := r.MakeApiCall(endpoint, true)
	if err != nil {
		return true, err
	}
	languagesUsed := make(map[string]int)

	err = json.Unmarshal(responseData, &languagesUsed)
	if err != nil {
		return true, err
	}

	r.Config.Logger.Trace(fmt.Sprintf("%v Languages found in repo: %v", len(languagesUsed), languagesUsed))
	if len(languagesUsed) == 0 {
		return false, nil
	}

	return true, nil
}
