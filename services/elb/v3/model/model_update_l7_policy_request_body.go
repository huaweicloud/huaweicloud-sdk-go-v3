package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateL7PolicyRequestBody struct {
	L7policy *UpdateL7PolicyOption `json:"l7policy"`
}

func (o UpdateL7PolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyRequestBody", string(data)}, " ")
}
