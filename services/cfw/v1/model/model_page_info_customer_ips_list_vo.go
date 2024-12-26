package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoCustomerIpsListVo struct {
	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Records *[]CustomerIpsListVo `json:"records,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

func (o PageInfoCustomerIpsListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoCustomerIpsListVo struct{}"
	}

	return strings.Join([]string{"PageInfoCustomerIpsListVo", string(data)}, " ")
}
