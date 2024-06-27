package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SrCreateInstanceRspInstanceTagsInfo 实例标签。
type SrCreateInstanceRspInstanceTagsInfo struct {

	// 用户标签。默认为空。
	Tags *[]CreateChInstanceInfoTagsInfoTags `json:"tags,omitempty"`

	// 系统标签。
	SysTags *[]SrCreateInstanceRspInstanceTagsInfoSysTags `json:"sys_tags,omitempty"`
}

func (o SrCreateInstanceRspInstanceTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrCreateInstanceRspInstanceTagsInfo struct{}"
	}

	return strings.Join([]string{"SrCreateInstanceRspInstanceTagsInfo", string(data)}, " ")
}
