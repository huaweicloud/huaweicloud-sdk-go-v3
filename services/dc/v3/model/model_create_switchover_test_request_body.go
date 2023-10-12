package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSwitchoverTestRequestBody struct {
	SwitchoverTestRecord *CreateSwitchoverTest `json:"switchover_test_record,omitempty"`
}

func (o CreateSwitchoverTestRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSwitchoverTestRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSwitchoverTestRequestBody", string(data)}, " ")
}
