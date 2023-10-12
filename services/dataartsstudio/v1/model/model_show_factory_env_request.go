package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryEnvRequest Request Object
type ShowFactoryEnvRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`
}

func (o ShowFactoryEnvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryEnvRequest struct{}"
	}

	return strings.Join([]string{"ShowFactoryEnvRequest", string(data)}, " ")
}
