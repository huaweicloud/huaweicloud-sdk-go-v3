package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataIoTime struct {

	// **参数解释**: 总耗时（单位：微秒）。 **取值范围**: 不涉及。
	AllTime int64 `json:"all_time"`

	DataIoTimeDetails *EventTimeInfo `json:"data_io_time_details"`
}

func (o DataIoTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataIoTime struct{}"
	}

	return strings.Join([]string{"DataIoTime", string(data)}, " ")
}
