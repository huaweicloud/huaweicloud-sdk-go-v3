package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainClassificationItem struct {

	// DomainItem的总数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// DomainItem详细信息
	Items *[]DomainItem `json:"items,omitempty" xml:"items"`
}

func (o DomainClassificationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainClassificationItem struct{}"
	}

	return strings.Join([]string{"DomainClassificationItem", string(data)}, " ")
}
