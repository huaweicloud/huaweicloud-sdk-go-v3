package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainStatsResponse Response Object
type ShowDomainStatsResponse struct {

	// 按指定的分组方式组织的数据
	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsResponse", string(data)}, " ")
}
