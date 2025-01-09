package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleSiteInfo 边缘小站信息。
type SimpleSiteInfo struct {

	// 云桌面边缘小站id。
	Id *string `json:"id,omitempty"`

	// 云办公服务的状态。 - PREPARING：准备开通。 - SUBSCRIBING：开通中。 - SUBSCRIBED：已开通。 - SUBSCRIPTION_FAILED：开通失败。 - DEREGISTERING：销户中。 - DEREGISTRATION_FAILED：销户失败。 - CLOSED：已销户未开通。
	Status *string `json:"status,omitempty"`
}

func (o SimpleSiteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleSiteInfo struct{}"
	}

	return strings.Join([]string{"SimpleSiteInfo", string(data)}, " ")
}
