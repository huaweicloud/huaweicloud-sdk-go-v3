package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchReviseAndUpdateUsingPostRequest Request Object
type ShowBatchReviseAndUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionReviseAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o ShowBatchReviseAndUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchReviseAndUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchReviseAndUpdateUsingPostRequest", string(data)}, " ")
}
