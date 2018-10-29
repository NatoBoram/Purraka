package main

import "os"

// Paths
const (
	rootFolder   = "./Purraka"
	discordPath  = rootFolder + "/discord.json"
	databasePath = rootFolder + "/db.json"
	headerPath   = rootFolder + "/header.json"
	errorPath    = rootFolder + "/errors.log"
)

// Permissions
const (
	permPrivateDirectory os.FileMode = 0700
	permPrivateFile      os.FileMode = 0600
)
