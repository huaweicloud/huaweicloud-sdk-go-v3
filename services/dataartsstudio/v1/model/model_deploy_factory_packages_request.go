package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployFactoryPackagesRequest Request Object
type DeployFactoryPackagesRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 有Body体的情况下必须，无Body体的情况下则无需填写和校验，默认值：application/json
	ContentType *string `json:"Content-Type,omitempty"`

	Body *DeployFactoryPackagesRequestBody `json:"body,omitempty"`
}

func (o DeployFactoryPackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployFactoryPackagesRequest struct{}"
	}

	return strings.Join([]string{"DeployFactoryPackagesRequest", string(data)}, " ")
}
