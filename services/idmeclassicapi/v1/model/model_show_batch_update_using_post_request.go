package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpdateUsingPostRequest Request Object
type ShowBatchUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelUpdateDto `json:"body,omitempty"`
}

func (o ShowBatchUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUpdateUsingPostRequest", string(data)}, " ")
}
