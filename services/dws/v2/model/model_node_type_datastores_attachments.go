package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeTypeDatastoresAttachments **参数解释**： 内核版本附加信息。 **取值范围**： 不涉及。
type NodeTypeDatastoresAttachments struct {

	// **参数解释**： 内核版本支持的最小CN。 **取值范围**： 大于0的正整数。
	MinCn *int32 `json:"min_cn,omitempty"`

	// **参数解释**： 内核版本支持的最大CN。 **取值范围**： 大于0的正整数。
	MaxCn *int32 `json:"max_cn,omitempty"`
}

func (o NodeTypeDatastoresAttachments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeDatastoresAttachments struct{}"
	}

	return strings.Join([]string{"NodeTypeDatastoresAttachments", string(data)}, " ")
}
