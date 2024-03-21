package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountGlobalEips struct {

	// 全域弹性公网IP个数
	Count int32 `json:"count"`
}

func (o CountGlobalEips) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGlobalEips struct{}"
	}

	return strings.Join([]string{"CountGlobalEips", string(data)}, " ")
}
