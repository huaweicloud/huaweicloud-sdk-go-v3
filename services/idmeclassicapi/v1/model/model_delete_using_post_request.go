package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUsingPostRequest Request Object
type DeleteUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdModifierDto `json:"body,omitempty"`
}

func (o DeleteUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUsingPostRequest struct{}"
	}

	return strings.Join([]string{"DeleteUsingPostRequest", string(data)}, " ")
}
