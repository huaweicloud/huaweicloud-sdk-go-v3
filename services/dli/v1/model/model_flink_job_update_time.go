package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobUpdateTime 作业更新信息。
type FlinkJobUpdateTime struct {

	// 作业更新时间, 毫秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o FlinkJobUpdateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobUpdateTime struct{}"
	}

	return strings.Join([]string{"FlinkJobUpdateTime", string(data)}, " ")
}
