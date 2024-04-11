package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSlowLogDownloadResponse Response Object
type CreateSlowLogDownloadResponse struct {

	// 慢日志下载信息列表
	List           *[]SlowLogDownloadInfo `json:"list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateSlowLogDownloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSlowLogDownloadResponse struct{}"
	}

	return strings.Join([]string{"CreateSlowLogDownloadResponse", string(data)}, " ")
}
