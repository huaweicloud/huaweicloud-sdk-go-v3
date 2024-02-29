package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchDeleteBranchUsingPostRequest Request Object
type ShowBatchDeleteBranchUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionMasterIdsDtoVersionModelVersionMasterDto `json:"body,omitempty"`
}

func (o ShowBatchDeleteBranchUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchDeleteBranchUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchDeleteBranchUsingPostRequest", string(data)}, " ")
}
