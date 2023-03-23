package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadFilePublisherResponse struct {

	// 返回值
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadFilePublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFilePublisherResponse struct{}"
	}

	return strings.Join([]string{"UploadFilePublisherResponse", string(data)}, " ")
}
