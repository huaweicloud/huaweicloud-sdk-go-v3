package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CnameStatus struct {

	// 域名cname状态
	Status int32 `json:"status"`

	// 加速域名
	DomainName string `json:"domain_name"`
}

func (o CnameStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CnameStatus struct{}"
	}

	return strings.Join([]string{"CnameStatus", string(data)}, " ")
}
