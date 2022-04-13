package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSlowLogResponse struct {
	// DDM慢sql日志条数。

	TotalRecord *int32 `json:"totalRecord,omitempty"`
	// DDM慢sql日志信息列表的集合。

	SlowLogList    *[]SlowLogList `json:"slowLogList,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogResponse", string(data)}, " ")
}
