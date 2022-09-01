package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSlowLogFileResponse struct {
	List *[]SlowLogFile `json:"list,omitempty" xml:"list"`

	// 文件总数。
	TotalCount     *int64 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogFileResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogFileResponse", string(data)}, " ")
}
