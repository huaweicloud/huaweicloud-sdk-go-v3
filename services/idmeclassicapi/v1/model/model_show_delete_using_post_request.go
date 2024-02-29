package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeleteUsingPostRequest Request Object
type ShowDeleteUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdModifierDto `json:"body,omitempty"`
}

func (o ShowDeleteUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeleteUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowDeleteUsingPostRequest", string(data)}, " ")
}
