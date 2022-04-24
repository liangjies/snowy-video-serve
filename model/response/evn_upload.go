package response

import "snowy-video-serve/model"

type FileUploadResponse struct {
	File model.FileUpload `json:"file"`
}
