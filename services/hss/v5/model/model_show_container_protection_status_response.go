package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowContainerProtectionStatusResponse Response Object
type ShowContainerProtectionStatusResponse struct {

	// 无风险的容器数
	NoRisk *int32 `json:"no_risk,omitempty"`

	// 有风险的容器数
	Risk *int32 `json:"risk,omitempty"`

	// 未防护的容器数
	NoProtect *int32 `json:"no_protect,omitempty"`

	// 容器总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowContainerProtectionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowContainerProtectionStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowContainerProtectionStatusResponse", string(data)}, " ")
}
