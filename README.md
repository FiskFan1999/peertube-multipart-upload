# peertube-multipart-upload
Upload a video to a Peertube instance via a multipart upload using the REST API

# Reference
Refer to the [Peertube API reference](https://docs.joinpeertube.org/api-rest-reference.html#operation/uploadResumableInit) for the documentation of the API calls used internally.

- [Initialize](https://docs.joinpeertube.org/api-rest-reference.html#operation/uploadResumableInit)
- [Upload](https://docs.joinpeertube.org/api-rest-reference.html#operation/uploadResumable)
- [Cancel](https://docs.joinpeertube.org/api-rest-reference.html#operation/uploadResumableCancel) (Likely won't be implemented, as multipart uploads are automatically cancelled after a certain length of time if they are abandoned)

Note that a `/api/v1/videos/upload-resumable` call will automatically finish the upload when the final part is called.

# Chat
Please feel free to discuss this project on the OFTC network at #peertube-multipart-upload (also mirrored on oftc).
