# Lair

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
  - Go 1.12.5 (or later)
  - [Codegansta's Gin](go get github.com/codegangsta/gin)
  - [Dep](https://github.com/golang/dep)
  - [Migrate](https://github.com/golang-migrate/migrate)

### Installing
  1. `git clone git@github.com:weblair/lair.git`
  2. `go build`
  3. `cp lair $GOPATH/bin/`
  
After the initial install---and assuming you have GNU Make---you can use the included Makefile to build and deploy code
changes to $GOPATH/bin. This will use Lair's `build version --increment` tool to increment the build number in the
VERSION string. You can check your update was installed by running `weblair --version`.

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Migrate](https://github.com/golang-migrate/migrate) - Database Migrations

## Contributing
Before committing, be sure to:
 1. Use [Git Flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)
 2. [Sign your commits](https://git-scm.com/book/ms/v2/Git-Tools-Signing-Your-Work)
 3. Run `gofmt -s -w .\'

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Robert Hawk** - *Initial work* - [DerHabicht](https://github.com/DerHabicht)

See also the list of [contributors](https://github.com/weblair/deploy/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [PurpleBooth](https://github.com/PurpleBooth) for the README template
* [Commissure](https://github.com/commissure) for the build metadata Makefile
* The GitHub team for the [gitignore](https://github.com/github/gitignore) template
