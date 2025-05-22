package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadRealTimeLogResponse Response Object
type DownloadRealTimeLogResponse struct {
	Result *RealTimeLogResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadRealTimeLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadRealTimeLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadRealTimeLogResponse", string(data)}, " ")
}
