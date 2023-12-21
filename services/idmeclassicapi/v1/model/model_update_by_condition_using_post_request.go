package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateByConditionUsingPostRequest Request Object
type UpdateByConditionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoUpdateByConditionVoPersistableModelUpdateDto `json:"body,omitempty"`
}

func (o UpdateByConditionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateByConditionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"UpdateByConditionUsingPostRequest", string(data)}, " ")
}
