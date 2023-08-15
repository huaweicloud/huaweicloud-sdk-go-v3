package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataArtsStudioInstancesRequest Request Object
type ListDataArtsStudioInstancesRequest struct {

	// 分页记录数。默认20
	Limit *int32 `json:"limit,omitempty"`

	// 分页偏移量。默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDataArtsStudioInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataArtsStudioInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListDataArtsStudioInstancesRequest", string(data)}, " ")
}
