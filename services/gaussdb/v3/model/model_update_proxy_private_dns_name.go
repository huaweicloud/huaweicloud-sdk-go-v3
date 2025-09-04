package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProxyPrivateDnsName struct {

	// **参数解释**：  新的dns名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	NewDnsName string `json:"new_dns_name"`

	// **参数解释**：            虚拟私有云ID，获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html)。  **约束限制**：   不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	VpcId string `json:"vpc_id"`
}

func (o UpdateProxyPrivateDnsName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyPrivateDnsName struct{}"
	}

	return strings.Join([]string{"UpdateProxyPrivateDnsName", string(data)}, " ")
}
