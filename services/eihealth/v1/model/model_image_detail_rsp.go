package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 镜像详情返回体
type ImageDetailRsp struct {

	// 镜像名称
	Name *string `json:"name,omitempty"`

	// 镜像ID
	Id *string `json:"id,omitempty"`

	// 镜像类型
	Type *string `json:"type,omitempty"`

	// 镜像芯片类型
	ChipType *string `json:"chip_type,omitempty"`

	// 镜像描述
	Description *string `json:"description,omitempty"`

	// 镜像版本列表
	Tags *[]string `json:"tags,omitempty"`

	// 镜像创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 镜像更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 源资源id
	SourceResourceId *string `json:"source_resource_id,omitempty"`
}

func (o ImageDetailRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetailRsp struct{}"
	}

	return strings.Join([]string{"ImageDetailRsp", string(data)}, " ")
}
