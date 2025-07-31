package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportImageQuickInSafeModeResponse Response Object
type ImportImageQuickInSafeModeResponse struct {

	// 异步任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportImageQuickInSafeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickInSafeModeResponse struct{}"
	}

	return strings.Join([]string{"ImportImageQuickInSafeModeResponse", string(data)}, " ")
}
