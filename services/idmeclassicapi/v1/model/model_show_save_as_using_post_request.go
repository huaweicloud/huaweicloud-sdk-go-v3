package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSaveAsUsingPostRequest Request Object
type ShowSaveAsUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistableModelSaveAsDto `json:"body,omitempty"`
}

func (o ShowSaveAsUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSaveAsUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowSaveAsUsingPostRequest", string(data)}, " ")
}
