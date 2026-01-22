package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteServiceItemResponseBodyData struct {

	// **参数解释**： 服务组成员ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 源、目的端口 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o DeleteServiceItemResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemResponseBodyData struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemResponseBodyData", string(data)}, " ")
}
