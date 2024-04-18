package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConformancePackRequestBody 更新合规规则包的请求体。
type UpdateConformancePackRequestBody struct {

	// 合规规则包名称。
	Name string `json:"name"`

	// 合规规则包参数。
	VarsStructure *[]VarsStructure `json:"vars_structure,omitempty"`
}

func (o UpdateConformancePackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConformancePackRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConformancePackRequestBody", string(data)}, " ")
}
