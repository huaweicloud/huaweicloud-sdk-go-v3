package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobLogResponse Response Object
type ShowJobLogResponse struct {

	// 作业日志条数
	Count *int32 `json:"count,omitempty"`

	// 作业日志内容列表
	Logs *[]LogContentDto `json:"logs,omitempty"`

	// 作业日志存储链接
	LogStorageLink *string `json:"log_storage_link,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobLogResponse struct{}"
	}

	return strings.Join([]string{"ShowJobLogResponse", string(data)}, " ")
}
