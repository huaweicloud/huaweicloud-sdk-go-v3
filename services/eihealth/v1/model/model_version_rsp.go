package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionRsp struct {

	// 资产版本号
	Version *string `json:"version,omitempty"`

	// 发布者
	Publisher *string `json:"publisher,omitempty"`

	// 资产长描述
	Descritpion *string `json:"descritpion,omitempty"`

	// 资产短描述
	Summary *string `json:"summary,omitempty"`

	// 许可证
	License *string `json:"license,omitempty"`

	// 资产状态
	Status *string `json:"status,omitempty"`

	// 资产发布失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 资产标签列表
	Labels *[]string `json:"labels,omitempty"`

	// 资产封面图访问链接
	Picture *string `json:"picture,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o VersionRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionRsp struct{}"
	}

	return strings.Join([]string{"VersionRsp", string(data)}, " ")
}
