package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DateTitleDto struct {

	// 时间（天）
	Day *string `json:"day,omitempty"`

	// 提交数量
	CommitsCount *int32 `json:"commits_count,omitempty"`
}

func (o DateTitleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DateTitleDto struct{}"
	}

	return strings.Join([]string{"DateTitleDto", string(data)}, " ")
}
