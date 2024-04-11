package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveChildNodeRequest Request Object
type BatchRemoveChildNodeRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoBatchOperateChildDto `json:"body,omitempty"`
}

func (o BatchRemoveChildNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveChildNodeRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveChildNodeRequest", string(data)}, " ")
}
