package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCreateUsingPostRequest Request Object
type ShowBatchCreateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelCreateDto `json:"body,omitempty"`
}

func (o ShowBatchCreateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCreateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchCreateUsingPostRequest", string(data)}, " ")
}
