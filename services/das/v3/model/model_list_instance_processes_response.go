package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceProcessesResponse Response Object
type ListInstanceProcessesResponse struct {

	// 同步时间
	DataSyncTime *int64 `json:"data_sync_time,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 数据列表
	Data *[]interface{} `json:"data,omitempty"`

	// 用户列表
	UserInfoList *[]string `json:"user_info_list,omitempty"`

	// 数据库列表
	DbInfoList *[]string `json:"db_info_list,omitempty"`

	// 来源IP列表
	HostInfoList *[]string `json:"host_info_list,omitempty"`

	// 状态列表
	StateInfoList *[]string `json:"state_info_list,omitempty"`

	// 命令列表
	CommandInfoList *[]string `json:"command_info_list,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o ListInstanceProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceProcessesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceProcessesResponse", string(data)}, " ")
}
