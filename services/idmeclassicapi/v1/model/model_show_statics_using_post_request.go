package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStaticsUsingPostRequest Request Object
type ShowStaticsUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoQueryRequestStaticsVo `json:"body,omitempty"`
}

func (o ShowStaticsUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStaticsUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowStaticsUsingPostRequest", string(data)}, " ")
}
