package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelEipResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 节点名称。
	NodeName       *string `json:"node_name,omitempty" xml:"node_name"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelEipResponse struct{}"
	}

	return strings.Join([]string{"CancelEipResponse", string(data)}, " ")
}
