package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpNumRequirementResponse Response Object
type ShowIpNumRequirementResponse struct {

	// 消耗的IP个数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowIpNumRequirementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpNumRequirementResponse struct{}"
	}

	return strings.Join([]string{"ShowIpNumRequirementResponse", string(data)}, " ")
}
