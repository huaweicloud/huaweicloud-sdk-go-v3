package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerGroupItem 服务器组统计信息
type ServerGroupItem struct {

	// 服务器组ID
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 服务器组名称
	ServerGroupName *string `json:"server_group_name,omitempty"`

	// 服务器组分配的主机数量
	HostNum *int32 `json:"host_num,omitempty"`
}

func (o ServerGroupItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerGroupItem struct{}"
	}

	return strings.Join([]string{"ServerGroupItem", string(data)}, " ")
}
