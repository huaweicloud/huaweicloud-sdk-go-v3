package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogicalDeleteByConditionUsingPostRequest Request Object
type ShowLogicalDeleteByConditionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoDeleteByConditionVo `json:"body,omitempty"`
}

func (o ShowLogicalDeleteByConditionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogicalDeleteByConditionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowLogicalDeleteByConditionUsingPostRequest", string(data)}, " ")
}
