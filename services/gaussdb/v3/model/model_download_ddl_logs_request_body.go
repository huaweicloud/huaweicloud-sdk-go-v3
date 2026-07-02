package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDdlLogsRequestBody 获取DDL日志下载链接请求体
type DownloadDdlLogsRequestBody struct {

	// **参数解释**：  DDL日志文件ID列表。  **约束限制**：  列表数量小于等于10。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	LogIds []string `json:"log_ids"`
}

func (o DownloadDdlLogsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDdlLogsRequestBody struct{}"
	}

	return strings.Join([]string{"DownloadDdlLogsRequestBody", string(data)}, " ")
}
