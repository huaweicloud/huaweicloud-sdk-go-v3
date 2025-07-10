package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAttachmentResponse Response Object
type DownloadAttachmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DownloadAttachmentResponse", string(data)}, " ")
}
