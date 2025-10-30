package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginCommonIpRequestInfo 设置白名单列表
type ModifyLoginCommonIpRequestInfo struct {

	// **参数解释**： 登录IP或登录网段,登录网段由IP地址和掩码组成,以'/'连接。 **取值范围**： 字符长度1-128位
	IpAddr string `json:"ip_addr"`

	// **参数解释**： 服务器ID列表,不可为NULL,删除常用登录IP时,需要将服务器ID列表置为空列表[]。 **取值范围**： 最小值0，最大值200
	HostIdList []string `json:"host_id_list"`
}

func (o ModifyLoginCommonIpRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginCommonIpRequestInfo struct{}"
	}

	return strings.Join([]string{"ModifyLoginCommonIpRequestInfo", string(data)}, " ")
}
