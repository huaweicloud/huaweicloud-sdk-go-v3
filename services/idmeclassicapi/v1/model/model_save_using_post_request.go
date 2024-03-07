package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveUsingPostRequest Request Object
type SaveUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelSaveDto `json:"body,omitempty"`
}

func (o SaveUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveUsingPostRequest struct{}"
	}

	return strings.Join([]string{"SaveUsingPostRequest", string(data)}, " ")
}
