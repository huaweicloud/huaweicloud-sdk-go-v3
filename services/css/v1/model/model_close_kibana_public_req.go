package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloseKibanaPublicReq struct {

	// 带宽。单位：Mbit/s
	EipSize *int32 `json:"eip_size,omitempty"`

	ElbWhiteList *StartKibanaPublicReqElbWhitelist `json:"elb_white_list,omitempty"`
}

func (o CloseKibanaPublicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseKibanaPublicReq struct{}"
	}

	return strings.Join([]string{"CloseKibanaPublicReq", string(data)}, " ")
}
