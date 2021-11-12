package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectBugStaticsV4Response struct {
	// bug统计

	BugStatistics  *[]BugStatisticResponseV4 `json:"bug_statistics,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListProjectBugStaticsV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectBugStaticsV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectBugStaticsV4Response", string(data)}, " ")
}
