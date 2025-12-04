package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPackLogInfosResponse Response Object
type ListPackLogInfosResponse struct {

	// **参数解释**：  binlog合并下载文件数量。  **约束限制**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  binlog合并下载文件信息。  **约束限制**：  不涉及。
	PackLogInfos   *[]PackLogInfo `json:"pack_log_infos,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPackLogInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPackLogInfosResponse struct{}"
	}

	return strings.Join([]string{"ListPackLogInfosResponse", string(data)}, " ")
}
