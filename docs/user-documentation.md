# **GoGit-Integration Help**

## Table of Contents

- [GoGit-Integration Help](#gogit-integration-help)
  - [Table of Contents](#table-of-contents)
  - [1 License](#1-license)
  - [2 Installation](#2-installation)
  - [3 Configuration](#3-configuration)
  - [4 Troubleshooting](#4-troubleshooting)

## 1 License

Copyright (c) Avanis GbmH, Anton Paul

All right reserved.

## 2 Installation

## 2.1 System Requirements

Since the GoGit-Integration is written in Go, it can be run on any system that supports Go.
Tho it has only really been tested on Windows 11.

## 2.2 Prerequisites

Install [Go 1.19](https://golang.org/doc/install) or newer.
```text
https://go.dev/dl/go1.19.5.windows-amd64.msi
```

## 2.3 Installation Process

## 2.3.1 Download the latest release

Download the latest release from the [releases page](https://github.com/AntonSkrub/GoGit-Integration/releases).
There are no releases yet tho.
Extract the archive and run the executable file called `GoGitIntegration.exe`.

### 2.3.2 Build from source

When you have Go installed, download the repository and run the following command in the root directory of the project:

```text
go build -o GoGitIntegration.exe cmd/main.go
```

This will create an executable file called `GoGitIntegration.exe` in the root directory of the project.

## 3 Configuration

At first startup, the application will create a default configuration file called `config.yml` in the root directory of the project.
The configuration file contains the following options:

`OrgaName: string`

The name of the GitHub organization to clone the repositories from.

`OrgaToken: string`

The access token of the GitHub organization to clone the repositories from.

`OrgaRepoType: string`

The type of repositories to clone from the organization. Can be either `all`, `public` or `private`.

`CloneUserRepos: bool`

Whether to clone repositories from configured users.

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

`UpdateInterval: int`

The interval in which the repositories should be updated.
The interval is given in cron syntax.

Example configuration:

```yaml
UpdateInterval: "0 */12 * * *"
```

This is for updating the repositories every 12 hours.
For more information about the cron syntax, see [this page](https://pkg.go.dev/github.com/robfig/cron/v3#hdr-CRON_Expression_Format).

`ListReferences: bool`

Whether to list the references of the repositories.

`LogCommits: bool`

Whether to log the latest commits of the repositories.

`LogLevel: int`

The log level of the application. Can be either `0` (no logging), `1` (error logging), `2` (warning logging), `3` (info logging), `4` (debug logging) or `5` (trace logging).

## 4 Troubleshooting

