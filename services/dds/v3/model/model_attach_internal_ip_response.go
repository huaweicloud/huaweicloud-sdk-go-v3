package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AttachInternalIpResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 新的内网IP。
	NewIp          *string `json:"new_ip,omitempty" xml:"new_ip"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachInternalIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternalIpResponse struct{}"
	}

	return strings.Join([]string{"AttachInternalIpResponse", string(data)}, " ")
}
