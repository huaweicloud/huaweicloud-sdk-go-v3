package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileHostResponseInfo 云服务器具体变更信息
type FileHostResponseInfo struct {

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 文件变更与注册表变更总数量 **取值范围**： 最小值0，最大值2147483647
	ChangeTotalNum *int32 `json:"change_total_num,omitempty"`

	// **参数解释**： 文件变更总数量 **取值范围**： 最小值0，最大值2147483647
	ChangeFileNum *int32 `json:"change_file_num,omitempty"`

	// **参数解释**： 注册表变更总数量 **取值范围**： 最小值0，最大值2147483647
	ChangeRegistryNum *int32 `json:"change_registry_num,omitempty"`

	// **参数解释**： 最后一次文件/注册表变更时间 **取值范围**： 非负长整数，时间格式：13位毫秒级时间戳（UTC时区，从1970-01-01 00:00:00开始计算），单位：ms
	LatestTime *int64 `json:"latest_time,omitempty"`
}

func (o FileHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileHostResponseInfo struct{}"
	}

	return strings.Join([]string{"FileHostResponseInfo", string(data)}, " ")
}
