package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBranchRequest Request Object
type BatchDeleteBranchRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto `json:"body,omitempty"`
}

func (o BatchDeleteBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBranchRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteBranchRequest", string(data)}, " ")
}
