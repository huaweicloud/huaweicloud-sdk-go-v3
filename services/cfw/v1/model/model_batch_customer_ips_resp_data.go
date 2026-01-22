package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCustomerIpsRespData **参数解释**： 批量操作自定义IPS规则响应数据 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**：   不涉及
type BatchCustomerIpsRespData struct {

	// **参数解释**： 防火墙id **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o BatchCustomerIpsRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCustomerIpsRespData struct{}"
	}

	return strings.Join([]string{"BatchCustomerIpsRespData", string(data)}, " ")
}
