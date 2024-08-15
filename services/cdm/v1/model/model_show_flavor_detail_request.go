package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorDetailRequest Request Object
type ShowFlavorDetailRequest struct {

	// 规格ID，获取方法请参见[查询版本规格](ShowFlavors.xml)。
	FlavorId string `json:"flavor_id"`
}

func (o ShowFlavorDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowFlavorDetailRequest", string(data)}, " ")
}
