package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePrePaidPublicipResponse struct {
	Publicip *PublicipCreateResp `json:"publicip,omitempty"`
	// 订单号（预付费场景返回该字段）

	OrderId *string `json:"order_id,omitempty"`
	// 弹性公网IP的ID（预付费场景返回该字段）

	PublicipId     *string `json:"publicip_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrePaidPublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrePaidPublicipResponse struct{}"
	}

	return strings.Join([]string{"CreatePrePaidPublicipResponse", string(data)}, " ")
}
