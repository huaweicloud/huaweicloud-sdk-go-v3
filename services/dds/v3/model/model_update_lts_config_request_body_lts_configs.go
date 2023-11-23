package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLtsConfigRequestBodyLtsConfigs 单个实例需要设置的LTS策略。
type UpdateLtsConfigRequestBodyLtsConfigs struct {

	// 实例ID，可以调用“[查询实例列表和详情](x-wc://file=zh-cn_topic_0000001369935045.xml)”接口获取。如果未申请实例，可以调用“[创建实例](x-wc://file=zh-cn_topic_0000001369734929.xml)”接口创建。
	InstanceId string `json:"instance_id"`

	LogType *LtsLogType `json:"log_type"`

	// LTS日志组ID。可通过云日志服务-“查询账号下所有日志组”API接口获取。
	LtsGroupId string `json:"lts_group_id"`

	// LTS日志流ID。可通过云日志服务-“查询指定日志组下的所有日志流”API接口获取。
	LtsStreamId string `json:"lts_stream_id"`
}

func (o UpdateLtsConfigRequestBodyLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigRequestBodyLtsConfigs struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigRequestBodyLtsConfigs", string(data)}, " ")
}
