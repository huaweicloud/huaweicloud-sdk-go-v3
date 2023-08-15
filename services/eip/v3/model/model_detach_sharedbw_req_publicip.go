package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachSharedbwReqPublicip 共享带宽移出弹性公网IP请求对象
type DetachSharedbwReqPublicip struct {
	Bandwidth *DetachSharedbwDict `json:"bandwidth,omitempty"`
}

func (o DetachSharedbwReqPublicip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachSharedbwReqPublicip struct{}"
	}

	return strings.Join([]string{"DetachSharedbwReqPublicip", string(data)}, " ")
}
