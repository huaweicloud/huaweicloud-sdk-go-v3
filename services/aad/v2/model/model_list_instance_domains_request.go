package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceDomainsRequest Request Object
type ListInstanceDomainsRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 限制条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInstanceDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceDomainsRequest", string(data)}, " ")
}
