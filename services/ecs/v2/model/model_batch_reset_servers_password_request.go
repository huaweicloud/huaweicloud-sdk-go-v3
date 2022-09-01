package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchResetServersPasswordRequest struct {
	Body *BatchResetServersPasswordRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchResetServersPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResetServersPasswordRequest struct{}"
	}

	return strings.Join([]string{"BatchResetServersPasswordRequest", string(data)}, " ")
}
