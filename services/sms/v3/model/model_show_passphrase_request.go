package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPassphraseRequest Request Object
type ShowPassphraseRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowPassphraseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPassphraseRequest struct{}"
	}

	return strings.Join([]string{"ShowPassphraseRequest", string(data)}, " ")
}
