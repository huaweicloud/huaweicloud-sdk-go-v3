package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UcbParam struct {

	// 折中参数。
	Alpha float64 `json:"alpha"`

	// 最小行为次数。
	MinUsedNum int32 `json:"min_used_num"`
}

func (o UcbParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcbParam struct{}"
	}

	return strings.Join([]string{"UcbParam", string(data)}, " ")
}
