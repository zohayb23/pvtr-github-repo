# Privateer Plugin for GitHub Repositories

This plugin is designed to test a GitHub repository using automated assessments compatible with the [Simplified Compliance Infrastructure](https://github.com/revanite-io/sci) Layer 4 data types.

## Work in Progress

Assessment development is currently addressing the [Open Source Project Security Baseline v2025.02.25](https://baseline.openssf.org).

As possible, the goal is to work on the OSPS Baseline maturity levels from the lowest to highest.

While development is ongoing, the best way to run the plugin is to pull the code locally and run the local Dockerfile.

1. Pull the repo
2. Modify `example-config.yml` to use your values, and rename it to `config.yml`
3. Build the Docker Image: `make docker build`
4. Run the Docker Image: `make docker run`
5. Review the output in the directory you've specified in your config file

> [NOTE!]
> The Dockerfile does not currenly provide a prettified output.
> Review the `.log` file to see the simplest results info.

