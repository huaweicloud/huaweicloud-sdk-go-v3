package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagerWorkSpacesRequest Request Object
type ListManagerWorkSpacesRequest struct {

	// DataArtsStudio实例id
	InstanceId string `json:"instance_id"`

	// 分页记录数，默认20
	Limit *int32 `json:"limit,omitempty"`

	// 数据偏移量。默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListManagerWorkSpacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagerWorkSpacesRequest struct{}"
	}

	return strings.Join([]string{"ListManagerWorkSpacesRequest", string(data)}, " ")
}
