package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 快照对象
type SnapshoDetail struct {
	// 快照对应的集群ID

	ClusterId string `json:"cluster_id"`
	// 快照大小，单位GB。

	Size float64 `json:"size"`
	// 快照名称

	Name string `json:"name"`
	// 快照描述

	Description string `json:"description"`
	// 快照完成的日期时间，格式为ISO8601: YYYY-MM-DDThh:mm:ssZ。

	Finished string `json:"finished"`
	// 快照创建的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。

	Started string `json:"started"`
	// 快照ID

	Id string `json:"id"`
	// 快照创建类型

	Type string `json:"type"`
	// 快照状态

	Status string `json:"status"`
}

func (o SnapshoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshoDetail struct{}"
	}

	return strings.Join([]string{"SnapshoDetail", string(data)}, " ")
}
