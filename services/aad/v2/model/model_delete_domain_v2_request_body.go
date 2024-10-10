package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDomainV2RequestBody struct {

	// 域名id列表
	DomainId []string `json:"domain_id"`
}

func (o DeleteDomainV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainV2RequestBody struct{}"
	}

	return strings.Join([]string{"DeleteDomainV2RequestBody", string(data)}, " ")
}
