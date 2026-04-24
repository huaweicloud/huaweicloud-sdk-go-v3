package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloseDomainNewReq 关闭域控的请求。
type CloseDomainNewReq struct {

	// 域id。
	DomainId string `json:"domain_id"`

	AuthType *DomainType `json:"auth_type"`
}

func (o CloseDomainNewReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloseDomainNewReq struct{}"
	}

	return strings.Join([]string{"CloseDomainNewReq", string(data)}, " ")
}
