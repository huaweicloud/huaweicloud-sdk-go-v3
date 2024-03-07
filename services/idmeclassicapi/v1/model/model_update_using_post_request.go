package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUsingPostRequest Request Object
type UpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistableModelUpdateDto `json:"body,omitempty"`
}

func (o UpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"UpdateUsingPostRequest", string(data)}, " ")
}
