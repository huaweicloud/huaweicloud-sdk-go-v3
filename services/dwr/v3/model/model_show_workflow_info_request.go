package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWorkflowInfoRequest struct {

	// 工作流名。名称必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符。
	GraphName string `json:"graph_name"`
}

func (o ShowWorkflowInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowInfoRequest", string(data)}, " ")
}
