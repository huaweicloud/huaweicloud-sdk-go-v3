package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcBaseInfo struct {
	// 云服务器ID

	EcsId *string `json:"ecs_id,omitempty"`
	// 云服务器名称

	EcsName *int32 `json:"ecs_name,omitempty"`
	// 是否使用级联方式  暂不支持

	CascadeFlag *bool `json:"cascade_flag,omitempty"`
}

func (o VpcBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcBaseInfo struct{}"
	}

	return strings.Join([]string{"VpcBaseInfo", string(data)}, " ")
}
