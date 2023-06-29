package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EipResourcesAvailableV3RequestBody This is a auto create Body Object
type EipResourcesAvailableV3RequestBody struct {

	// 公共池类型
	Type *string `json:"type,omitempty"`

	// 查询的公共池数量
	Limit int32 `json:"limit"`
}

func (o EipResourcesAvailableV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipResourcesAvailableV3RequestBody struct{}"
	}

	return strings.Join([]string{"EipResourcesAvailableV3RequestBody", string(data)}, " ")
}
