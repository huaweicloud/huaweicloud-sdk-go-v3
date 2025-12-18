package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportInstanceInfosRequestBody struct {

	// **参数解释**:   实例id列表。 **约束限制**:   不涉及。 **取值范围**:   不涉及 **默认取值**:   不涉及。
	InstanceList *[]string `json:"instance_list,omitempty"`

	// **参数解释**:   导出字段列表。 **约束限制**:   不涉及。 **取值范围**:   - name：实例名称   - id：实例ID   - alias：实例备注   - editionMode：产品类型   - haModel：实例类型   - deployMode：部署形态   - engine：数据库引擎版本   - hotfixVersions：已升级热补丁   - instanceStatus：实例状态   - payModel：计费模式   - targetEngineVersion：目标版本   - flavor：性能规格   - availableZones：可用区   - privateIp：内网地址   - dnsName：DNS   - ipv6：IPv6地址   - dbPort：数据库端口   - publicIp：弹性公网IP   - createAt：创建时间   - volumeType：存储空间类型   - volumeSize：存储空间大小   - vpcName：虚拟私有云名称   - vpcId：虚拟私有云ID   - securityGroupName：安全组   - enterpriseProjectName：企业项目  **默认取值**:   不涉及。
	UserDefinedColumns []string `json:"user_defined_columns"`

	// **参数解释**:   时区。 **约束限制**:   不涉及。 **取值范围**:   - +08:00 **默认取值**:   +08:00
	TimeZone *string `json:"time_zone,omitempty"`

	// **参数解释**:   语言。 **约束限制**:   不涉及。 **取值范围**:   - zh-cn    - en-us  **默认取值**:   zh-cn
	Language *string `json:"language,omitempty"`
}

func (o ExportInstanceInfosRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceInfosRequestBody struct{}"
	}

	return strings.Join([]string{"ExportInstanceInfosRequestBody", string(data)}, " ")
}
