package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HeadGroupSpec Ray Head配置
type HeadGroupSpec struct {

	// 资源规格，从规格列表查询获取。
	SpecCode *string `json:"spec_code,omitempty"`
}

func (o HeadGroupSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeadGroupSpec struct{}"
	}

	return strings.Join([]string{"HeadGroupSpec", string(data)}, " ")
}
