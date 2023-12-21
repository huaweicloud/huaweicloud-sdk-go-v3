package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteByConditionUsingPostRequest Request Object
type DeleteByConditionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoDeleteByConditionVo `json:"body,omitempty"`
}

func (o DeleteByConditionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteByConditionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"DeleteByConditionUsingPostRequest", string(data)}, " ")
}
