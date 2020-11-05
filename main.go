package main

import (
	"context"
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

	log.Printf("calling server '%s'", baseURL)
	log.Printf("sending notification for image '%s'", image)

	client := update_docker_image.NewUpdateDockerImageProtobufClient(baseURL, &http.Client{})
	_, err := client.UpdateImage(context.Background(), &update_docker_image.UpdateImageReq{
		Image: image,
	})
	if err != nil {
		log.Fatal("error calling UpdateImage: " + err.Error())
	}
}
