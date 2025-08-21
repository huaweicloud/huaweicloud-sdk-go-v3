package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomerIpsRespData **参数解释**： 更新自定义IPS规则响应数据 **取值范围**： 不涉及
type UpdateCustomerIpsRespData struct {

	// **参数解释**： 防火墙id **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o UpdateCustomerIpsRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomerIpsRespData struct{}"
	}

	return strings.Join([]string{"UpdateCustomerIpsRespData", string(data)}, " ")
}
