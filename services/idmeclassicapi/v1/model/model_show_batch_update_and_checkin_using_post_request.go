package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpdateAndCheckinUsingPostRequest Request Object
type ShowBatchUpdateAndCheckinUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUpdateAndCheckinDtoVersionModel `json:"body,omitempty"`
}

func (o ShowBatchUpdateAndCheckinUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpdateAndCheckinUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUpdateAndCheckinUsingPostRequest", string(data)}, " ")
}
