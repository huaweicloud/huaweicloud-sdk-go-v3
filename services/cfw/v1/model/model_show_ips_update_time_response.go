package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpsUpdateTimeResponse Response Object
type ShowIpsUpdateTimeResponse struct {
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
