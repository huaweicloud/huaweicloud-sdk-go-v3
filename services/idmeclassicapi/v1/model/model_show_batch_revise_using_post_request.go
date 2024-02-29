package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchReviseUsingPostRequest Request Object
type ShowBatchReviseUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionReviseDto `json:"body,omitempty"`
}

func (o ShowBatchReviseUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchReviseUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchReviseUsingPostRequest", string(data)}, " ")
}
