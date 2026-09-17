package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateNodeRequestBody 为资源绑定节点结构体
type AssociateNodeRequestBody struct {

	// 资源所关联节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o AssociateNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateNodeRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateNodeRequestBody", string(data)}, " ")
}
