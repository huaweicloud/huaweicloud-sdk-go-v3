package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowGetUsingPostRequest Request Object
type BatchShowGetUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsDecryptDto `json:"body,omitempty"`
}

func (o BatchShowGetUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowGetUsingPostRequest struct{}"
	}

	return strings.Join([]string{"BatchShowGetUsingPostRequest", string(data)}, " ")
}
