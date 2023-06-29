package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowErrorLogResponse Response Object
type ShowErrorLogResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	ErrorLogList   *[]ErrorLogList `json:"error_log_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowErrorLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowErrorLogResponse struct{}"
	}

	return strings.Join([]string{"ShowErrorLogResponse", string(data)}, " ")
}
