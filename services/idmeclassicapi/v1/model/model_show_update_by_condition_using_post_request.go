package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateByConditionUsingPostRequest Request Object
type ShowUpdateByConditionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoUpdateByConditionVoPersistableModelUpdateDto `json:"body,omitempty"`
}

func (o ShowUpdateByConditionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateByConditionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowUpdateByConditionUsingPostRequest", string(data)}, " ")
}
