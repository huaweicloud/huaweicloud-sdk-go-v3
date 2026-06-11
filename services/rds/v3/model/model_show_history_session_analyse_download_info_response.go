package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistorySessionAnalyseDownloadInfoResponse Response Object
type ShowHistorySessionAnalyseDownloadInfoResponse struct {

	// 记录个数
	Count *int32 `json:"count,omitempty"`

	// 下载信息
	List           *[]DownloadInfo `json:"list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHistorySessionAnalyseDownloadInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistorySessionAnalyseDownloadInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowHistorySessionAnalyseDownloadInfoResponse", string(data)}, " ")
}
