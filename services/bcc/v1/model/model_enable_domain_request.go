package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDomainRequest Request Object
type EnableDomainRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`
}

func (o EnableDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDomainRequest struct{}"
	}

	return strings.Join([]string{"EnableDomainRequest", string(data)}, " ")
}
