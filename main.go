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
		log.Fatal("error writing %s: %s", path, err)
	} else {
		log.Printf("wrote %s successfully", path)
	}

	secret, err := client.Logical().Read(path)
	if err != nil {
		log.Fatal("error reading %s: %s", path, err)
	} else {
		data, err := json.Marshal(secret.Data)
		if err != nil {
			log.Fatal("error decoding %s: %s", path, err)
		} else {
			log.Printf("read %s as: %s", path, data)
		}
	}

	_, err = client.Logical().Delete(path)
	if err != nil {
		log.Fatal("error deleting %s: %s", path, err)
	} else {
		secret, _ := client.Logical().Read(path)
		if secret != nil {
			log.Fatal("error deleting %s: still readable after delete", path)
		} else {
			log.Printf("deleted %s successfully", path)
		}
	}
}
