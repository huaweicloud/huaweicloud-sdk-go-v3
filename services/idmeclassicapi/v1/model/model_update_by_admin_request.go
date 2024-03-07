package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateByAdminRequest Request Object
type UpdateByAdminRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModel `json:"body,omitempty"`
}

func (o UpdateByAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateByAdminRequest struct{}"
	}

	return strings.Join([]string{"UpdateByAdminRequest", string(data)}, " ")
}
