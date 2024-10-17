package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteInstanceDemandRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	Id string `json:"id"`

	// 是否删除弹性IP
	DeletePublicip *bool `json:"delete_publicip,omitempty"`

	// 是否删除磁盘
	DeleteVolume *bool `json:"delete_volume,omitempty"`
}

func (o DeleteInstanceDemandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceDemandRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceDemandRequest", string(data)}, " ")
}
