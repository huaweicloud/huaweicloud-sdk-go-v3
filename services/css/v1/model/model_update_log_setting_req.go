package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLogSettingReq struct {

	// 委托名称，委托给CSS，允许CSS调用您的其他云服务。
	Agency string `json:"agency"`

	// 日志在OBS桶中的备份路径。
	LogBasePath string `json:"log_base_path"`

	// 用于存储日志的OBS桶的桶名。
	LogBucket string `json:"log_bucket"`

	// 设置保存日志的索引前缀。
	IndexPrefix *string `json:"index_prefix,omitempty"`

	// 修改日志保存的天数。
	KeepDays *int32 `json:"keep_days,omitempty"`

	// 设置保存日志的目标集群ID。
	TargetClusterId *string `json:"target_cluster_id,omitempty"`

	// 集群日志是否开启自动备份。
	AutoEnable *bool `json:"auto_enable,omitempty"`

	// 集群日志备份开始时间。
	Period *string `json:"period,omitempty"`
}

func (o UpdateLogSettingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogSettingReq struct{}"
	}

	return strings.Join([]string{"UpdateLogSettingReq", string(data)}, " ")
}
