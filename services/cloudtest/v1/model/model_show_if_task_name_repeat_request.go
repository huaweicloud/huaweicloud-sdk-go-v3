package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIfTaskNameRepeatRequest Request Object
type ShowIfTaskNameRepeatRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// UUID
	Id *string `json:"id,omitempty"`

	// 查询的模板名称
	Name string `json:"name"`
}

func (o ShowIfTaskNameRepeatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIfTaskNameRepeatRequest struct{}"
	}

	return strings.Join([]string{"ShowIfTaskNameRepeatRequest", string(data)}, " ")
}
