package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartKibanaPublicReq struct {

	// 带宽。单位：Mbit/s
	EipSize int32 `json:"eip_size"`

	ElbWhiteList *StartKibanaPublicReqElbWhitelist `json:"elb_white_list,omitempty"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o StartKibanaPublicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKibanaPublicReq struct{}"
	}

	return strings.Join([]string{"StartKibanaPublicReq", string(data)}, " ")
}
