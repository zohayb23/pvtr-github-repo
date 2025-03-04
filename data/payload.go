package data

import (
	"context"
	"fmt"

	"github.com/ossf/si-tooling/v2/si"
	"github.com/privateerproj/privateer-sdk/config"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Payload struct {
	GraphqlData
	Insights si.SecurityInsights
	Metadata RepoMetadata
	Config   *config.Config
}

func Loader(config *config.Config) (payload interface{}, err error) {
	graphql, err := getGraphqlData(config)
	if err != nil {
		return nil, err
	}
	rest, err := getRestData(config)
	if err != nil {
		return nil, err
	}
	return interface{}(Payload{
		GraphqlData: graphql,
		Insights:    rest.Insights,
		Metadata:    rest.Metadata,
		Config:      config,
	}), nil
}

func getGraphqlData(config *config.Config) (data GraphqlData, err error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GetString("token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	variables := map[string]interface{}{
		"owner": githubv4.String(config.GetString("owner")),
		"name":  githubv4.String(config.GetString("repo")),
	}

	err = client.Query(context.Background(), &data, variables)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Error querying GitHub GraphQL API: %s", err.Error()))
	}
	return
}

func getRestData(config *config.Config) (data *RestData, err error) {
	r := &RestData{
		Config: config,
	}
	return r, r.Setup()
}
