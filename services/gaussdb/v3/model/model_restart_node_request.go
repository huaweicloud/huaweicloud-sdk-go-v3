package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartNodeRequest 重启节点请求体
type RestartNodeRequest struct {

	// 实例节点是否延迟重启。默认false，立即重启。  - true: 延迟重启，实例节点将在运维时间窗内自动重启。 - false: 立即重启。
	Delay *bool `json:"delay,omitempty"`
}

func (o RestartNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartNodeRequest struct{}"
	}

	return strings.Join([]string{"RestartNodeRequest", string(data)}, " ")
}
