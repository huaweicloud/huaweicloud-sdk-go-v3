package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPublisherIconResponse Response Object
type UploadPublisherIconResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadPublisherIconResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPublisherIconResponse struct{}"
	}

	return strings.Join([]string{"UploadPublisherIconResponse", string(data)}, " ")
}
