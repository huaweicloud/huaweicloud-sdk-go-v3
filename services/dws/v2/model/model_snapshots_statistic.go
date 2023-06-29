package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotsStatistic 快照统计信息。
type SnapshotsStatistic struct {

	// 资源统计信息名称。 - storage.free：免费容量。 - storage.paid：付费容量。 - storage.used：已用容量。
	Name string `json:"name"`

	// 资源统计信息值。
	Value float32 `json:"value"`

	// 资源统计信息单位。
	Unit string `json:"unit"`
}

func (o SnapshotsStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotsStatistic struct{}"
	}

	return strings.Join([]string{"SnapshotsStatistic", string(data)}, " ")
}
