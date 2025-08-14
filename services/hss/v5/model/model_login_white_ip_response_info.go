package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoginWhiteIpResponseInfo struct {

	// 白名单启用状态，包含如下：   - true：已启用   - false：已禁用
	Enabled *bool `json:"enabled,omitempty"`

	// 白名单IP或IP网段，IP网段由IP地址和掩码组成，以‘/’连接。
	WhiteIp *string `json:"white_ip,omitempty"`

	// 服务器ID总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 服务器ID列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o LoginWhiteIpResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginWhiteIpResponseInfo struct{}"
	}

	return strings.Join([]string{"LoginWhiteIpResponseInfo", string(data)}, " ")
}
