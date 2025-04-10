package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspProtectStatisticsResponse Response Object
type ShowRaspProtectStatisticsResponse struct {

	// 防护主机数
	ProtectHostNum *int64 `json:"protect_host_num,omitempty"`

	// 防御篡改攻击数
	AntiTamperingNum *int64 `json:"anti_tampering_num,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o ShowRaspProtectStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRaspProtectStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowRaspProtectStatisticsResponse", string(data)}, " ")
}
