package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogDownloadResponse Response Object
type ShowSlowLogDownloadResponse struct {

	// 慢日志下载信息列表
	List           *[]SlowLogDownloadInfo `json:"list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowSlowLogDownloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDownloadResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDownloadResponse", string(data)}, " ")
}
