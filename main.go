package main

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"flag"
	"log"
	"net/http"

	update_docker_image "github.com/bsdlp/update-docker-image"
)

func main() {
	flag.Parse()
	baseURL := flag.Arg(0)
	if baseURL == "" {
		log.Fatal("server_url is required")
	}
	image := flag.Arg(1)
	if image == "" {
		log.Fatal("image is required")
	}
	encodedPrivateKey := flag.Arg(2)
	var privateKey ed25519.PrivateKey
	if encodedPrivateKey != "" {
		decodedKey, err := base64.StdEncoding.DecodeString(encodedPrivateKey)
		if err != nil {
			log.Fatal("invalid private key")
		}
		privateKey = ed25519.PrivateKey(decodedKey)
	}

	log.Printf("calling server '%s'", baseURL)
	log.Printf("sending notification for image '%s'", image)

	req := &update_docker_image.UpdateImageReq{
		Image: image,
	}

	if privateKey != nil {
		req.Signature = ed25519.Sign(privateKey, []byte(image))
	}

	client := update_docker_image.NewUpdateDockerImageProtobufClient(baseURL, &http.Client{})
	_, err := client.UpdateImage(context.Background(), req)
	if err != nil {
		log.Fatal("error calling UpdateImage: " + err.Error())
	}
}
