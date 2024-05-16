package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlameLineTreeInfo struct {

	// 开始时间，比如1704271204595
	From int64 `json:"from"`

	// 结束时间, 比如1704275169491
	To int64 `json:"to"`

	// 数据类型, CPU 或者 LATENCY
	Type string `json:"type"`

	// 实例id
	InstanceId int64 `json:"instance_id"`

	// api的url,比如: GET_/user/{id}
	Api string `json:"api"`

	// 实例所在区域
	Region string `json:"region"`
}

func (o FlameLineTreeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlameLineTreeInfo struct{}"
	}

	return strings.Join([]string{"FlameLineTreeInfo", string(data)}, " ")
}
