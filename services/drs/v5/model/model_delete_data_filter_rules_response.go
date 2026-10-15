package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataFilterRulesResponse Response Object
type DeleteDataFilterRulesResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteDataFilterRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataFilterRulesResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataFilterRulesResponse", string(data)}, " ")
}
