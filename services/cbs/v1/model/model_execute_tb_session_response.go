package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteTbSessionResponse struct {

	// 所有数据的信息。
	Questions      *[]ExecuteTbQuestion `json:"questions,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ExecuteTbSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTbSessionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTbSessionResponse", string(data)}, " ")
}
