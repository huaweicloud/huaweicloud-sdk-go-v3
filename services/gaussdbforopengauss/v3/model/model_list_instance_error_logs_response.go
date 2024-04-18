package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceErrorLogsResponse Response Object
type ListInstanceErrorLogsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 日志文件列表
	LogFiles       *[]InstanceLogFile `json:"log_files,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListInstanceErrorLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceErrorLogsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceErrorLogsResponse", string(data)}, " ")
}
