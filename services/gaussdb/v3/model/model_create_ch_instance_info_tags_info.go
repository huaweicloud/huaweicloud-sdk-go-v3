package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceInfoTagsInfo 实例标签。
type CreateChInstanceInfoTagsInfo struct {

	// 用户标签。
	Tags *[]CreateChInstanceInfoTagsInfoTags `json:"tags,omitempty"`

	// 系统标签。
	SysTags *[]CreateChInstanceInfoTagsInfoTags `json:"sys_tags,omitempty"`
}

func (o CreateChInstanceInfoTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceInfoTagsInfo struct{}"
	}

	return strings.Join([]string{"CreateChInstanceInfoTagsInfo", string(data)}, " ")
}
