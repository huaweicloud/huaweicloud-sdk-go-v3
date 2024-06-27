package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceRequestBodyTagsInfo 标签值。
type CreateChInstanceRequestBodyTagsInfo struct {

	// 系统标签。
	SysTags []CreateChInstanceRequestBodyTagsInfoSysTags `json:"sys_tags"`
}

func (o CreateChInstanceRequestBodyTagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceRequestBodyTagsInfo struct{}"
	}

	return strings.Join([]string{"CreateChInstanceRequestBodyTagsInfo", string(data)}, " ")
}
