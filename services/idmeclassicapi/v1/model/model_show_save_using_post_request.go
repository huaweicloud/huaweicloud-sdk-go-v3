package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSaveUsingPostRequest Request Object
type ShowSaveUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelSaveDto `json:"body,omitempty"`
}

func (o ShowSaveUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSaveUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowSaveUsingPostRequest", string(data)}, " ")
}
