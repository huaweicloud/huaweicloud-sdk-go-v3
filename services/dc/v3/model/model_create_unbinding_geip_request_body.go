package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUnbindingGeipRequestBody struct {
	GlobalEips *[]UnbindingGeipBody `json:"global_eips,omitempty"`
}

func (o CreateUnbindingGeipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUnbindingGeipRequestBody struct{}"
	}

	return strings.Join([]string{"CreateUnbindingGeipRequestBody", string(data)}, " ")
}
