package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainRequest Request Object
type ListDomainRequest struct {

	// 限制条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainRequest struct{}"
	}

	return strings.Join([]string{"ListDomainRequest", string(data)}, " ")
}
