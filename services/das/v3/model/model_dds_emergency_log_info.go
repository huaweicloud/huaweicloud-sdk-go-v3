package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DdsEmergencyLogInfo DDSEmergencyLogInfo对象
type DdsEmergencyLogInfo struct {

	// ID
	Id *int64 `json:"id,omitempty"`

	// 执行SQL
	ExecuteSql *string `json:"execute_sql,omitempty"`

	// 创建时间
	CreateAt *int64 `json:"create_at,omitempty"`
}

func (o DdsEmergencyLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DdsEmergencyLogInfo struct{}"
	}

	return strings.Join([]string{"DdsEmergencyLogInfo", string(data)}, " ")
}
