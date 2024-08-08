package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLatestAttachedServerAppRequest Request Object
type ListLatestAttachedServerAppRequest struct {

	// 镜像实例唯一标识。
	ServerId string `json:"server_id"`
}

func (o ListLatestAttachedServerAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLatestAttachedServerAppRequest struct{}"
	}

	return strings.Join([]string{"ListLatestAttachedServerAppRequest", string(data)}, " ")
}
