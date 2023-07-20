package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMeteringResponse Response Object
type ShowMeteringResponse struct {

	// 资源信息列表。
	ProductInfoList *[]ProductInfo `json:"product_info_list,omitempty"`
	HttpStatusCode  int            `json:"-"`
}

func (o ShowMeteringResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeteringResponse struct{}"
	}

	return strings.Join([]string{"ShowMeteringResponse", string(data)}, " ")
}
