package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareBusinessVersionRequest Request Object
type CompareBusinessVersionRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelCompareVersionVo `json:"body,omitempty"`
}

func (o CompareBusinessVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareBusinessVersionRequest struct{}"
	}

	return strings.Join([]string{"CompareBusinessVersionRequest", string(data)}, " ")
}
