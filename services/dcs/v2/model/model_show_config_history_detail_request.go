package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigHistoryDetailRequest Request Object
type ShowConfigHistoryDetailRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 历史实例ID
	HistoryId string `json:"history_id"`
}

func (o ShowConfigHistoryDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigHistoryDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigHistoryDetailRequest", string(data)}, " ")
}
