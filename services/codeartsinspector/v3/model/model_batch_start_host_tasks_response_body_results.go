package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchStartHostTasksResponseBodyResults struct {

	// 扫描ID
	ScanId *string `json:"scan_id,omitempty"`

	// 主机ID
	HostId *string `json:"host_id,omitempty"`
}

func (o BatchStartHostTasksResponseBodyResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartHostTasksResponseBodyResults struct{}"
	}

	return strings.Join([]string{"BatchStartHostTasksResponseBodyResults", string(data)}, " ")
}
