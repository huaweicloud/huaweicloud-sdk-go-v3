package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventSourceRequest Request Object
type DeleteEventSourceRequest struct {

	// 指定查询的事件源ID
	SourceId string `json:"source_id"`
}

func (o DeleteEventSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventSourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventSourceRequest", string(data)}, " ")
}
