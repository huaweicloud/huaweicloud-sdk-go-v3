package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorDetailRequest Request Object
type ShowFlavorDetailRequest struct {

	// 参数解释： 规格id。 约束限制： 不涉及 取值范围： 实例规格列表接口返回的id。 默认取值： 不涉及
	FlavorId string `json:"flavor_id"`
}

func (o ShowFlavorDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowFlavorDetailRequest", string(data)}, " ")
}
