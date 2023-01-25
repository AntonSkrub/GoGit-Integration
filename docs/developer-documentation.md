# **GoGit-Integration Developer documentation**

## Table of Contents

- [GoGit-Integration Developer documentation](#gogit-integration-developer-documentation)
  - [Table of Contents](#table-of-contents)
  - [1 Development](#1-development)
  - [2 Development environment](#2-development-environment)
  - [3 Dependencies](#3-dependencies)
  - [4 Building requirements](#4-building-requirements)
  - [5 Versioning](#5-versioning)
  - [6 Packages](#6-packages)

## 1 Development

```text
Anton Paul

Avanis GmbH
Meisenstraße 79a
33607 Bielefeld
Germany
https://avanis.de
```
## 2 Development environment

The software was developed under Windows using Visual Studio Code.
The project is purely written in Go in the version 1.19.5.

## 3 Dependencies

- Go 1.19
  - github.com/go-git/go-git/v5
  - github.com/robfig/cron/v3
  - github.com/sirupsen/logrus
  - gopkg.in/yaml.v3
  - github.com/Microsoft/go-winio
  - github.com/ProtonMail/go-crypto
  - github.com/acomagu/bufpipe
  - github.com/cloudflare/circl
  - github.com/emirpasic/gods
  - github.com/go-git/gcfg
  - github.com/go-git/go-billy/v5
  - github.com/imdario/mergo
  - github.com/jbenet/go-context
  - github.com/kevinburke/ssh_config
  - github.com/pjbgf/sha1cd
  - github.com/sergi/go-diff
  - github.com/skeema/knownhosts
  - github.com/xanzy/ssh-agent
  - golang.org/x/crypto
  - golang.org/x/net
  - golang.org/x/sys
  - gopkg.in/warnings.v0

## 4 Building requirements

[Golang 1.19.5](https://golang.org/dl/) or higher must be installed on the machine. The modules listed in the **Dependencies** section must be available.

Verify Go installation
```bash
go version
```

## 5 Versioning

There is no versioning yet.

## 6 Packages

### main

The `main` package is the starting point of the application. It contains the main function and the application logic. It's responsible for starting other software modules and handling the application lifecycle.

| Package dependencies | Description |
| -------------------- | ----------- |
| config | Used to initialize and handle the configuration during runtime |
| gitapi | Used to get a list of repository names to clone |
| gogit  | Used to created and update the local copies of the repositories |

### config

The `config` package defines the configuration structure and provides methods to initialize, load and handle the configuration

| Package dependencies | Description |
| -------------------- | ----------- |
| - | - |

### gitapi

The `gitapi` package provides methods to get a list of the repository names to clone.

| Package dependencies | Description |
| -------------------- | ----------- |
| config | Provides config-values like organization-names and usernames |

### gogit

The `gogit` package provides methods to clone and update the local copies of the repositories. It also provides methods to get some more information about the repositories, base on the `ListReferences` and `LogCommits` config parameters.

| Package dependencies | Description |
| -------------------- | ----------- |
| config | Provides config-values like the local `OutPutPath` for the repositories |
| gitapi | Provides the `gitapi.Repository` struct used to extract repository information like the `full_name` and `owner` |