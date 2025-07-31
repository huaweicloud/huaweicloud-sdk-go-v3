package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostProtectionStatusResponse Response Object
type ShowHostProtectionStatusResponse struct {

	// 无风险的主机数
	NoRisk *int32 `json:"no_risk,omitempty"`

	// 有风险的主机数
	Risk *int32 `json:"risk,omitempty"`

	// 未防护的主机数
	NoProtect *int32 `json:"no_protect,omitempty"`

	// 主机总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHostProtectionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostProtectionStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowHostProtectionStatusResponse", string(data)}, " ")
}
