package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteLtsConfigRequestBodyLtsConfigs struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	LogType *LtsLogType `json:"log_type,omitempty"`
}

func (o DeleteLtsConfigRequestBodyLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLtsConfigRequestBodyLtsConfigs struct{}"
	}

	return strings.Join([]string{"DeleteLtsConfigRequestBodyLtsConfigs", string(data)}, " ")
}
