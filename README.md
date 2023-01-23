# GoGit-Integration

[![Linters](https://github.com/AntonSkrub/GoGit-Integration/actions/workflows/linters.yml/badge.svg)](https://github.com/AntonSkrub/GoGit-Integration/actions/workflows/linters.yml)

## Description

The GoGit-Integration is a small application which maintains local copies of an organizations, or users, GitHub repositories. The application keeps the local copies up to date with the remote repositories, based on a timed interval. Later the application should also offer a visual overview of the repositories and their status.

## Installation guide

### Prerequisites

To install the application you need to have Go installed on your machine.
If you don't have Go installed, you can find the installation instructions [here](https://golang.org/doc/install).

### Build from source

When you have Go installed, download the repository and run the following command in the root directory of the project:

```text
go build -o GoGitIntegration.exe cmd/main.go
```

This will create an executable file called `GoGitIntegration.exe` in the root directory of the project.

### Configuration

At first startup, the application will create a default configuration file called `config.yml` in the root directory of the project.
The configuration file contains the following options:

`OrgaName: string`

The name of the organization to clone the repositories from.

`OrgaToken: string`

The token of the organization to clone the repositories from.

`OrgaRepoType: string`

The type of repositories to clone from the organization. Can be either `all`, `public` or `private`.

`CloneUserRepos: bool`

Whether to also clone repositories of configured users.

`Users: map[string]User`

A map of users to clone the repositories from. It contains the following options:

`Name: string`

The name of the user to clone the repositories from.

`Token: string`

The token of the user to clone the repositories from.

`Affiliation: string`

The affiliation of the user to the repositories to clone. Can be either `owner`, `collaborator` or `organization_member`.

Example configuration:

```yaml
Users:
  1st:
    Name: "AntonSkrub/"
    Token: AccessToken
    Affiliation: "owner,collaborator,organization_member"
```

`OutputPath: string`

The path to the directory where the repositories should be cloned to.

`ListReferences: bool`

Whether to list the references of the repositories.

`LogCommits: bool`

Whether to log the latest commits of the repositories.

`LogLevel: int`

The log level of the application. Can be either `0` (no logging), `1` (error logging), `2` (warning logging), `3` (info logging), `4` (debug logging) or `5` (trace logging).


