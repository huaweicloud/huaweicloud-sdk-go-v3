package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviseUsingPostRequest Request Object
type ShowReviseUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionReviseDto `json:"body,omitempty"`
}

func (o ShowReviseUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviseUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowReviseUsingPostRequest", string(data)}, " ")
}
