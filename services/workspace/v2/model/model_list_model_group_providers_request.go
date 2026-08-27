package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelGroupProvidersRequest Request Object
type ListModelGroupProvidersRequest struct {

	// 模型组id。
	GroupId string `json:"group_id"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小，默认20。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListModelGroupProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelGroupProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListModelGroupProvidersRequest", string(data)}, " ")
}
