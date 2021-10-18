// Package loader
package loader

// Load - Loading config and architecture for app executing
func Load() {
	// 1. Init config
	loadConfig()
	// 2. Init Logger
	loggerToFile()
	// 3. Init PSQL
	loadPSQL()
	// 4. Init User
	initUser()
}
