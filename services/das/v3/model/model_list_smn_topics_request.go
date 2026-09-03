package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmnTopicsRequest Request Object
type ListSmnTopicsRequest struct {

	// 偏移数值
	Offset *int32 `json:"offset,omitempty"`

	// 每页的数据量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSmnTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmnTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListSmnTopicsRequest", string(data)}, " ")
}
