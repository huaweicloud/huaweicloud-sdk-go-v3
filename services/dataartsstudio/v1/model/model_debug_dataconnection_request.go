package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugDataconnectionRequest Request Object
type DebugDataconnectionRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	Body *ApigDataSourceVo `json:"body,omitempty"`
}

func (o DebugDataconnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugDataconnectionRequest struct{}"
	}

	return strings.Join([]string{"DebugDataconnectionRequest", string(data)}, " ")
}
