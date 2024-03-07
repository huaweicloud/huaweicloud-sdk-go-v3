package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteReviseRequest Request Object
type ExecuteReviseRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionReviseDto `json:"body,omitempty"`
}

func (o ExecuteReviseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteReviseRequest struct{}"
	}

	return strings.Join([]string{"ExecuteReviseRequest", string(data)}, " ")
}
