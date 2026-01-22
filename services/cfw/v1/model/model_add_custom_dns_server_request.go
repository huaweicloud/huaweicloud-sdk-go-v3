package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCustomDnsServerRequest Request Object
type AddCustomDnsServerRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： DNS服务器IP，可通过查询dns服务器列表接口查询获得，通过返回值中的data.server_ip（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServerIp string `json:"server_ip"`
}

func (o AddCustomDnsServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCustomDnsServerRequest struct{}"
	}

	return strings.Join([]string{"AddCustomDnsServerRequest", string(data)}, " ")
}
