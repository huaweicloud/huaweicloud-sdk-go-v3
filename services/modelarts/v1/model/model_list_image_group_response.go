package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageGroupResponse Response Object
type ListImageGroupResponse struct {

	// **参数解释**：当前页数。 **取值范围**：正整数。
	Current *int32 `json:"current,omitempty"`

	// **参数解释**：镜像信息概览数据。
	Data *[]ImageGroup `json:"data,omitempty"`

	// **参数解释**：总的页数。 **取值范围**：正整数。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释**：每一页的数量。 **取值范围**：正整数。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：总的记录数量。 **取值范围**：非负整数。
	Total *int64 `json:"total,omitempty"`

	// **参数解释**：当前账号是否存在swr企业版镜像。 **约束限制**：true或false。 **取值范围**：布尔类型 **默认取值**：false。
	IsSwrEnterprise *bool `json:"is_swr_enterprise,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListImageGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageGroupResponse struct{}"
	}

	return strings.Join([]string{"ListImageGroupResponse", string(data)}, " ")
}
