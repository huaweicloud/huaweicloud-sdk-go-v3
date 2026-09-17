package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetSqlLimitingSwitchRequest Request Object
type BatchSetSqlLimitingSwitchRequest struct {
	Body *BatchSetSqlLimitingSwitchRequestBody `json:"body,omitempty"`
}

func (o BatchSetSqlLimitingSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSqlLimitingSwitchRequest struct{}"
	}

	return strings.Join([]string{"BatchSetSqlLimitingSwitchRequest", string(data)}, " ")
}
