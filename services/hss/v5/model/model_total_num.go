package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TotalNum 总数
type TotalNum struct {
}

func (o TotalNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalNum struct{}"
	}

	return strings.Join([]string{"TotalNum", string(data)}, " ")
}
