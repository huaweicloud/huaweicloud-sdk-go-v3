package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleRdsRequest Request Object
type ShowAvalibleRdsRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	// 目标实例 ID。
	TargetInstanceId string `json:"target_instance_id"`

	// 源DN实例 ID。
	SourceDnInstanceId string `json:"source_dn_instance_id"`

	// 恢复时间。
	RestoreTime string `json:"restore_time"`
}

func (o ShowAvalibleRdsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleRdsRequest struct{}"
	}

	return strings.Join([]string{"ShowAvalibleRdsRequest", string(data)}, " ")
}
