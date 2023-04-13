package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Throughput struct {

	// 冻结标签。
	Frozened bool `json:"frozened"`

	// 云硬盘吞吐量标识。
	Id string `json:"id"`

	// 吞吐量大小。
	TotalVal int32 `json:"total_val"`
}

func (o Throughput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Throughput struct{}"
	}

	return strings.Join([]string{"Throughput", string(data)}, " ")
}
