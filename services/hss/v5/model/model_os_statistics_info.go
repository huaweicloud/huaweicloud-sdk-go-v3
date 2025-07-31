package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsStatisticsInfo 操作系统统计数据
type OsStatisticsInfo struct {

	// os_name
	OsName *string `json:"os_name,omitempty"`

	// os_type
	OsType *string `json:"os_type,omitempty"`

	// os_number
	Number *int32 `json:"number,omitempty"`
}

func (o OsStatisticsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsStatisticsInfo struct{}"
	}

	return strings.Join([]string{"OsStatisticsInfo", string(data)}, " ")
}
