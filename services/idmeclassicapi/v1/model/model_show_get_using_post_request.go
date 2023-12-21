package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetUsingPostRequest Request Object
type ShowGetUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdDecryptDto `json:"body,omitempty"`
}

func (o ShowGetUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowGetUsingPostRequest", string(data)}, " ")
}
