package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductExtendResp **参数解释**： 当前产品信息的扩展信息，例如存储、IO等。 **取值范围**： 不涉及。
type ProductExtendResp struct {

	// **参数解释**： 集群维度规格扩展信息编号。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 集群维度规格扩展信息值。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o ProductExtendResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductExtendResp struct{}"
	}

	return strings.Join([]string{"ProductExtendResp", string(data)}, " ")
}
