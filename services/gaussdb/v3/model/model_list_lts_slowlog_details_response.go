package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsSlowlogDetailsResponse Response Object
type ListLtsSlowlogDetailsResponse struct {

	// 慢日志列表。
	SlowLogList    *[]LtsLogSlowDetail `json:"slow_log_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListLtsSlowlogDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowlogDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsSlowlogDetailsResponse", string(data)}, " ")
}
