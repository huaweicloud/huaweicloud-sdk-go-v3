package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照对象
type Snapshot struct {
	// 指定创建快照的集群ID

	ClusterId string `json:"cluster_id"`
	// 快照名称，要求唯一性且必须以字母开头，不区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符，长度为4～64个字符。

	Name string `json:"name"`
	// 快照描述

	Description *string `json:"description,omitempty"`
}

func (o Snapshot) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Snapshot struct{}"
	}

	return strings.Join([]string{"Snapshot", string(data)}, " ")
}
