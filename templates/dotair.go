package templates

// Dotair is the template for the air tool's local config file.
// This template is a modified version of the example config locate here:
// https://github.com/cosmtrek/air/blob/master/air.conf.example
const Dotair = `# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory
# . or absolute path, please note that the directories following must be under root
root = "." 
# Optional! If 'watch_dir' is empty, use 'root'.
watch_dir = ""
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use 'make' as well.
cmd = "make"
# Binary file yields from 'bin'.
bin = "bin/$2"
# This log file places in your tmp_dir.
log = "air_errors.log"
# Watch these filename extensions.
include_ext = ["go", "yml"]
# Ignore these filename extensions or directories.
exclude_dir = []
# There's no need to trigger build each time file changes if it's too frequently.
delay = 1000 # ms

[log]
# Show log time
time = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"
# app = "white"
`
