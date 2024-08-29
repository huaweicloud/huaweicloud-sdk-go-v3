package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDigitalHumanBusinessCardResponse Response Object
type ListDigitalHumanBusinessCardResponse struct {

	// 数字人名片制作任务总数。
	Count *int32 `json:"count,omitempty"`

	// 数字人名片制作任务列表。
	Jobs *[]DigitalHumanBusinessCardJobInfo `json:"jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDigitalHumanBusinessCardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDigitalHumanBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"ListDigitalHumanBusinessCardResponse", string(data)}, " ")
}
