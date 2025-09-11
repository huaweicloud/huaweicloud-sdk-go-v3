package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlStatisticsBean struct {

	// 删除数量
	DeleteNum *int64 `json:"delete_num,omitempty"`

	// 插入数量
	InsertNum *int64 `json:"insert_num,omitempty"`

	// 其他数量
	OtherNum *int64 `json:"other_num,omitempty"`

	// 周期
	Period *string `json:"period,omitempty"`

	// 查询数量
	SelectNum *int64 `json:"select_num,omitempty"`

	// 更新数量
	UpdateNum *int64 `json:"update_num,omitempty"`
}

func (o SqlStatisticsBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlStatisticsBean struct{}"
	}

	return strings.Join([]string{"SqlStatisticsBean", string(data)}, " ")
}
