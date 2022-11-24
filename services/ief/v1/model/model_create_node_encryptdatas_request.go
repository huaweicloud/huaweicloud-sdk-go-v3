package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNodeEncryptdatasRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *EncryptDataNodeReq `json:"body,omitempty"`
}

func (o CreateNodeEncryptdatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodeEncryptdatasRequest struct{}"
	}

	return strings.Join([]string{"CreateNodeEncryptdatasRequest", string(data)}, " ")
}
