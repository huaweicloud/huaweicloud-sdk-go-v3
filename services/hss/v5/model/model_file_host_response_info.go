package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileHostResponseInfo 云服务器具体变更信息
type FileHostResponseInfo struct {

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 变更总数
	ChangeTotalNum *int32 `json:"change_total_num,omitempty"`

	// 变更文件
	ChangeFileNum *int32 `json:"change_file_num,omitempty"`

	// 变更注册表
	ChangeRegistryNum *int32 `json:"change_registry_num,omitempty"`

	// 最后变更时间
	LatestTime *int64 `json:"latest_time,omitempty"`
}

func (o FileHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileHostResponseInfo struct{}"
	}

	return strings.Join([]string{"FileHostResponseInfo", string(data)}, " ")
}
