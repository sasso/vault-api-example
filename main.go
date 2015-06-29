package main

import (
	"encoding/json"
	"github.com/hashicorp/vault/api"
	"log"
)

func main() {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("err: %s", err)
	}

	if token := client.Token(); token != "" {
		log.Printf("using token client %s", token)
	} else {
		log.Fatal("no VAULT_TOKEN supplied!")
	}

	path := `secret/hello`
	_, err = client.Logical().Write(path, map[string]interface{}{"audience": "world"})
	if err != nil {
		log.Fatal("error writing secret/hello: %s", err)
	} else {
		log.Printf("wrote %s successfully", path)
	}

	secret, err := client.Logical().Read(path)
	if err != nil {
		log.Fatal("error reading secret/hello: %s", err)
	} else {
		data, err := json.Marshal(secret.Data)
		if err != nil {
			log.Fatal("error decoding secret/hello: %s", err)
		} else {
			log.Printf("read %s as: %s", path, data)
		}
	}

}
