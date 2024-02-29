package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeleteLatestVersionUsingPostRequest Request Object
type ShowDeleteLatestVersionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterModifierDto `json:"body,omitempty"`
}

func (o ShowDeleteLatestVersionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeleteLatestVersionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowDeleteLatestVersionUsingPostRequest", string(data)}, " ")
}
