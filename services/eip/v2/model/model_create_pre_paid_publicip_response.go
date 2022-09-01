package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePrePaidPublicipResponse struct {
	Publicip *PublicipCreateResp `json:"publicip,omitempty" xml:"publicip"`

	// 订单号（预付费场景返回该字段）
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 弹性公网IP的ID（预付费场景返回该字段）
	PublicipId     *string `json:"publicip_id,omitempty" xml:"publicip_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrePaidPublicipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrePaidPublicipResponse struct{}"
	}

	return strings.Join([]string{"CreatePrePaidPublicipResponse", string(data)}, " ")
}
