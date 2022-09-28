package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteNodeEncryptdatasRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 加密数据ID
	EncryptdataId string `json:"encryptdata_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeleteNodeEncryptdatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeEncryptdatasRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodeEncryptdatasRequest", string(data)}, " ")
}
