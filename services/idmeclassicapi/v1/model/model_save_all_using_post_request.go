package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveAllUsingPostRequest Request Object
type SaveAllUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelSaveAllDto `json:"body,omitempty"`
}

func (o SaveAllUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAllUsingPostRequest struct{}"
	}

	return strings.Join([]string{"SaveAllUsingPostRequest", string(data)}, " ")
}
