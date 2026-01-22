package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpsUpdateTimeResponse Response Object
type ShowIpsUpdateTimeResponse struct {

	// 查询ips规则细节时间数据
	Data           *[]IpsRuleUpdateTimeVo `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowIpsUpdateTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpsUpdateTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowIpsUpdateTimeResponse", string(data)}, " ")
}
