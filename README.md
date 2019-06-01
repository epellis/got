# Got
Got is a Toy Version Control System (Like `git`, `hg`, `svn`)

# Installation
`go get -t -v github.com/epellis/got/...`

# Workflows
Here are the steps for interacting with the Got VCS.

## Creating a new project
```
# Create a new project in the current directory
$ got init <project-name>

# create your files, directories...

# Add all files in the current directory to project
$ got add .

# Sync your changes to the master server
$ got sync
```

## Editing a file in your project
```
# Sync for the latest changes
$ got sync

# edit your file, call it <file-a>

$ got add <file-a>

$ got sync
```

## Reverting to a previous version
_Work in Progress..._

## Switching branches
_Work in Progress..._
