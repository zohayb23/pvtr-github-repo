package data

import (
	"github.com/shurcooL/githubv4"
)

type GraphqlData struct {
	// Need to update token for this
	Organization struct {
		RequiresTwoFactorAuthentication bool
		WebCommitSignoffRequired        bool
	} `graphql:"organization(login: $owner)"`

	Repository struct {
		WebCommitSignoffRequired bool
		Name                     string
		HasDiscussionsEnabled    bool
		HasIssuesEnabled         bool
		IsSecurityPolicyEnabled  bool
		DefaultBranchRef         struct {
			Name          string
			RefUpdateRule struct {
				AllowsDeletions              bool
				AllowsForcePushes            bool
				RequiredApprovingReviewCount int
			}
			BranchProtectionRule struct {
				RestrictsPushes          bool // This didn't give an accurate result
				RequiresApprovingReviews bool // This gave an accurate result
				RequiresCommitSignatures bool
				RequiresStatusChecks     bool
			}
		}
		LicenseInfo struct {
			Name   string
			SpdxId string
			Url    string
		}
		LatestRelease struct {
			Description string
		}
		ContributingGuidelines struct {
			Body         string
			ResourcePath githubv4.URI
		}
	} `graphql:"repository(owner: $owner, name: $name)"`
}
