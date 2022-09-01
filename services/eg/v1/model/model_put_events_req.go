package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutEventsReq struct {
	Events *[]CloudEvents `json:"events,omitempty" xml:"events"`
}

func (o PutEventsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsReq struct{}"
	}

	return strings.Join([]string{"PutEventsReq", string(data)}, " ")
}
