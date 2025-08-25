package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppAccessKeyRequest Request Object
type CreateAppAccessKeyRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	Body *CreateAppAccessKeyRequestBody `json:"body,omitempty"`
}

func (o CreateAppAccessKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateAppAccessKeyRequest", string(data)}, " ")
}
