package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespVhosts struct {

	// **参数解释**： Vhost名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o ShowCeshierarchyRespVhosts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespVhosts struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespVhosts", string(data)}, " ")
}
