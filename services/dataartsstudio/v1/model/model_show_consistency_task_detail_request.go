package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsistencyTaskDetailRequest Request Object
type ShowConsistencyTaskDetailRequest struct {

	// 对账作业ID
	Id string `json:"id"`

	// workspace 信息
	Workspace string `json:"workspace"`
}

func (o ShowConsistencyTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsistencyTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConsistencyTaskDetailRequest", string(data)}, " ")
}
