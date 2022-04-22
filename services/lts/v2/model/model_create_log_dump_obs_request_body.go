package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 此参数在请求实体中，采用json字符串格式
type CreateLogDumpObsRequestBody struct {

	// 日志组id。
	LogGroupId string `json:"log_group_id"`

	// 日志流id列表, 可以指定一个或多个日志流进行obs周期性转储
	LogStreamIds []string `json:"log_stream_ids"`

	// obs 桶名称。
	ObsBucketName string `json:"obs_bucket_name"`

	// 周期性转储, 必须填 cycle。
	Type string `json:"type"`

	// 转储格式 RAW/JSON， 默认为 RAW。
	StorageFormat string `json:"storage_format"`

	// 是否开启转储 true/false, 默认为 true
	SwitchOn *bool `json:"switch_on,omitempty"`

	// 转储至OBS桶中的日志文件前缀。
	PrefixName *string `json:"prefix_name,omitempty"`

	// 自定义文件夹路径。
	DirPrefixName *string `json:"dir_prefix_name,omitempty"`

	// 转储周期的长度， 与 period_unit 拼接后必须在该列表中 [\"2min\",\"5min\",\"30min\",\"1hour\",\"3hour\",\"6hour\",\"12hour\"]。
	Period int32 `json:"period"`

	// 转储周期的单位， 与 period 拼接后必须在该列表中 [\"2min\",\"5min\",\"30min\",\"1hour\",\"3hour\",\"6hour\",\"12hour\"]。
	PeriodUnit string `json:"period_unit"`
}

func (o CreateLogDumpObsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsRequestBody", string(data)}, " ")
}
