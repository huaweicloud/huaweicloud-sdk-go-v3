package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像版本详情
type GetTagDetailRsp struct {

	// 镜像版本名称
	Tag *string `json:"tag,omitempty"`

	// 镜像版本大小
	Size *int64 `json:"size,omitempty"`

	// 镜像版本创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 镜像版本更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o GetTagDetailRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTagDetailRsp struct{}"
	}

	return strings.Join([]string{"GetTagDetailRsp", string(data)}, " ")
}
