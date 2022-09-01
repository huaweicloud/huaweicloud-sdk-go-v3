package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateKmsTagRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	// 密钥ID
	KeyId string `json:"key_id" xml:"key_id"`

	Body *CreateKmsTagRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateKmsTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagRequest struct{}"
	}

	return strings.Join([]string{"CreateKmsTagRequest", string(data)}, " ")
}
