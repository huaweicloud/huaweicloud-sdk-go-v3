package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportImageQuickResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportImageQuickResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickResponse struct{}"
	}

	return strings.Join([]string{"ImportImageQuickResponse", string(data)}, " ")
}
