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

The software was developed under Windows using Visual Studio Code to run on any machine.
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

The `main` package is the starting point of the application. It contains the main function and some additional application logic. It's responsible for starting other software modules and handling the applications lifecycle.

| Package dependencies | Description |
| -------------------- | ----------- |
| config | Used to initialize and handle the configuration during runtime |
| gitapi | Used to get a list of repository names to clone |
| gogit  | Used to created and update the local copies of the repositories |

#### **main()** function

First the function declares some flags, which can be passed to the program in the command-line. If one of these flags is detected, the application will perform certain actions, such as outputting a help text and exit afterwards. Otherwise the application will continue with the default behavior.
It loads the configuration file and initializes the logger. After that the names of the repositories to clone are requested from the GitHub API and will attempt to clone or update the local copies of the repositories. After the initial run, the application starts a cron-job to keep the local copies up to date with the remote repositories.

*Input values:*

The `main()` function does not accept any input values.

*Return values:*

The `main()` function does not return any values.

*Error handling:*

The `main()` function does not handle any errors.

---
#### **printHelp()** function

The `printHelp()` function is executed when the "-help" flag is passed in the CLI on startup. It print some basic information about the program, lists other usable starting flags and then lead to the immediate end of the program.

*Input values:*

There are no input parameters for this function.

*Return values:*

There are no return values in this function

*Error handling:*

The function is expected to never cause an error.

---
#### **printConfigExplanation()** function

The `printConfigExplanation()` function is executed when the "-confighelp" flag is passed in the CLI on startup. It displays the list of cluster configuration parameters with a short description to each one and then leads to the immediate end of the program

*Input values:*

There are no input parameters for this function.

*Return values:*

There are no return values in this function

*Error handling:*

The function is expected to never cause an error.

---
---
### config

The `config` package defines the configuration structure and provides methods to initialize, load and handle the configuration.

| Package dependencies | Description |
| -------------------- | ----------- |
| - | - |

#### **GetConfig()** function

This function returns the current configuration. If no configuration is detected at the specified location, by the `initConfig()` function, a new one will be created through the `createConfig()` function.

*Input values:*

There are no input parameters for this function.

*Return values:*

| Return value | Type | Usage |
| -------------- | ---- | ----- |
| instance | *Config | The current configuration, provides GitHub account names and access-tokens for example |

*Error handling:*

When no configuration is found, the `initConfig()` function will be called to create a new one. If this fails, the application will exit with an error.

---
#### **initConfig()** function

This function tries to load the configuration file from the specified location. If no configuration is found, a new one will be created with default values by the `createConfig()` function.

*Input values:*

There are no input parameters for this function.

*Return values:*
| Return value | Type | Usage |
| -------------- | ---- | ----- |
| error | error | Used to stores and return occurring errors |

*Error handling:*

Should an error occur while reading or creating the configuration file, the error will be passed to the caller-function and the application will exit with an error.

---
#### **createConfig()** function

This function creates a new configuration file with default values. It will be called when no configuration file is found at the specified location.

*Input values:*

There are no input parameters for this function.

*Return values:*
| Return value | Type | Usage |
| -------------- | ---- | ----- |
| error | error | Used to stores and return occurring errors |

*Error handling:*

When an error occurs while creating the configuration file, the error will be passed to the caller-function and the application will exit with an error.

---
---
### gitapi

The `gitapi` package provides methods to get a list of the repository names from the GitHub API.

| Package dependencies | Description |
| -------------------- | ----------- |
| config | Provides config-values like organization-names and usernames |

#### **GetRepoList()** function

The `GetRepoList()` function requests a/the list of repositories from the GitHub API. It will return a list of all repositories, accessible by the provided access-token, for each of the configured GitHub accounts.

*Input values:*
| Input variable | Type | Usage |
| -------------- | ---- | ----- |
| account | *config.Account | The GitHub Account to request repositories from |

*Return values:*

| Return value | Type | Usage |
| -------------- | ---- | ----- |
| repos | []Repositories | A list of repositories |

*Error handling:*

Errors that occur in this function will be directly printed to the console and the application will continue with the next account.

---

#### **buildURL()** function

The `buildURL()` function builds the URL to request the list of repositories from the GitHub API. It will return the URL as a string to it's caller function.

*Input values:*
| Input variable | Type | Usage |
| -------------- | ---- | ----- |
| paramValue | string | The value of the "type" parameter to add to the request URL |

*Return values:*
| Return value | Type | Usage |
| -------------- | ---- | ----- |
| urlString | string | The URL to request the list of repositories from the GitHub API |

*Error handling:*

Should an error occur while building the URL, it will be directly printed to the console and the application will continue with the next account.
---
---

### gogit

The `gogit` package provides methods to clone and update the local copies of the repositories. It also provides a method to list tags and branches of a repository.

| Package dependencies | Description |
| -------------------- | ----------- |
| config | Provides config-values like the local `OutPutPath` for the repositories |
| gitapi | Provides the `gitapi.Repository` struct used to extract repository information like the `full_name` and `owner` |

#### **UpdateLocalCopies()** function

The `UpdateLocalCopies()` function will, for each repository, try to open a folder named after the repositories `full_name` in the `OutputPath` folder. If the folder does not exist, it will be created and the `Clone()` function will be called to create a local copy in the just created folder. If the folder already exists, the local copy will be updated by pulling the latest changes from the remote repository.

*Input values:*
| Input variable | Type | Usage |
| -------------- | ---- | ----- |
| repos | []gitapi.Repository | A list of repositories |
| config | *config.Config | The current configuration |
| account | *config.Account | The GitHub Account to clone/update repositories from |

*Return values:*

This function does not have any return values.

*Error handling:*

Errors that occur in this function will be directly printed to the console and the application will continue with the next repository of the given account.

---

#### **Clone()** function

Should a repository not be found in the local `OutputPath` folder, the `Clone()` function will be called to create a local copy of the repository. It will clone the repository into a folder named after the repositories `full_name` in the `OutputPath` folder.

*Input values:*
| Input variable | Type | Usage |
| -------------- | ---- | ----- |
| name | string | The name of the repository to clone |
| config | *config.Config | The current configuration |
| account | *config.Account | The GitHub Account to clone repositories from |

*Return values:*

This function does not have any return values.

*Error handling:*

Errors that occur in this function will be directly printed to the console and the application will continue with the next repository of the given account.

---

#### **AccessRepo()** function

The `AccessRepo()` function requests a list of references (tags and branches) from GitHub, the returned list will be passed to the `ListRefs()` function, which print the list to the console.

*Input values:*
| Input variable | Type | Usage |
| -------------- | ---- | ----- |
| r | *gitapi.Repository | The repository to request references from |
| config | *config.Config | The current configuration |
| account | *config.Account | The GitHub Account the repository belongs to |

*Return values:*

There are no return values for this function.

*Error handling:*

Error will directly be printed to the console and the application will continue without listing the references of the repository.

---

#### **ListRefs()** function

The `ListRefs()` function first groups the references by their `ref_type` (tag or branch) and then prints the ordered list to the console.

*Input values:*
| Input variable | Type | Usage |
| -------------- | ---- | ----- |
| refs | []*plumbing.Reference | The list of references found on the repository |

*Return values:*

There are no return values for this function.

*Error handling:*

This function is expected to not cause any errors.