package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventLogRequest Request Object
type ListEventLogRequest struct {

	// 页数
	Page *int32 `json:"page,omitempty"`

	// 单页显示数
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListEventLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventLogRequest struct{}"
	}

	return strings.Join([]string{"ListEventLogRequest", string(data)}, " ")
}
