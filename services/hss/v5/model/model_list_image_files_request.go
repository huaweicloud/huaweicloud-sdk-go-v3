package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageFilesRequest Request Object
type ListImageFilesRequest struct {

	// Region ID
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像类型，包含如下:   - private_image : 私有镜像仓库   - shared_image : 共享镜像仓库   - instance_image : 企业镜像   - cicd : cicd镜像   - harbor ：Harbor仓库镜像
	ImageType string `json:"image_type"`

	// 镜像id
	ImageId string `json:"image_id"`

	// 组织名称
	Namespace *string `json:"namespace,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本名称
	TagName *string `json:"tag_name,omitempty"`

	// 文件名称过滤查询(支持模糊匹配)
	FileName *string `json:"file_name,omitempty"`

	// 返回的条数，默认10
	Limit *int32 `json:"limit,omitempty"`

	// 起始索引，默认是0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListImageFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageFilesRequest struct{}"
	}

	return strings.Join([]string{"ListImageFilesRequest", string(data)}, " ")
}
