package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetByUniqueKeyRequest Request Object
type ShowGetByUniqueKeyRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	// 方法名称，格式为：getBy{uniqueName}。  uniqueName：表示“唯一键”为“是”的属性英文名称。
	GetUniqueFieldMethod string `json:"getUniqueFieldMethod"`

	Body *RdmParamVoPersistableModelUniqueKeyDto `json:"body,omitempty"`
}

func (o ShowGetByUniqueKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetByUniqueKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowGetByUniqueKeyRequest", string(data)}, " ")
}
