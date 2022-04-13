package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogConfiguration struct {
	// ID。

	Id *string `json:"id,omitempty"`
	// 集群ID。

	ClusterId *string `json:"clusterId,omitempty"`
	// OBS桶。

	ObsBucket *string `json:"obsBucket,omitempty"`
	// IAM 委托。

	Agency *string `json:"agency,omitempty"`
	// 更新时间。

	UpdateAt *int64 `json:"updateAt,omitempty"`
	// 备份路径。

	BasePath *string `json:"basePath,omitempty"`
	// 自动备份开关。

	AutoEnable *bool `json:"autoEnable,omitempty"`
	// 备份开始时间。

	Period *string `json:"period,omitempty"`
	// 日志开关。

	LogSwitch *bool `json:"logSwitch,omitempty"`
}

func (o LogConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfiguration struct{}"
	}

	return strings.Join([]string{"LogConfiguration", string(data)}, " ")
}
