package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeSqlSwitchBody struct {

	// 开关类型。取值DAS SQL Explorer和DAS Slow Query Log，分别表示DAS收集全量SQL开关和DAS收集慢SQL开关。
	Type string `json:"type" xml:"type"`

	// 开关状态，取值0和1，分别代表要关闭和开启。
	Status int32 `json:"status" xml:"status"`

	// 数据库类型。当前全量SQL支持的数据库类型包括MySQL和GaussDB(for MySQL)，慢SQL支持的类型：MySQL、GaussDB(for MySQL)、PostgreSQL。
	DatastoreType string `json:"datastore_type" xml:"datastore_type"`

	// SQL数据保存时长（天）。默认为7天，最长可保留30天，到期后数据自动删除。如果要保留30天以上，请到DAS页面进行操作。
	RetentionDays *int64 `json:"retention_days,omitempty" xml:"retention_days"`
}

func (o ChangeSqlSwitchBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSqlSwitchBody struct{}"
	}

	return strings.Join([]string{"ChangeSqlSwitchBody", string(data)}, " ")
}
