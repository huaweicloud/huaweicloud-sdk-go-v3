package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoginCommonIpResponseInfo struct {

	// IP地址
	IpAddr *string `json:"ip_addr,omitempty"`

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 服务器列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o LoginCommonIpResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginCommonIpResponseInfo struct{}"
	}

	return strings.Join([]string{"LoginCommonIpResponseInfo", string(data)}, " ")
}
