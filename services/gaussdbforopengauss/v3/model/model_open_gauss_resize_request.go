package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规格变更时必填。
type OpenGaussResizeRequest struct {

	// 规格变更时选定的目标规格。新规格的资源规格编码。参考表1中GaussDB(for openGauss)的“规格编码”列内容获取。
	FlavorRef string `json:"flavor_ref" xml:"flavor_ref"`

	// 创建包周期实例时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。true，表示自动从账户中支付。false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty" xml:"is_auto_pay"`
}

func (o OpenGaussResizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussResizeRequest struct{}"
	}

	return strings.Join([]string{"OpenGaussResizeRequest", string(data)}, " ")
}
