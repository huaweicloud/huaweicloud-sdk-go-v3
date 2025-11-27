package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectCollectRequest Request Object
type ListEnterpriseProjectCollectRequest struct {

	// **参数解释：** 唯一标识id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 分页参数，上一页请求最后一个id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释：** 分页查询每页显示的条目数量。 **约束限制：** 不涉及。 **取值范围：** 自定义，在1~500范围。 **默认取值：** 不涉及。
	Limit int32 `json:"limit"`
}

func (o ListEnterpriseProjectCollectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectCollectRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectCollectRequest", string(data)}, " ")
}
