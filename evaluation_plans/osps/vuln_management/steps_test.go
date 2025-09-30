package vuln_management

import (
	"testing"

	"github.com/ossf/gemara/layer4"
	"github.com/ossf/si-tooling/v2/si"
	"github.com/revanite-io/pvtr-github-repo/data"
	"github.com/stretchr/testify/assert"
)

type testingData struct {
	expectedResult   layer4.Result
	expectedMessage  string
	payloadData      interface{}
	assertionMessage string
}

func TestSastToolDefined(t *testing.T) {

	testData := []testingData{
		{
			expectedResult:   layer4.Passed,
			expectedMessage:  "Static Application Security Testing documented in Security Insights",
			assertionMessage: "Test for SAST integration enabled",
			payloadData: data.Payload{
				RestData: &data.RestData{
					Insights: si.SecurityInsights{
						Repository: si.Repository{
							Security: si.SecurityInfo{
								Tools: []si.Tool{
									{
										Type: "SAST",
										Integration: si.Integration{
											Adhoc: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			expectedResult:   layer4.Failed,
			expectedMessage:  "No Static Application Security Testing documented in Security Insights",
			assertionMessage: "Test for SAST integration present but not explicitly enabled",
			payloadData: data.Payload{
				RestData: &data.RestData{
					Insights: si.SecurityInsights{
						Repository: si.Repository{
							Security: si.SecurityInfo{
								Tools: []si.Tool{
									{
										Type: "SAST",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			expectedResult:   layer4.Failed,
			expectedMessage:  "No Static Application Security Testing documented in Security Insights",
			assertionMessage: "Test for Non SAST tool defined",
			payloadData: data.Payload{
				RestData: &data.RestData{
					Insights: si.SecurityInsights{
						Repository: si.Repository{
							Security: si.SecurityInfo{
								Tools: []si.Tool{
									{
										Type: "NotSast",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			expectedResult:   layer4.Failed,
			expectedMessage:  "No Static Application Security Testing documented in Security Insights",
			assertionMessage: "Test for no tools defined",
			payloadData: data.Payload{
				RestData: &data.RestData{
					Insights: si.SecurityInsights{
						Repository: si.Repository{
							Security: si.SecurityInfo{},
						},
					},
				},
			},
		},
	}

	for _, test := range testData {
		result, message := sastToolDefined(test.payloadData, nil)

		assert.Equal(t, test.expectedResult, result, test.assertionMessage)
		assert.Equal(t, test.expectedMessage, message, test.assertionMessage)
	}

}

func TestHasVulnerabilityDisclosurePolicy(t *testing.T) {
	tests := []struct {
		name            string
		payloadData     any
		expectedResult  layer4.Result
		expectedMessage string
	}{
		{
			name:            "Vulnerability disclosure policy present",
			expectedResult:  layer4.Passed,
			expectedMessage: "Vulnerability disclosure policy was specified in Security Insights data",
			payloadData: data.Payload{
				RestData: &data.RestData{
					Insights: si.SecurityInsights{
						Project: si.Project{
							Vulnerability: si.VulnReport{
								SecurityPolicy: "https://example.com/SECURITY.md",
							},
						},
					},
				},
				GraphqlRepoData: &data.GraphqlRepoData{},
			},
		},
		{
			name:            "Vulnerability disclosure policy missing",
			expectedResult:  layer4.Failed,
			expectedMessage: "Vulnerability disclosure policy was NOT specified in Security Insights data",
			payloadData: data.Payload{
				RestData: &data.RestData{
					Insights: si.SecurityInsights{
						Project: si.Project{
							Vulnerability: si.VulnReport{
								SecurityPolicy: "",
							},
						},
					},
				},
				GraphqlRepoData: &data.GraphqlRepoData{},
			},
		},
		{
			name:            "Invalid payload",
			expectedResult:  layer4.Unknown,
			expectedMessage: "Malformed assessment: expected payload type data.Payload, got string (invalid_payload)",
			payloadData:     "invalid_payload",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, message := hasVulnerabilityDisclosurePolicy(test.payloadData, nil)
			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedMessage, message)
		})
	}
}

func TestHasPublicVulnerabilityDisclosure(t *testing.T) {
	tests := []struct {
		name            string
		payloadData     any
		expectedResult  layer4.Result
		expectedMessage string
	}{
		{
			name:            "One published security advisory",
			expectedResult:  layer4.Passed,
			expectedMessage: "Found 1 published security advisory",
			payloadData: data.Payload{
				RestData: &data.RestData{
					SecurityAdvisories: []data.SecurityAdvisory{
						{
							GhsaId:      "GHSA-xxxx-xxxx-xxxx",
							Summary:     "Test advisory",
							State:       "published",
							PublishedAt: "2024-01-01T00:00:00Z",
						},
					},
				},
				GraphqlRepoData: &data.GraphqlRepoData{},
			},
		},
		{
			name:            "Multiple published security advisories",
			expectedResult:  layer4.Passed,
			expectedMessage: "Found 3 published security advisories",
			payloadData: data.Payload{
				RestData: &data.RestData{
					SecurityAdvisories: []data.SecurityAdvisory{
						{
							GhsaId:      "GHSA-xxxx-xxxx-xxxx",
							Summary:     "First advisory",
							State:       "published",
							PublishedAt: "2024-01-01T00:00:00Z",
						},
						{
							GhsaId:      "GHSA-yyyy-yyyy-yyyy",
							Summary:     "Second advisory",
							State:       "published",
							PublishedAt: "2024-02-01T00:00:00Z",
						},
						{
							GhsaId:      "GHSA-zzzz-zzzz-zzzz",
							Summary:     "Third advisory",
							State:       "published",
							PublishedAt: "2024-03-01T00:00:00Z",
						},
					},
				},
				GraphqlRepoData: &data.GraphqlRepoData{},
			},
		},
		{
			name:            "No published security advisories",
			expectedResult:  layer4.Failed,
			expectedMessage: "No published security advisories found",
			payloadData: data.Payload{
				RestData: &data.RestData{
					SecurityAdvisories: []data.SecurityAdvisory{},
				},
				GraphqlRepoData: &data.GraphqlRepoData{},
			},
		},
		{
			name:            "Invalid payload",
			expectedResult:  layer4.Unknown,
			expectedMessage: "Malformed assessment: expected payload type data.Payload, got string (invalid_payload)",
			payloadData:     "invalid_payload",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, message := hasPublicVulnerabilityDisclosure(test.payloadData, nil)
			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedMessage, message)
		})
	}
}
