package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserUpdateAttribute 更细工作项时输入的用户信息
type UserUpdateAttribute struct {

	// **参数解释**： 用户ID，可通过[查询项目成员列表](ListProjectUsers.xml)接口获取，响应消息体中的**id**字段的值就是用户ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 用户名称。 **约束限制**： 当**id**属性有值时，优先使用**id**和项目成员进行匹配，匹配失败再按**name**匹配。 **取值范围**： 2~64个字符。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 用户昵称。 **约束限制**： 当**id**，**name**属性有值时，优先使用**id**，**name**和项目成员进行匹配，匹配失败再按**nick_name**匹配。 **取值范围**： 2~30个字符。 **默认取值**： 不涉及。
	NickName *string `json:"nick_name,omitempty"`
}

func (o UserUpdateAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserUpdateAttribute struct{}"
	}

	return strings.Join([]string{"UserUpdateAttribute", string(data)}, " ")
}
