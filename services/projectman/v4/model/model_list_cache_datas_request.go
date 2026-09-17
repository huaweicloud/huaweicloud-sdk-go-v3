package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCacheDatasRequest struct {

	// **参数解释**： 项目的32位uuid，项目唯一标识，通过[查询项目列表](ListProjectsV4.xml)接口获取，响应消息体中的**project_id**字段的值就是项目ID。 **约束限制**： 32位的数字和字母组成的字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectUUId *string `json:"projectUUId,omitempty"`

	// **参数解释：** 字段类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： backlog。
	Type *string `json:"type,omitempty"`
}

func (o ListCacheDatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCacheDatasRequest struct{}"
	}

	return strings.Join([]string{"ListCacheDatasRequest", string(data)}, " ")
}
