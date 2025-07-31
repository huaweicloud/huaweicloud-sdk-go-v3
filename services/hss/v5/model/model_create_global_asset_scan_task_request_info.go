package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalAssetScanTaskRequestInfo 下发任务的列表，三种参数选择其中一种即可，有多个参数则优先级顺序为all_hosts，server_group，host_ids
type CreateGlobalAssetScanTaskRequestInfo struct {

	// 下发任务的主机列表
	HostIds *[]string `json:"host_ids,omitempty"`

	// 下发任务的主机组列表
	ServerGroup *[]string `json:"server_group,omitempty"`

	// 下发全部主机的扫描
	AllHosts bool `json:"all_hosts"`
}

func (o CreateGlobalAssetScanTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalAssetScanTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateGlobalAssetScanTaskRequestInfo", string(data)}, " ")
}
