package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCssClusterRequest Request Object
type ListCssClusterRequest struct {

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListCssClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCssClusterRequest struct{}"
	}

	return strings.Join([]string{"ListCssClusterRequest", string(data)}, " ")
}
