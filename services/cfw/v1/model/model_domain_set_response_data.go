package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetResponseData struct {

	// 域名组id
	Id *string `json:"id,omitempty"`

	// 域名组名称
	Name *string `json:"name,omitempty"`
}

func (o DomainSetResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetResponseData struct{}"
	}

	return strings.Join([]string{"DomainSetResponseData", string(data)}, " ")
}
