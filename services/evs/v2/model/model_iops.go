package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Iops
type Iops struct {

	// 冻结标签。
	Frozened bool `json:"frozened"`

	// 云硬盘iops标识。
	Id string `json:"id"`

	// iops大小。
	TotalVal int32 `json:"total_val"`
}

func (o Iops) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Iops struct{}"
	}

	return strings.Join([]string{"Iops", string(data)}, " ")
}
