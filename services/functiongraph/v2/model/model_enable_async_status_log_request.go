package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAsyncStatusLogRequest Request Object
type EnableAsyncStatusLogRequest struct {
}

func (o EnableAsyncStatusLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAsyncStatusLogRequest struct{}"
	}

	return strings.Join([]string{"EnableAsyncStatusLogRequest", string(data)}, " ")
}
