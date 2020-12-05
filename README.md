# Cognality Pose Server
Cognality's Go REST API saves real-time AR pose data stream with relation to S3 key to retrieve later for real-time 3d reconstruction and semantic segmentation.

# Endpoints
 ```
 GET https://cognality-pose-server-prod.herokuapp.com/frames/<token>/<id>
 ```
 retrieves a JSON of all the frames with the specified <id> using a securely hashed <token>
 
 ```
 POST https://cognality-pose-server-prod.herokuapp.com/frames
 BODY 
 {
	"token": "<token>",
	"recId": "<recording ID>",
	"frame": {
		"filename": "<S3 Key>",
		"pose": "<4x4 pose matrix>"
	}
}

RETURNS 
{
  "ok": <true/false>,
  "error": <error if ok is false>
}
 ```
 Posts a frame of a recording. Pose matrix must stay consistent, shape is 4x4.
 
 ## S3
 when storing images in the S3 bucket, this pattern must be used: 
 ``` https://<bucket>.s3.amazonaws.com/<recId>/<filename>```
 

# Todos
 - [ ] update hashing mechanism to something faster and more cryptographically secure
 - [x] update from express to fiber to make it go more zoom zoom
 - [x] docker
 - [x] transition from postgres to mongo
