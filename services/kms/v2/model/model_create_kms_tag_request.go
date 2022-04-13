package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateKmsTagRequest struct {
	// 密钥ID

	KeyId string `json:"key_id"`

	Body *CreateKmsTagRequestBody `json:"body,omitempty"`
}

func (o CreateKmsTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagRequest struct{}"
	}

	return strings.Join([]string{"CreateKmsTagRequest", string(data)}, " ")
}
