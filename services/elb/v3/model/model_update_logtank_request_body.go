package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLogtankRequestBody struct {
	Logtank *UpdateLogtankOption `json:"logtank"`
}

func (o UpdateLogtankRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLogtankRequestBody", string(data)}, " ")
}
