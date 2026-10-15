package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveDataFilterRulesResponse Response Object
type SaveDataFilterRulesResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SaveDataFilterRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveDataFilterRulesResponse struct{}"
	}

	return strings.Join([]string{"SaveDataFilterRulesResponse", string(data)}, " ")
}
