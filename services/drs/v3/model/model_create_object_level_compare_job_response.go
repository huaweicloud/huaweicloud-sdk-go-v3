package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectLevelCompareJobResponse Response Object
type CreateObjectLevelCompareJobResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateObjectLevelCompareJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectLevelCompareJobResponse struct{}"
	}

	return strings.Join([]string{"CreateObjectLevelCompareJobResponse", string(data)}, " ")
}
