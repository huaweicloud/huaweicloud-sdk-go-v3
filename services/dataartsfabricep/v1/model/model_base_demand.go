package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseDemand 资源需求量配置
type BaseDemand struct {

	// 最小数
	Min int32 `json:"min"`

	// 最大数，最小值为1，最大值为1000。
	Max int32 `json:"max"`
}

func (o BaseDemand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseDemand struct{}"
	}

	return strings.Join([]string{"BaseDemand", string(data)}, " ")
}
