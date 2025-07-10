package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlViolationsResponse Response Object
type ListControlViolationsResponse struct {

	// 控制策略合规性。
	ControlViolations *[]ControlViolation `json:"control_violations,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListControlViolationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlViolationsResponse struct{}"
	}

	return strings.Join([]string{"ListControlViolationsResponse", string(data)}, " ")
}
