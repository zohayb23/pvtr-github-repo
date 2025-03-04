package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ossf/si-tooling/v2/si"
	"github.com/privateerproj/privateer-sdk/config"
)

type RestData struct {
	owner    string
	repo     string
	Metadata RepoMetadata
	Insights si.SecurityInsights
	Config   *config.Config
}

type RepoMetadata struct {
	Name     string `json:"name"`
	Private  bool   `json:"private"`
	Releases []ReleaseData
	Contents struct {
		TopLevel []DirContents
		ForgeDir []DirContents
	}
	// WorkflowPermissions WorkflowPermissions
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

type DirContents struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	SHA         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

type FileAPIResponse struct {
	ByteContent []byte `json:"content"`
	SHA         string `json:"sha"`
}

type WorkflowPermissions struct {
	DefaultWorkflowPermissions string `json:"default_workflow_permissions"`
	CanApprovePullRequest      bool   `json:"can_approve_pull_request_reviews"`
}

var APIBase = "https://api.github.com/repos"

func (r *RestData) Setup() error {
	r.owner = r.Config.GetString("owner")
	r.repo = r.Config.GetString("repo")

	r.getMetadata()
	r.loadSecurityInsights()
	return nil
}

func (r *RestData) makeApiCall(endpoint string) (body []byte, err error) {
	r.Config.Logger.Trace(fmt.Sprintf("GET %s", endpoint))
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+r.Config.GetString("token"))
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

func (r *RestData) getSourceFile(owner, repo, path string) (response FileAPIResponse, err error) {
	endpoint := fmt.Sprintf("%s/%s/%s/contents/%s", APIBase, owner, repo, path)
	responseData, err := r.makeApiCall(endpoint)
	if err != nil {
		return
	}
	err = json.Unmarshal(responseData, &response)
	return
}

func (r *RestData) loadSecurityInsights() {
	r.getTopDirContents()
	if len(r.Metadata.Contents.TopLevel) == 0 {
		r.Config.Logger.Error("no contents retrieved from the top level of the repository")
		return
	}
	for _, content := range r.Metadata.Contents.TopLevel {
		if r.foundSecurityInsights(content) {
			insights, err := si.Read(r.owner, r.repo, "security-insights.yml")
			r.Insights = insights
			if err != nil {
				r.Config.Logger.Error(fmt.Sprintf("error reading security insights file: %s", err.Error()))
			}
			return
		}
	}
	r.getForgeDirContents()
	for _, content := range r.Metadata.Contents.ForgeDir {
		if r.foundSecurityInsights(content) {
			insights, err := si.Read(r.owner, r.repo, ".github/security-insights.yml")
			r.Insights = insights
			if err != nil {
				r.Config.Logger.Error(fmt.Sprintf("error reading security insights file: %s", err.Error()))
			}
			return
		}
	}
	r.Config.Logger.Error("no security insights file found")
}

func (r *RestData) foundSecurityInsights(content DirContents) bool {
	if strings.Contains(strings.ToLower(content.Name), "security-insights.") {
		response, err := r.getSourceFile(r.owner, r.repo, content.Path)
		if err != nil {
			r.Config.Logger.Error(fmt.Sprintf("error unmarshalling API response for security insights file: %s", err.Error()))
			return false
		}
		r.Config.Logger.Trace(fmt.Sprintf("Security Insights Exists - SHA: %v", response.SHA))
		return true
	}
	return false
}

func (r *RestData) getTopDirContents() {
	endpoint := fmt.Sprintf("%s/%s/%s/contents", APIBase, r.owner, r.repo)
	responseData, err := r.makeApiCall(endpoint)
	if err != nil {
		r.Config.Logger.Error(fmt.Sprintf("error getting top level contents: %s", err.Error()))
		return
	}
	json.Unmarshal(responseData, &r.Metadata.Contents.TopLevel)
}

func (r *RestData) getForgeDirContents() {
	endpoint := fmt.Sprintf("%s/%s/%s/contents/.github", APIBase, r.owner, r.repo)
	responseData, err := r.makeApiCall(endpoint)
	if err != nil {
		r.Config.Logger.Error(fmt.Sprintf("error getting forge contents: %s", err.Error()))
		return
	}
	json.Unmarshal(responseData, &r.Metadata.Contents.ForgeDir)
}

func (r *RestData) getMetadata() error {
	endpoint := fmt.Sprintf("%s/%s/%s", APIBase, r.owner, r.repo)
	responseData, err := r.makeApiCall(endpoint)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, &r.Metadata)
}

// func (r *RepoMetadata) getReleases(owner, repo string) error {
// 	endpoint := fmt.Sprintf("%s/%s/%s/releases", APIBase, owner, repo)
// 	responseData, err := r.makeApiCall(endpoint)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal(responseData, &r.Releases)
// }

// func (r *RestData) getWorkflowPermissions() (WorkflowPermissions, error) {
// 	if r.Metadata.WorkflowPermissions != (WorkflowPermissions{}) {
// 		return r.Metadata.WorkflowPermissions, nil
// 	}

// 	endpoint := fmt.Sprintf("%s/%s/%s/actions/permissions/workflow", APIBase, r.owner, r.repo)
// 	responseData, err := r.makeApiCall(endpoint)
// 	if err != nil {
// 		return WorkflowPermissions{}, err
// 	}

// 	var permResp WorkflowPermissions
// 	if err := json.Unmarshal(responseData, &permResp); err != nil {
// 		return WorkflowPermissions{}, fmt.Errorf("failed to parse permissions: %v", err)
// 	}

// 	return permResp, nil
// }

func (r *RestData) getFileContentByURL(downloadURL string) (string, error) {
	responseData, err := r.makeApiCall(downloadURL)
	if err != nil {
		return "", err
	}
	return string(responseData), nil
}
