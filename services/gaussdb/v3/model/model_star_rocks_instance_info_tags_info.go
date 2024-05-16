package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksInstanceInfoTagsInfo 实例标签。
type StarRocksInstanceInfoTagsInfo struct {

	// 用户标签。
	Tags *[]StarRocksInstanceInfoTagsInfoTags `json:"tags,omitempty"`

	// 系统标签。
	SysTags *[]StarRocksInstanceInfoTagsInfoSysTags `json:"sys_tags,omitempty"`
}

func (o StarRocksInstanceInfoTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoTagsInfo struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoTagsInfo", string(data)}, " ")
}
