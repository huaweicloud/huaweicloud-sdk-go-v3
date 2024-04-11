package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetRootRequest Request Object
type ShowGetRootRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdDto `json:"body,omitempty"`
}

func (o ShowGetRootRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetRootRequest struct{}"
	}

	return strings.Join([]string{"ShowGetRootRequest", string(data)}, " ")
}
