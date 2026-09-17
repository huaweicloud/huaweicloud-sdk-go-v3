package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogArchiveLinkResponse Response Object
type ShowSlowLogArchiveLinkResponse struct {

	// 慢日志归档文件下载地址
	DownloadLink   *string `json:"download_link,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSlowLogArchiveLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogArchiveLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogArchiveLinkResponse", string(data)}, " ")
}
