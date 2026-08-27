package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadOnlineDdlTaskLogResponse Response Object
type DownloadOnlineDdlTaskLogResponse struct {

	// **参数解释**：  下载实例无锁变更任务日志的链接。  **取值范围**：   不涉及。
	DownloadLink   *string `json:"download_link,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadOnlineDdlTaskLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadOnlineDdlTaskLogResponse struct{}"
	}

	return strings.Join([]string{"DownloadOnlineDdlTaskLogResponse", string(data)}, " ")
}
