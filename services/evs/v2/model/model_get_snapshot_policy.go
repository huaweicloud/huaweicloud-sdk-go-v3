package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetSnapshotPolicy 查询自动快照策略模型。
type GetSnapshotPolicy struct {

	// 自动快照策略ID。
	Id string `json:"id"`

	// 自动快照策略名称。最大支持255个字节。  最小长度：0  最大长度：255
	Name string `json:"name"`

	// 自动快照的重复日期，单位为天，周期为星期。取值范围为 1~7，如 1 表示周一。  当一星期内需要创建多次自动快照时，可以传入多个时间点：  最多传入 7 个时间点。 多个时间点用一个格式类似\"1\", \"2\", … \"7\"的 JSON 数组表示，时间点之间用半角逗号（,）隔开。
	RepeatWeekdays []string `json:"repeat_weekdays"`

	// 自动快照的间隔日期，单位为天，取值范围为 1~30,如1表示每天。
	RepeatIntervalDays int32 `json:"repeat_interval_days"`

	// 自动快照的创建时间点。使用 UTC +8 时间，单位为小时。取值范围为 0~23，代表 00:00 至 23:00 共 24 个时间点，如 1 表示 01:00。  当一天内需要创建多次自动快照时，可以传入多个时间点：  最多传入 24 个时间点。 多个时间点用一个格式类似\"0\", \"1\", … \"23\"的 JSON 数组表示，时间点之间用半角逗号（,）隔开。
	TimePoints []string `json:"time_points"`

	// 设置快照的保留时间，单位为天。保留时间到期后快照会被自动释放，取值范围：1~65536。 缺省值：-1，表示快照持续保留。
	RetentionDays int32 `json:"retention_days"`

	// 自动快照策略的创建时间。
	CreatedAt string `json:"created_at"`

	// 自动快照策略的更新时间。
	UpdatedAt string `json:"updated_at"`

	// 设置自动快照策略的标签。
	Tags map[string]string `json:"tags"`

	// 企业项目id。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o GetSnapshotPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetSnapshotPolicy struct{}"
	}

	return strings.Join([]string{"GetSnapshotPolicy", string(data)}, " ")
}
