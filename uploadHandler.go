package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type MultipartUploadHandlerHandlerInput struct {
	Hostname        string
	Username        string
	Password        string
	ContentType     string
	ChannelID       int
	File            *VideoFileReader
	FileName        string
	DisplayName     string
	Privacy         int8
	Category        int
	CommentsEnabled bool
	DescriptionText string
	DownloadEnabled bool
	Language        string
	Licence         int
	NSFW            bool
	SupportText     string
	Tags            []string
}

func MultipartUploadHandler(input MultipartUploadHandlerHandlerInput) (err error) {
	/*
		This function handles the entire process
		of a multipart upload.

		Steps:
	*/
	/*
		Retrieve or request a new oauth token
		from the given username and password.

		(note possibility of failed authentication)
	*/
	oauthToken, err := GetUserTokenFromAPI(input.Hostname, input.Username, input.Password)
	if err != nil {
		return
	}
	log.Println(oauthToken)
	/*
		Initialize the multipart upload via
		https://docs.joinpeertube.org/api-rest-reference.html#operation/uploadResumableInit

		Oauth authorization via http header
		Params:
		-X-Upload-Content-Length
		-X-Upload-Content-Type (video/mp4 etc.)
		-channelId
		-filename
		-name
		required by me:
		-privacy
		not required:
		-category
		-commentsEnabled
		-description
		-downloadEnabled
		-language
		-licence (int)
		-nsfw
		-originallyPublishedAt (will not handle)
		-previewfile (will not handle this for now)
		-scheduleUpdate (will not handle)
		-support
		-tags
		-thumbnailfile (will not handle for now)
		-waitTranscoding (hard-code to true)

		Responces: 201 (best, all good to go). (200 already exists, send resumable instead. What do i do with this?)
		413 (bad, file size too large) 415 (bad, filetype not supported)

		if 201:
		JSON response:
		-Location (note that this includes the full path to the api resumable api call INCLUDING the upload id, the rest will be specified in headers in the PUT request)
		-Content-Length (What to do with this?)

	*/
	client := &http.Client{}
	initializeUrl := fmt.Sprintf("%s/api/v1/videos/upload-resumable", input.Hostname)
	initializePayload := map[string]interface{}{
		"channelId": input.ChannelID,
		"filename":  input.FileName,
		"name":      input.DisplayName,
	}
	initializePayloadBytes, err := json.Marshal(initializePayload)
	if err != nil {
		panic(err)
	}
	initialize, err := http.NewRequest("POST", initializeUrl, bytes.NewReader(initializePayloadBytes))
	if err != nil {
		panic(err)
	}

	// add oauth token in header
	initialize.Header.Add("Authorization", fmt.Sprintf("Bearer %s", oauthToken))

	initialize.Header.Add("X-Upload-Content-Length", fmt.Sprintf("%d", input.File.TotalBytes))
	initialize.Header.Add("X-Upload-Content-Type", input.ContentType)
	initialize.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(initialize)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	if resp.StatusCode != 201 {
		// BAD,
		log.Printf("initialize api call returned status code %d.", resp.StatusCode)
		fmt.Println(resp.Header)
		body, err2 := io.ReadAll(resp.Body)
		if err2 != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(body))
		panic("returned non 201 status code")
	}
	// response will be in headers
	fmt.Printf("%+v\n", resp.Header)

	// Get upload location
	defer resp.Body.Close()
	uploadLocation := resp.Header.Get("Location")

	if strings.HasPrefix(uploadLocation, "//") {
		uploadLocation = "https:" + uploadLocation
	} else {
		log.Println("Warning: recieved an upload location that doesn't begin with \"//\", i don't know what to do with this.")
		panic(nil)
	}
	fmt.Println("upload location", uploadLocation)

	/*
		For each part of the file, upload that part
		via https://docs.joinpeertube.org/api-rest-reference.html#operation/uploadResumable

		(note that this call will automatically finish
		the multipart upload when the final part is
		uploaded.)

		url:
		hostname + /api/v1/videos/upload-resumable?upload_id=12345abcde (given by Location in json response of initializer)

		PUT request
		headers:
		-Content-Length
		-Content-Range
		REQUEST BODY SCHEMA: application/octet-stream

		Responses:
		308 (good, not complete)
		200 (good, last chunk recieved, done.)
	*/
	for {
		chunk, err := input.File.GetNextChunk()
		if err != nil {
			panic(err)
		}
		if chunk.Finished {
			break
		}
		fmt.Println(chunk.MinByte, chunk.MaxByte, chunk.Length)
	}
	return
}
