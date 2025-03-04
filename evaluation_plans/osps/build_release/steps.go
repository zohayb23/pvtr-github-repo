package build_release

import (
	"fmt"
	"strings"

	"github.com/revanite-io/sci/pkg/layer4"

	"github.com/revanite-io/pvtr-github-repo/data"
	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
)

func passIfNoRelases(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if len(data.Releases) == 0 {
		return layer4.PassAndHalt, "No releases found to evaluate"
	}

	return layer4.Passed, fmt.Sprintf("Found %v releases, continuing to evaluate", len(data.Releases))
}

func releaseHasUniqueIdentifier(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	var noNameCount int
	var sameNameFound []string
	var releaseNames = make(map[string]int)

	for _, release := range data.Releases {
		if release.Name == "" {
			noNameCount++
		} else if _, ok := releaseNames[release.Name]; ok {
			sameNameFound = append(sameNameFound, release.Name)
		} else {
			releaseNames[release.Name] = release.Id
		}
	}
	if noNameCount > 0 || len(sameNameFound) > 0 {
		sameNames := strings.Join(sameNameFound, ", ")
		message := []string{fmt.Sprintf("Found %v releases with no name", noNameCount)}
		if len(sameNameFound) > 0 {
			message = append(message, fmt.Sprintf("Found %v releases with the same name: %v", len(sameNameFound), sameNames))
		}
		return layer4.Failed, strings.Join(message, ". ")
	}
	return layer4.Passed, "All releases found have a unique name"
}

func getLinksFromInsights(data data.Payload) []string {
	si := data.Insights
	links := []string{
		si.Header.URL,
		si.Header.ProjectSISource,
		si.Project.Homepage,
		si.Project.Roadmap,
		si.Project.Funding,
		si.Project.Documentation.DetailedGuide,
		si.Project.Documentation.CodeOfConduct,
		si.Project.Documentation.QuickstartGuide,
		si.Project.Documentation.ReleaseProcess,
		si.Project.Documentation.SignatureVerification,
		si.Project.Vulnerability.BugBountyProgram,
		si.Project.Vulnerability.SecurityPolicy,
		si.Repository.URL,
		si.Repository.License.URL,
		si.Repository.Security.Assessments.Self.Evidence,
	}
	for _, repo := range si.Project.Repositories {
		links = append(links, repo.URL)
	}

	for _, repo := range si.Repository.Security.Assessments.ThirdParty {
		links = append(links, repo.Evidence)
	}

	for _, tool := range si.Repository.Security.Tools {
		links = append(links, tool.Results.Adhoc.Location)
		links = append(links, tool.Results.CI.Location)
		links = append(links, tool.Results.Release.Location)
	}
	return links
}

func insecureURI(uri string) bool {
	if !strings.HasPrefix(uri, "https://") ||
		!strings.HasPrefix(uri, "ssh:") ||
		!strings.HasPrefix(uri, "git:") ||
		!strings.HasPrefix(uri, "git@") {
		return false
	}
	return true
}

func ensureInsightsLinksUseHTTPS(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	links := getLinksFromInsights(data)
	var badURIs []string
	for _, link := range links {
		if insecureURI(link) {
			badURIs = append(badURIs, link)
		}
	}
	if len(badURIs) > 0 {
		return layer4.Failed, fmt.Sprintf("The following links do not use HTTPS: %v", strings.Join(badURIs, ", "))
	}
	return layer4.Passed, "All links use HTTPS"
}

func ensureGitHubWebsiteLinkUsesHTTPS(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if insecureURI(data.WebsiteURL) {
		return layer4.Passed, fmt.Sprintf("The website URI linked from GitHub uses an insecure protocol: %v", data.WebsiteURL)
	}
	return layer4.Passed, fmt.Sprintf("The website URI linked from GitHub uses a secure protocol: %v", data.WebsiteURL)
}
