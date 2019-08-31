package templates

// ReadmeMd is the template for generating the README.md file in a Gin project.
const ReadmeMd = `# $1 $2

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing 
purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
  - Go 1.12.5 (or later)
  - Postgres 11.4 (or later)
  - GNU Make 3.81 (or later)
  - [Swaggo CLI](https://github.com/swaggo/swag)
  - [Lair](https://github.com/weblair/lair.git) (Optional, but recommended).
  - [Air](https://github.com/cosmtrek/air) (Necessary only if you want to use ` + "`make run`" + `)

### Dev Setup (with Lair)
  1. ` + "`git clone git@github.com:$1/$2.git`" + `
  2. ` + "`go mod tidy`" + `
  3. ` + "`lair db create --seed`" + `
  4. ` + "`make run`" + `

### Dev Setup (without Lair)
  1. ` + "`git clone git@github.com:$1/$2.git`" + `
  2. ` + "`psql -U postgres -c \"CREATE DATABASE $2_development;\"`" + `
  3. ` + "`migrate -source file://db/migrations -database postgres://localhost:5432/$2_development?sslmode=disable up`" + `
  4. ` + "`go mod tidy`" + `
  5. ` + "`make run`" + `

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

` + "```" + `
Give an example
` + "```" + `

### And coding style tests

Explain what these tests test and why

` + "```" + `
Give an example
` + "```" + `

## Deployment

Add additional notes about how to deploy this on a live system

## Built With
* [Gin](https://gin-gonic.com) - API framework
* [Postgres](https://postgresql.org) - Database
* [Migrate](https://github.com/golang-migrate/migrate) - Database Migrations
* [Swaggo](https://github.com/swaggo/swag) - API Documentation

## Contributing
Before committing, be sure to:
1. Use [Git Flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)
2. [Sign your commits](https://git-scm.com/book/ms/v2/Git-Tools-Signing-Your-Work)
3. Run ` + "`gofmt -w -s .`" + `
4. Run ` + "`swag init`" + `
5. Use the [Go Reportcard CLI](https://github.com/gojp/goreportcard)

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the 
[tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Robert Hawk** - *Initial work* - [DerHabicht](https://github.com/DerHabicht)

See also the list of [contributors](https://github.com/$1/$2/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* [PurpleBooth](https://github.com/PurpleBooth) for the README template
* [Commissure](https://github.com/commissure) for the build meta-data Makefile
`
