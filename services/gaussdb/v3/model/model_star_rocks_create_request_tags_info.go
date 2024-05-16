package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksCreateRequestTagsInfo 标签信息。
type StarRocksCreateRequestTagsInfo struct {

	// 系统标签。
	SysTags []StarRocksCreateRequestTagsInfoSysTags `json:"sys_tags"`
}

func (o StarRocksCreateRequestTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestTagsInfo struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestTagsInfo", string(data)}, " ")
}
