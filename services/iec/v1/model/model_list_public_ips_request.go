package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPublicIpsRequest struct {

	// 查询返回弹性IP列表数量。取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 站点ID。
	SiteId *string `json:"site_id,omitempty" xml:"site_id"`

	// 端口ID
	PortId *string `json:"port_id,omitempty" xml:"port_id"`
}

func (o ListPublicIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicIpsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicIpsRequest", string(data)}, " ")
}
