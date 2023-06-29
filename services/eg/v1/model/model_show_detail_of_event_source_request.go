package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailOfEventSourceRequest Request Object
type ShowDetailOfEventSourceRequest struct {

	// 指定查询的事件源ID
	SourceId string `json:"source_id"`
}

func (o ShowDetailOfEventSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventSourceRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventSourceRequest", string(data)}, " ")
}
