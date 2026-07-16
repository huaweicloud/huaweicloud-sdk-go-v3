package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerOsRequest Lite Server操作系统镜像相关操作请求体。
type ServerOsRequest struct {

	// **参数解释**：用于登录服务器密码。注意弹性云服务器和裸金属服务器admin_pass和key_pair_name必须二选一，超节点不支持admin_pass。 **约束限制**：密码规则： - 长度为8至26个 - 至少包含大写字母、小写字母、数字及特殊符号(!@%-_=+[{}]:,./?)中的3种 - 不能与用户名或倒序的用户名相同 - 不能包含root或administrator及其逆序 **取值范围**：不涉及 **默认取值**：不涉及
	AdminPass *string `json:"admin_pass,omitempty"`

	// **参数解释**：服务器登录密钥对名称。注意admin_pass和key_pair_name必须二选一。 **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及
	KeyPairName *string `json:"key_pair_name,omitempty"`

	// **参数解释**：镜像ID，切换操作系统场景，该参数必填。 **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 创建云服务器过程中待注入实例自定义数据。支持注入文本、文本文件。 示例： base64编码前： * Linux服务器： ```bash #!/bin/bash echo user_test > /home/user.txt ``` base64编码后： * Linux服务器： ```bash IyEvYmluL2Jhc2gKZWNobyB1c2VyX3Rlc3QgPiAvaG9tZS91c2VyLnR4dA== ``` 了解更多实例自定义数据注入请参考[[用户数据注入](https://support.huaweicloud.com/usermanual-ecs/zh-cn_topic_0032380449.html)](tag:hc)[[用户数据注入](https://support.huaweicloud.com/intl/zh-cn/usermanual-ecs/zh-cn_topic_0032380449.html)][ECS服务“通过实例自定义数据配置ECS实例”章节](tag:fcs,hcso)。 用户需明确user_data的使用效果，可能产生的影响和风险由用户自行承担。 **约束限制**： - user_data的值为base64编码之后的内容。 - 注入内容（编码之前的内容）最大长度为32K。  **取值范围**：不涉及。 **默认取值**：不涉及。
	UserData *string `json:"user_data,omitempty"`
}

func (o ServerOsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerOsRequest struct{}"
	}

	return strings.Join([]string{"ServerOsRequest", string(data)}, " ")
}
