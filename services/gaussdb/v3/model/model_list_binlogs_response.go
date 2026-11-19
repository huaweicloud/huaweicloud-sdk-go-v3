package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogsResponse Response Object
type ListBinlogsResponse struct {

	// **参数解释**：  Binlog日志文件总数。   **取值范围**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  Binlog日志文件列表。
	Binlogs        *[]BinlogInfo `json:"binlogs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListBinlogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogsResponse struct{}"
	}

	return strings.Join([]string{"ListBinlogsResponse", string(data)}, " ")
}
