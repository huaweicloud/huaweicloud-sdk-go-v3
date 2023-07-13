package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnctionRequest Request Object
type DeleteConnctionRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 连接名称.
	ConnectionName string `json:"connection_name"`
}

func (o DeleteConnctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnctionRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnctionRequest", string(data)}, " ")
}
