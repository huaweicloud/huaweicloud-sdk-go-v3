package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageStatisticDo struct {

	// 参数解释: 文件总数。 取值范围: 不涉及。
	FilesCount *int64 `json:"filesCount,omitempty"`

	// 参数解释: 占用空间(字节)。 取值范围: 不涉及。
	UsedSpaceLength *int64 `json:"usedSpaceLength,omitempty"`

	// 参数解释: 占用空间。 取值范围: 不涉及。
	UsedSpace *string `json:"usedSpace,omitempty"`

	// 参数解释: 统计日期。 取值范围: 不涉及。
	SummaryDate *string `json:"summaryDate,omitempty"`

	// 参数解释: 文件夹总数。 取值范围: 不涉及。
	FoldersCount *int64 `json:"foldersCount,omitempty"`

	// 参数解释: 文件及文件夹总数。 取值范围: 不涉及。
	ItemsCount *int64 `json:"itemsCount,omitempty"`
}

func (o StorageStatisticDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageStatisticDo struct{}"
	}

	return strings.Join([]string{"StorageStatisticDo", string(data)}, " ")
}
