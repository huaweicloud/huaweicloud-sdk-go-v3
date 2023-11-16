package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAsyncStatusLogResponse Response Object
type EnableAsyncStatusLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableAsyncStatusLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAsyncStatusLogResponse struct{}"
	}

	return strings.Join([]string{"EnableAsyncStatusLogResponse", string(data)}, " ")
}
