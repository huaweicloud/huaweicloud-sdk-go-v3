package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowlogResponse Response Object
type ListSlowlogResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 慢日志列表
	Slowlogs       *[]SlowlogItem `json:"slowlogs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSlowlogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogResponse struct{}"
	}

	return strings.Join([]string{"ListSlowlogResponse", string(data)}, " ")
}
