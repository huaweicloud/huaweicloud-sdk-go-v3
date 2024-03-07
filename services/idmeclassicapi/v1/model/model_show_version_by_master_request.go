package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVersionByMasterRequest Request Object
type ShowVersionByMasterRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterQueryDto `json:"body,omitempty"`
}

func (o ShowVersionByMasterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionByMasterRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionByMasterRequest", string(data)}, " ")
}
