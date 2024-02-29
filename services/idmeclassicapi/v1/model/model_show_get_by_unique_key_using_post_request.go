package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetByUniqueKeyUsingPostRequest Request Object
type ShowGetByUniqueKeyUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	// 方法名称，格式为：getBy{uniqueName}。  uniqueName：表示“唯一键”为“是”的属性英文名称。
	GetUniqueFieldMethod string `json:"getUniqueFieldMethod"`

	Body *RdmParamVoPersistableModelUniqueKeyDto `json:"body,omitempty"`
}

func (o ShowGetByUniqueKeyUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetByUniqueKeyUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowGetByUniqueKeyUsingPostRequest", string(data)}, " ")
}
