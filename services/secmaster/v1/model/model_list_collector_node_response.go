package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorNodeResponse Response Object
type ListCollectorNodeResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 相关描述信息
	Records        *[]CollectorNode `json:"records,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListCollectorNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorNodeResponse struct{}"
	}

	return strings.Join([]string{"ListCollectorNodeResponse", string(data)}, " ")
}
