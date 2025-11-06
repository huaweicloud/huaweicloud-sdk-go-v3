package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersonalRepositoryImportRecordsResponse Response Object
type ListPersonalRepositoryImportRecordsResponse struct {

	// 当前用户仓库导入任务列表
	Body *[]RepositoryImportRecordDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPersonalRepositoryImportRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonalRepositoryImportRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListPersonalRepositoryImportRecordsResponse", string(data)}, " ")
}
