package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveAsUsingPostRequest Request Object
type SaveAsUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelSaveAsDto `json:"body,omitempty"`
}

func (o SaveAsUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAsUsingPostRequest struct{}"
	}

	return strings.Join([]string{"SaveAsUsingPostRequest", string(data)}, " ")
}
