package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AttachEipResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 节点名称。
	NodeName *string `json:"node_name,omitempty" xml:"node_name"`

	// 公网IP的ID。
	PublicIpId *string `json:"public_ip_id,omitempty" xml:"public_ip_id"`

	// 公网IP。
	PublicIp       *string `json:"public_ip,omitempty" xml:"public_ip"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipResponse struct{}"
	}

	return strings.Join([]string{"AttachEipResponse", string(data)}, " ")
}
