package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBuildRealTimeLogResponse Response Object
type DownloadBuildRealTimeLogResponse struct {
	Result *OctopusV3LogResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadBuildRealTimeLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBuildRealTimeLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadBuildRealTimeLogResponse", string(data)}, " ")
}
