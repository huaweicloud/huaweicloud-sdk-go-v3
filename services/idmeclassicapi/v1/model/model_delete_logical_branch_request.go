package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogicalBranchRequest Request Object
type DeleteLogicalBranchRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterModifierDto `json:"body,omitempty"`
}

func (o DeleteLogicalBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogicalBranchRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogicalBranchRequest", string(data)}, " ")
}
