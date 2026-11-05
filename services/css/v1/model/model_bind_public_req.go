package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindPublicReq struct {
	Eip *BindPublicReqEip `json:"eip"`

	// 公网访问控制白名单。
	WhiteList *string `json:"white_list,omitempty"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o BindPublicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicReq struct{}"
	}

	return strings.Join([]string{"BindPublicReq", string(data)}, " ")
}
