package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorParserResponse Response Object
type DeleteCollectorParserResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteCollectorParserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorParserResponse struct{}"
	}

	return strings.Join([]string{"DeleteCollectorParserResponse", string(data)}, " ")
}
