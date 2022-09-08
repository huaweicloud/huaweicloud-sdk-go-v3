package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadProductRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *UploadProductRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadProductRequest struct{}"
	}

	return strings.Join([]string{"UploadProductRequest", string(data)}, " ")
}
