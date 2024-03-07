package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateByAdminRequest Request Object
type BatchUpdateByAdminRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModel `json:"body,omitempty"`
}

func (o BatchUpdateByAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateByAdminRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateByAdminRequest", string(data)}, " ")
}
