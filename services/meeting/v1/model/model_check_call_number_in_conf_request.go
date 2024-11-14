package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCallNumberInConfRequest Request Object
type CheckCallNumberInConfRequest struct {

	// 呼叫号码
	CallNumber string `json:"call_number"`
}

func (o CheckCallNumberInConfRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCallNumberInConfRequest struct{}"
	}

	return strings.Join([]string{"CheckCallNumberInConfRequest", string(data)}, " ")
}
