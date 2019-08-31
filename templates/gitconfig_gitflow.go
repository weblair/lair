package templates

// TODO: Build git flow init into Lair.
// GitFlowGitconfig is some add-on text intended to be appended to the .git/config file.
// This appended text sets up Git Flow without having to run the git flow init command.
const GitFlowGitconfig = `[gitflow "branch"]
	master = master
	develop = develop
[gitflow "prefix"]
	feature = feature/
	release = release/
	hotfix = hotfix/
	support = support/
	versiontag = v
`
