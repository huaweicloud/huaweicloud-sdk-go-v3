package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInternalIpResponse Response Object
type AttachInternalIpResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 新的内网IP。
	NewIp          *string `json:"new_ip,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachInternalIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternalIpResponse struct{}"
	}

	return strings.Join([]string{"AttachInternalIpResponse", string(data)}, " ")
}
