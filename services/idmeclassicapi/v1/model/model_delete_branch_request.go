package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBranchRequest Request Object
type DeleteBranchRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterModifierDto `json:"body,omitempty"`
}

func (o DeleteBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBranchRequest struct{}"
	}

	return strings.Join([]string{"DeleteBranchRequest", string(data)}, " ")
}
