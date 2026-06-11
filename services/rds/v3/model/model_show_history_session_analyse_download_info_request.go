package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistorySessionAnalyseDownloadInfoRequest Request Object
type ShowHistorySessionAnalyseDownloadInfoRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 类型。 session    表示下载历史会话 wait_event    表示下载历史等待事件 top_sql    表示下载历史top sql
	Type string `json:"type"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowHistorySessionAnalyseDownloadInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistorySessionAnalyseDownloadInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowHistorySessionAnalyseDownloadInfoRequest", string(data)}, " ")
}
