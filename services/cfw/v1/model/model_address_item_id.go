package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressItemId struct {

	// **参数解释**： 地址组成员ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： IP地址 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o AddressItemId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressItemId struct{}"
	}

	return strings.Join([]string{"AddressItemId", string(data)}, " ")
}
