package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorDetailResponse Response Object
type ShowFlavorDetailResponse struct {

	// 规格名称
	Name *string `json:"name,omitempty"`

	// 规格ID
	StrId *string `json:"str_id,omitempty"`

	// 规格详细列表
	FlavorDetail   *[]FlavorAttribute `json:"flavor_detail,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowFlavorDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorDetailResponse", string(data)}, " ")
}
