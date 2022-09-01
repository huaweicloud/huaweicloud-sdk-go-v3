package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type OpenGaussInstanceActionRequest struct {
	ExpandCluster *OpenGaussExpandCluster `json:"expand_cluster,omitempty" xml:"expand_cluster"`

	EnlargeVolume *OpenGaussEnlargeVolume `json:"enlarge_volume,omitempty" xml:"enlarge_volume"`

	// 包周期实例时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。  true，表示自动从账户中支付。 false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *string `json:"is_auto_pay,omitempty" xml:"is_auto_pay"`
}

func (o OpenGaussInstanceActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussInstanceActionRequest struct{}"
	}

	return strings.Join([]string{"OpenGaussInstanceActionRequest", string(data)}, " ")
}
