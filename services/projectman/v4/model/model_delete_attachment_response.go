package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAttachmentResponse Response Object
type DeleteAttachmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteAttachmentResponse", string(data)}, " ")
}
