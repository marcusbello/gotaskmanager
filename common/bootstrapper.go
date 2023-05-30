package common

func StartUp() {
	// initialize config
	initConfig()
	// Initialize private/public keys for JWT autentication
	initKeys()
	// Start a mongoDB Session
	createDbSession()
	// Add indexes into mongoDB
	addIndexes()
}
