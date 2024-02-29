package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateUsingPostRequest Request Object
type ShowUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistableModelUpdateDto `json:"body,omitempty"`
}

func (o ShowUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowUpdateUsingPostRequest", string(data)}, " ")
}
