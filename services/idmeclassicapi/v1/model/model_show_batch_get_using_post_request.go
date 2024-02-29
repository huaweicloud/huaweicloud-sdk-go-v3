package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchGetUsingPostRequest Request Object
type ShowBatchGetUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsDecryptDto `json:"body,omitempty"`
}

func (o ShowBatchGetUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchGetUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchGetUsingPostRequest", string(data)}, " ")
}
