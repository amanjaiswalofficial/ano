package main

import (
	"cron/job"

	"os"
)

func main() {

	job.SetupEnvironment("../config.json")
	credential_path := os.Getenv("CREDENTIALS_PATH") // Where to read credentials from
	token_path := os.Getenv("TOKEN_PATH")            // Where to store auth_token

	job.ConnectToGoogle(credential_path, token_path)
}
