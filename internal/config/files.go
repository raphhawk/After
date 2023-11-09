package config

import (
	"os"
	"path/filepath"
)

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, ".proglog", filename)
}

var (
	// self signed ca certificate
	CAFile = configFile("ca.pem")

	// server certificate
	ServerCertFile = configFile("server.pem")
	ServerKeyFile  = configFile("server-key.pem")

	// client(multiple) certificates
	RootClientCertFile   = configFile("root-client.pem")
	RootClientKeyFile    = configFile("root-client-key.pem")
	NobodyClientCertFile = configFile("nobody-client.pem")
	NobodyClientKeyFile  = configFile("nobody-client-key.pem")

	// acl model and policy files
	ACLModelFile  = configFile("model.conf")
	ACLPolicyFile = configFile("policy.csv")
)
