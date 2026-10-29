package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogStreamParams 创建日志流参数。
type CreateLogStreamParams struct {

	// 需要创建的日志流名称。
	LogStreamName string `json:"log_stream_name"`

	// 日志存储时间 说明： 该参数仅对华东-上海一、华北-北京四、华南-广州用户开放。
	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`

	// 日志流名称别名
	LogStreamNameAlias *string `json:"log_stream_name_alias,omitempty"`

	// 企业项目名称 >只能由中文、英文字母、数字、下划线、中划线组成，且不能使用任何大小写形式的“default”； 描述不超过512个字符。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释：** 是否开启智能冷存储。 **约束限制：** 开启智能冷存储需要日志存储时间大于7天。开启智能冷存储后，需要设置标准存储层数据的保存时间。标准存储至少需保存7天才能转换为智能冷存储层数据。当日志保存时间大于标准存储层数据保存时间，且小于日志存储时间时，日志数据将自动转换为智能冷存储层数据存储。 **取值范围：** true/false。 **默认取值：** false
	HotColdSeparation *bool `json:"hot_cold_separation,omitempty"`

	// **参数解释：** 日志流标准存储时间（单位：天），即日志上报到LTS后日志标准存储的时间。 **约束限制：** 开启智能冷存储需要日志总存储时间大于7天。 **取值范围：** 最小值：7天 最大值：365天 **默认取值：** 不涉及。
	HotStorageDays *int32 `json:"hot_storage_days,omitempty"`
}

func (o CreateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamParams struct{}"
	}

	return strings.Join([]string{"CreateLogStreamParams", string(data)}, " ")
}
