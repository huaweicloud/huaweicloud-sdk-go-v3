package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogicalDeleteUsingPostRequest Request Object
type ShowLogicalDeleteUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdModifierDto `json:"body,omitempty"`
}

func (o ShowLogicalDeleteUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogicalDeleteUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowLogicalDeleteUsingPostRequest", string(data)}, " ")
}
