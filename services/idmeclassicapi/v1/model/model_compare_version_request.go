package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareVersionRequest Request Object
type CompareVersionRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoCompareVersionVo `json:"body,omitempty"`
}

func (o CompareVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionRequest struct{}"
	}

	return strings.Join([]string{"CompareVersionRequest", string(data)}, " ")
}
