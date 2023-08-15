package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkitemStatus struct {

	// 工作项状态变更记录的id,每次变更产生一条新的记录id
	Id *string `json:"id,omitempty"`

	Status *WorkitemStatusStatus `json:"status,omitempty"`
}

func (o WorkitemStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemStatus struct{}"
	}

	return strings.Join([]string{"WorkitemStatus", string(data)}, " ")
}
