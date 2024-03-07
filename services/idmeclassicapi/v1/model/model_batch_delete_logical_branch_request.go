package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteLogicalBranchRequest Request Object
type BatchDeleteLogicalBranchRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto `json:"body,omitempty"`
}

func (o BatchDeleteLogicalBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLogicalBranchRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLogicalBranchRequest", string(data)}, " ")
}
