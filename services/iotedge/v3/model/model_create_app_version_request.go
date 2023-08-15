package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppVersionRequest Request Object
type CreateAppVersionRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用版本
	Version string `json:"version"`

	Body *CreateAppVersionRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateAppVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateAppVersionRequest", string(data)}, " ")
}
