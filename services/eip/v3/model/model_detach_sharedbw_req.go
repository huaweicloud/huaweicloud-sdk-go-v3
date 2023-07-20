package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachSharedbwReq 共享带宽移出弹性公网IP请求参数
type DetachSharedbwReq struct {
	Publicip *DetachSharedbwReqPublicip `json:"publicip"`
}

func (o DetachSharedbwReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachSharedbwReq struct{}"
	}

	return strings.Join([]string{"DetachSharedbwReq", string(data)}, " ")
}
