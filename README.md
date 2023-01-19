# GoGit-Integration

[![Linters](https://github.com/AntonSkrub/GoGit-Integration/actions/workflows/linters.yml/badge.svg)](https://github.com/AntonSkrub/GoGit-Integration/actions/workflows/linters.yml)

## Description

The GoGit-Integration is a small application which maintains local copies of an organizations, or users, GitHub repositories. The application keeps the local copies up to date with the remote repositories, based on a timed interval [(later configurable)]. Later in development the application should also offer a visual overview of the repositories and their status.

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

At first startup, the program will create a default configuration file called `config.yml` in the root directory of the project.
The configuration file contains the following options:

```yaml
OrgaName: Default Orga
OrgaToken: ""
UserName: Default User
UserToken: ""
OutputPath: ../Repo-Backups/
ListReferences: true
LogCommits: false
LogLevel: 6
```
