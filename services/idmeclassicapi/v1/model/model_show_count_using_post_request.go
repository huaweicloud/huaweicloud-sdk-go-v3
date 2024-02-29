package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCountUsingPostRequest Request Object
type ShowCountUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoQueryRequestCountVo `json:"body,omitempty"`
}

func (o ShowCountUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCountUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowCountUsingPostRequest", string(data)}, " ")
}
