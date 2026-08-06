package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeyPolicyRequest Request Object
type ListKeyPolicyRequest struct {

	// **参数解释：** 密钥空间ID **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyspaceId string `json:"keyspace_id"`

	// **参数解释：** 分页参数，每一页显示的记录数 **约束限制：** 数字类型 **取值范围：** 1-100 **默认取值：** 50
	Limit *string `json:"limit,omitempty"`

	// **参数解释：** 分页参数，下一页的标志 **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	Marker *string `json:"marker,omitempty"`
}

func (o ListKeyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListKeyPolicyRequest", string(data)}, " ")
}
