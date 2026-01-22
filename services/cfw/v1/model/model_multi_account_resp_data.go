package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultiAccountRespData **参数解释**： 多账号响应信息 **取值范围**： 不涉及
type MultiAccountRespData struct {

	// **参数解释**： 防火墙ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 防火墙名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o MultiAccountRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiAccountRespData struct{}"
	}

	return strings.Join([]string{"MultiAccountRespData", string(data)}, " ")
}
