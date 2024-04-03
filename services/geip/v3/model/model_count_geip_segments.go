package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountGeipSegments struct {

	// 全域弹性公网IP段个数
	Count int32 `json:"count"`
}

func (o CountGeipSegments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGeipSegments struct{}"
	}

	return strings.Join([]string{"CountGeipSegments", string(data)}, " ")
}
