package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchoverTestResponse Response Object
type SwitchoverTestResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	SwitchoverTestRecord *SwitchoverTestRecord `json:"switchover_test_record,omitempty"`
	HttpStatusCode       int                   `json:"-"`
}

func (o SwitchoverTestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchoverTestResponse struct{}"
	}

	return strings.Join([]string{"SwitchoverTestResponse", string(data)}, " ")
}
