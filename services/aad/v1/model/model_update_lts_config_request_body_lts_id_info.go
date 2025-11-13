package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLtsConfigRequestBodyLtsIdInfo 日志配置信息
type UpdateLtsConfigRequestBodyLtsIdInfo struct {

	// 企业项目id
	EpsId string `json:"eps_id"`

	// region
	Region string `json:"region"`

	// 日志组id
	LtsGroupId string `json:"lts_group_id"`

	// 日志流id
	LtsAttackStreamId string `json:"lts_attack_stream_id"`
}

func (o UpdateLtsConfigRequestBodyLtsIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigRequestBodyLtsIdInfo struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigRequestBodyLtsIdInfo", string(data)}, " ")
}
