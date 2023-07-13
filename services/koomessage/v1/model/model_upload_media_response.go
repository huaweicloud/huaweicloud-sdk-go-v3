package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadMediaResponse Response Object
type UploadMediaResponse struct {
	Data           *UploadMediaResponseModel `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o UploadMediaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadMediaResponse struct{}"
	}

	return strings.Join([]string{"UploadMediaResponse", string(data)}, " ")
}
