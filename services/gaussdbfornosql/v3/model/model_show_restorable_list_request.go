package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRestorableListRequest struct {

	// 备份文件ID
	BackupId string `json:"backup_id"`

	// 索引位置偏移量。取值大于或等于0。不传该参数时，查询偏移量默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。取值范围：1~100。不传该参数时，默认查询前100条信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowRestorableListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestorableListRequest struct{}"
	}

	return strings.Join([]string{"ShowRestorableListRequest", string(data)}, " ")
}
