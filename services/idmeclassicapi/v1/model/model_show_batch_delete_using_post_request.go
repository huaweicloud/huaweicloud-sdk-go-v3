package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchDeleteUsingPostRequest Request Object
type ShowBatchDeleteUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsModifierDto `json:"body,omitempty"`
}

func (o ShowBatchDeleteUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchDeleteUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchDeleteUsingPostRequest", string(data)}, " ")
}
