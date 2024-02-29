package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCreateUsingPostRequest Request Object
type ShowCreateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistableModelCreateDto `json:"body,omitempty"`
}

func (o ShowCreateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCreateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowCreateUsingPostRequest", string(data)}, " ")
}
