# Go Bootstrap (base/skeleton)

Introduction
This is a repository intended to serve as a starting point if you want to bootstrap a project in Go.

It could be useful if you want to start from scratch a project. The idea is that you don't have to worry about the boilerplate.


## How To Start
You could manually clone this repo or just us it as a template

### Cloning the repository
We recommend to follow the next step by step process in order to avoid adding the bootstrap project commits to your project Git history:

Use this repositoy template
 - Clone your project
 - Move to the project directory: cd your-project-name
 - Run go mod tidy to download all dependencies


### Folder structure

```bash
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   ├── main.go
│   └── main_test.go
├── go.mod
├── go.sum
├── internal
│   └── private_func.go
├── pkg
│   └── logger
│       └── logger.go
```

We'll try to maintain this project as simple as possible, but Pull Requests are welcome!