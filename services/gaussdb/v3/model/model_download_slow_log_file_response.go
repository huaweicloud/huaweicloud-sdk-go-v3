package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSlowLogFileResponse Response Object
type DownloadSlowLogFileResponse struct {

	// 慢日志下载链接列表。
	List *[]DownloadSlowLogFileItem `json:"list,omitempty"`

	// 条数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DownloadSlowLogFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowLogFileResponse struct{}"
	}

	return strings.Join([]string{"DownloadSlowLogFileResponse", string(data)}, " ")
}
