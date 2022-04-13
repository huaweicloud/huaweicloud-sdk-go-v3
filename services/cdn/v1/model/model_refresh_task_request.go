package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RefreshTaskRequest struct {
	RefreshTask *RefreshTaskRequestBody `json:"refresh_task"`
}

func (o RefreshTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshTaskRequest struct{}"
	}

	return strings.Join([]string{"RefreshTaskRequest", string(data)}, " ")
}
