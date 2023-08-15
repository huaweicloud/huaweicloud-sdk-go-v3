package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsErrorLogDetailsResponse Response Object
type ListLtsErrorLogDetailsResponse struct {

	// 错误日志列表。
	ErrorLogList   *[]LtsLogErrorDetail `json:"error_log_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListLtsErrorLogDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsErrorLogDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsErrorLogDetailsResponse", string(data)}, " ")
}
