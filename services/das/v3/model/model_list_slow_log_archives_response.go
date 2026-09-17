package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogArchivesResponse Response Object
type ListSlowLogArchivesResponse struct {

	// 慢日志归档文件列表
	ArchiveList *[]SlowLogArchiveDto `json:"archive_list,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogArchivesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogArchivesResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogArchivesResponse", string(data)}, " ")
}
