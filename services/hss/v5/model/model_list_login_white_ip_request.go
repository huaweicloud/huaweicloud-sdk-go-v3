package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginWhiteIpRequest Request Object
type ListLoginWhiteIpRequest struct {

	// **参数解释**: 主机所属的企业项目ID。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。 **约束限制**: 开通企业项目功能后才需要配置企业项目。 **取值范围**: 字符长度1-256位 **默认取值**: 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 白名单IP/IP网段,IP网段由IP地址和掩码组成,以'/'连接。
	WhiteIp *string `json:"white_ip,omitempty"`
}

func (o ListLoginWhiteIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginWhiteIpRequest struct{}"
	}

	return strings.Join([]string{"ListLoginWhiteIpRequest", string(data)}, " ")
}
