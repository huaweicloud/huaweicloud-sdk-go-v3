package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerHyperScaleUpRequest struct {

	// **参数解释**：服务器规格名称。 **约束限制**：不涉及。 **取值范围**：长度为1至128个字符，只能包含字母和数字及点。 **默认取值**：不涉及。
	Flavor string `json:"flavor"`

	RootVolume *EvsVolume `json:"root_volume,omitempty"`

	DataVolume *ServerDataVolume `json:"data_volume,omitempty"`

	// **参数解释**：服务器镜像ID。 **约束限制**：不涉及。 **取值范围**：长度为36个字符，符合UUID格式。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**： 创建云服务器过程中待注入实例自定义数据。支持注入文本、文本文件。 示例： base64编码前： * Linux服务器： ```bash #!/bin/bash echo user_test > /home/user.txt ``` base64编码后： * Linux服务器： ```bash IyEvYmluL2Jhc2gKZWNobyB1c2VyX3Rlc3QgPiAvaG9tZS91c2VyLnR4dA== ``` 了解更多实例自定义数据注入请参考[[用户数据注入](https://support.huaweicloud.com/usermanual-ecs/zh-cn_topic_0032380449.html)](tag:hc)[[用户数据注入](https://support.huaweicloud.com/intl/zh-cn/usermanual-ecs/zh-cn_topic_0032380449.html)][ECS服务“通过实例自定义数据配置ECS实例”章节](tag:fcs,hcso)。 用户需明确user_data的使用效果，可能产生的影响和风险由用户自行承担。 **约束限制**： - user_data的值为base64编码之后的内容。 - 注入内容（编码之前的内容）最大长度为32K。  **取值范围**：不涉及。 **默认取值**：不涉及。
	Userdata *string `json:"userdata,omitempty"`

	// **参数解释**：服务器登录密钥对名称。admin_pass和key_pair_name必须二选一。注意超节点扩容仅支持使用密钥对创建。 **约束限制**：admin_pass和key_pair_name不能同时存在。 **取值范围**：长度为1至64个字符，只能包含字母、数字、中划线、下划线和点。 **默认取值**：不涉及。
	KeyPairName *string `json:"key_pair_name,omitempty"`
}

func (o ServerHyperScaleUpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerHyperScaleUpRequest struct{}"
	}

	return strings.Join([]string{"ServerHyperScaleUpRequest", string(data)}, " ")
}
