package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubnetsRequest struct {

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 查询返回边缘子网列表数量。取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 站点ID。
	SiteId *string `json:"site_id,omitempty" xml:"site_id"`
}

func (o ListSubnetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetsRequest", string(data)}, " ")
}
