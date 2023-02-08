# **GoGit-Integration Help**

## Table of Contents

- [GoGit-Integration Help](#gogit-integration-help)
  - [Table of Contents](#table-of-contents)
  - [1 License](#1-license)
  - [2 Installation](#2-installation)
    - [2.1 System Requirements](#21-system-requirements)
    - [2.2 Network Requirements](#22-network-requirements)
    - [2.3 Prerequisites](#23-prerequisites)
    - [2.3 Installation Process](#23-installation-process)
  - [3 Configuration](#3-configuration)
    - [3.1 Accounts](#31-accounts)
    - [3.2 Remaining Options](#32-remaining-options)
  - [4 Troubleshooting](#4-troubleshooting)

## 1 License

Copyright (c) Avanis GbmH, Anton Paul

All right reserved.

## 2 Installation

## 2.1 System Requirements

Since the GoGit-Integration is written in Go, the system requirements are fairly slim and it should be able to run on most systems.

## 2.2 Network Requirements

The following protocols are needed to run the application:
`HTTP` -> To download the repositories from GitHub.

## 2.3 Prerequisites

Install [Go 1.19](https://golang.org/doc/install) or newer.
```text
https://go.dev/dl/go1.19.5.windows-amd64.msi
```

## 2.3 Installation Process

You can either download the latest release from the [release page](https://github.com/AntonSkrub/GoGit-Integration/releases), or build the application from source.
When you decide to download the latest release, you simply need to extract the files from the archive.
If you want to build the application from the source, you need to download/clone the repository and run the following command in the root directory of the project:

```text
go build -o GoGitIntegration.exe cmd/main.go
```

In either case you will now have a file called `GoGitIntegration.exe`, place it in the desired location and run it. The initial execution will create a configuration file called `config.yml` in the same directory as the executable.

## 3 Configuration

At first startup, the application will create a default configuration file called `config.yml` in the root directory of the project.
The configuration file contains the following options:

### 3.1 Accounts
`Accounts: map[string]Account`

The `Accounts` configuration option is a map containing the names of and some additional information about the GitHub accounts to clone the repositories from.
The map consists of the following options:

- `Name: string`: The name of the GitHub account to clone the repositories from. Mainly used for identification purposes.  
- `Token: string`: The access token of the GitHub account. This is needed to clone private repositories.  
- `Option: string`: The options can be combined with a comma and determine which repositories should be cloned. The available options are:  
  - `all`: Clones all repositories of the account. Default value.  
  - `owner`: Clones only the repositories, of which the account is the owner.  
  - `public`: Clones only the public repositories of the account.  
  - `private`: Clones only the private repositories of the account.  
  - `member`: Clones only the repositories, in which the account is a member.  
- `BackupRepos: bool`: Whether to backup the repositories of the given account. If set to `true` the application clones the repositories of the account to the `OutputPath` directory. Otherwise the account and it's repositories will be skipped.  
- `ValidateName: bool`: Whether the name of the account should be contained in the repositories `full_name` field. If set to `true` the application will only clone repositories, in which the `full_name` contains the name of the account. Otherwise the application will clone all repositories found in the account.  

Example configuration, with default values:

```yaml
Accounts:
  1st:
    Name: "GitHub-Username"
    Token: "<GitHub-Access-Token>"
    Option: "all"
    BackupRepos: true
    ValidateName: false
```
### 3.2 Remaining Options
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

The "loglevel" configuration option, sets the level of detail for the information that is output during runtime of the application.

The available log levels are:
- **0** >> Panic
- **1** >> Fatal
- **2** >> Error
- **3** >> Warn
- **4** >> Info
- **5** >> Debug
- **6** >> Trace
> The default is "Info"

## 4 Troubleshooting

