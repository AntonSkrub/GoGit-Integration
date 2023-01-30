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

Since the GoGit-Integration is written in Go, the system requirements are fairly slim and it should be able to run on most systems.

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

`Organizations: map[string]Account`

A map of organizations to clone the repositories from. It contains the following options:

`Name: string`

The name of the organization to clone the repositories from.

`Token: string`

The access token of the organization

`Option: string`

Determines which repositories should be cloned. 

The following options are available, of which the first is the default option:
- `all`: Clones all repositories of the organization.
- `public`: Clones all public repositories of the organization.
- `private`: Clones all private repositories of the organization.
- `forks`: Clones all forked repositories of the organization.
- `sources`: Clones all source repositories of the organization.
- `member`: Clones all repositories of the organization, in which the user is a member.

`BackupRepos: bool`

Whether to backup the repositories of the given organization.
If set to `true` the application clones the repositories of the organization to the `OutputPath` directory. Otherwise the organization and it's repositories will be skipped.

`ValidateName: bool`

Whether to validate the name of the organization. 
If set to `true` the application will check if the `full_name` of the repository contains the name of the organization and if not, the repository will not be cloned. Otherwise the application will clone all repositories found in the organization.

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

`LogLevel: int`

The log level of the application. Can be either `0` (no logging), `1` (error logging), `2` (warning logging), `3` (info logging), `4` (debug logging) or `5` (trace logging).

## 4 Troubleshooting

