package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSaveAllUsingPostRequest Request Object
type ShowSaveAllUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelSaveAllDto `json:"body,omitempty"`
}

func (o ShowSaveAllUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSaveAllUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowSaveAllUsingPostRequest", string(data)}, " ")
}
