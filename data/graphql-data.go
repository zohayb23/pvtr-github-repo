package data

import (
	"github.com/shurcooL/githubv4"
)

type GraphqlData struct {
	Repository struct {
		Name                    string
		IsPrivate               bool
		HasDiscussionsEnabled   bool
		HasIssuesEnabled        bool
		IsSecurityPolicyEnabled bool

		DefaultBranchRef struct {
			Name          string
			RefUpdateRule struct {
				AllowsDeletions              bool
				AllowsForcePushes            bool
				RequiredApprovingReviewCount int
			}
			BranchProtectionRule struct {
				RestrictsPushes             bool // This didn't give an accurate result
				RequiresApprovingReviews    bool // This gave an accurate result
				RequiresCommitSignatures    bool
				RequiresStatusChecks        bool
				RequiredStatusCheckContexts []string
			}
			Target struct {
				OID    string `graphql:"oid"` // Latest commit SHA
				Commit struct {
					Status struct {
						State    string // Overall commit status
						Contexts []struct {
							Context     string
							Description string
							State       string
							TargetURL   string `graphql:"targetUrl"`
						}
					} `graphql:"status"` // Classic status API

					AssociatedPullRequests struct {
						Nodes []struct {
							StatusCheckRollup struct {
								Commit struct {
									CheckSuites struct {
										Nodes []struct {
											CheckRuns struct {
												Nodes []struct {
													Name string `graphql:"name"`
												}
											} `graphql:"checkRuns(first: 25)"`
										}
									} `graphql:"checkSuites(first: 25)"`
								}
							}
						}
					} `graphql:"associatedPullRequests(last: 1)"`
				} `graphql:"... on Commit"`
			} `graphql:"target"`
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
