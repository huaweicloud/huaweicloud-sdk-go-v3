package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// os-extended-volumes:volumes_attached字段数据结构说明
type OsExtendedVolumes struct {
	// 云硬盘ID

	Id *string `json:"id,omitempty"`
	// 删裸金属服务器时是否一并删除该卷。true：是false：否

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o OsExtendedVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsExtendedVolumes struct{}"
	}

	return strings.Join([]string{"OsExtendedVolumes", string(data)}, " ")
}
