package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrsJobNameResponse Response Object
type ShowDrsJobNameResponse struct {

	// 参数解释： 任务名字。
	DrsName        *string `json:"drs_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDrsJobNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrsJobNameResponse struct{}"
	}

	return strings.Join([]string{"ShowDrsJobNameResponse", string(data)}, " ")
}
