package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUsingPostRequest Request Object
type CreateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelCreateDto `json:"body,omitempty"`
}

func (o CreateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"CreateUsingPostRequest", string(data)}, " ")
}
