package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogicalDeleteLatestVersionUsingPostRequest Request Object
type ShowLogicalDeleteLatestVersionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterModifierDto `json:"body,omitempty"`
}

func (o ShowLogicalDeleteLatestVersionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogicalDeleteLatestVersionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowLogicalDeleteLatestVersionUsingPostRequest", string(data)}, " ")
}
