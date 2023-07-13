package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainItem struct {
	Metadata *MetaDomain `json:"metadata,omitempty"`
}

func (o DomainItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainItem struct{}"
	}

	return strings.Join([]string{"DomainItem", string(data)}, " ")
}
