package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDdlLogsResponse Response Object
type DownloadDdlLogsResponse struct {

	// **参数解释**：  每个日志文件的下载链接详情。  **取值范围**：  不涉及。
	DownloadFiles  *[]DownLoadFileInfoItem `json:"download_files,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o DownloadDdlLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDdlLogsResponse struct{}"
	}

	return strings.Join([]string{"DownloadDdlLogsResponse", string(data)}, " ")
}
