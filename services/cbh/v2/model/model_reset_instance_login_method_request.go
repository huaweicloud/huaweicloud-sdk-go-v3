package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetInstanceLoginMethodRequest Request Object
type ResetInstanceLoginMethodRequest struct {
	Body *CommonCbhRequestBody `json:"body,omitempty"`
}

func (o ResetInstanceLoginMethodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetInstanceLoginMethodRequest struct{}"
	}

	return strings.Join([]string{"ResetInstanceLoginMethodRequest", string(data)}, " ")
}
