package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeleteByConditionUsingPostRequest Request Object
type ShowDeleteByConditionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoDeleteByConditionVo `json:"body,omitempty"`
}

func (o ShowDeleteByConditionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeleteByConditionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowDeleteByConditionUsingPostRequest", string(data)}, " ")
}
