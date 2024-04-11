package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddChildNodeRequest Request Object
type BatchAddChildNodeRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoBatchOperateChildDto `json:"body,omitempty"`
}

func (o BatchAddChildNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddChildNodeRequest struct{}"
	}

	return strings.Join([]string{"BatchAddChildNodeRequest", string(data)}, " ")
}
