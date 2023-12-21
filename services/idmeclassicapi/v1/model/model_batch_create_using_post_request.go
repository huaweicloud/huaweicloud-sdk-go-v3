package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateUsingPostRequest Request Object
type BatchCreateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelCreateDto `json:"body,omitempty"`
}

func (o BatchCreateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateUsingPostRequest", string(data)}, " ")
}
