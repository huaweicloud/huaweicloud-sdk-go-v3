package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchLogicalDeleteUsingPostRequest Request Object
type ShowBatchLogicalDeleteUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsModifierDto `json:"body,omitempty"`
}

func (o ShowBatchLogicalDeleteUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchLogicalDeleteUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchLogicalDeleteUsingPostRequest", string(data)}, " ")
}
