package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserAuthorityRequest Request Object
type ExportUserAuthorityRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 数据库用户/角色名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10000
	Name string `json:"name"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ExportUserAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ExportUserAuthorityRequest", string(data)}, " ")
}
