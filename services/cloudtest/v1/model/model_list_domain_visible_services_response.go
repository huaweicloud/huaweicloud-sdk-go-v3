package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainVisibleServicesResponse Response Object
type ListDomainVisibleServicesResponse struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value          *[]DomainVisibleServiceVo `json:"value,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListDomainVisibleServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainVisibleServicesResponse struct{}"
	}

	return strings.Join([]string{"ListDomainVisibleServicesResponse", string(data)}, " ")
}
