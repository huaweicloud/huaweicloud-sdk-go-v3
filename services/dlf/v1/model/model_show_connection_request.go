package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectionRequest Request Object
type ShowConnectionRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 连接名称.
	ConnectionName string `json:"connection_name"`
}

func (o ShowConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectionRequest", string(data)}, " ")
}
