package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcbResourceByTagRequest Request Object
type ListGcbResourceByTagRequest struct {

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量。
	Offset *int32 `json:"offset,omitempty"`

	Body *QueryResourceByTagRequestBody `json:"body,omitempty"`
}

func (o ListGcbResourceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcbResourceByTagRequest struct{}"
	}

	return strings.Join([]string{"ListGcbResourceByTagRequest", string(data)}, " ")
}
