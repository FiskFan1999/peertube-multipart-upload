package main

type MultipartUploadHandlerHandlerInput struct {
	Hostname        string
	Username        string
	Password        string
	ContentType     string
	ChannelID       int
	Filename        string
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
	return
}
