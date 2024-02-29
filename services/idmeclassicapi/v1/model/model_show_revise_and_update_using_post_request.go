package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviseAndUpdateUsingPostRequest Request Object
type ShowReviseAndUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionReviseAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o ShowReviseAndUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviseAndUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowReviseAndUpdateUsingPostRequest", string(data)}, " ")
}
