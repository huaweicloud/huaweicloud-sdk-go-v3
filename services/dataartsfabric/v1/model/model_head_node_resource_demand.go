package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HeadNodeResourceDemand headNode资源需求量配置
type HeadNodeResourceDemand struct {

	// 资源规格，从规格列表查询获取。
	SpecCode string `json:"spec_code"`
}

func (o HeadNodeResourceDemand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeadNodeResourceDemand struct{}"
	}

	return strings.Join([]string{"HeadNodeResourceDemand", string(data)}, " ")
}
