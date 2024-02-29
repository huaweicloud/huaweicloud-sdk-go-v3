package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogicalDeleteBranchUsingPostRequest Request Object
type ShowLogicalDeleteBranchUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterModifierDto `json:"body,omitempty"`
}

func (o ShowLogicalDeleteBranchUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogicalDeleteBranchUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowLogicalDeleteBranchUsingPostRequest", string(data)}, " ")
}
