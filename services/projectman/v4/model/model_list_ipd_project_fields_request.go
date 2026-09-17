package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpdProjectFieldsRequest Request Object
type ListIpdProjectFieldsRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// **参数解释**： 关键字搜索，支持标题、字段类型、创建人搜索。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Keyword *string `json:"keyword,omitempty"`

	// **参数解释**： 分页起始，从0开始，为limit整数倍。 **约束限制**： 取值为limit的倍数。 **取值范围**： ≥ 0 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小。 **约束限制**： 不涉及 **取值范围**： ≥ 1 **默认取值**： 20
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIpdProjectFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdProjectFieldsRequest struct{}"
	}

	return strings.Join([]string{"ListIpdProjectFieldsRequest", string(data)}, " ")
}
