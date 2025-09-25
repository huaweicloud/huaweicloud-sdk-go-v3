package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageBasicImageRequest Request Object
type ListImageBasicImageRequest struct {

	// Region ID
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 镜像id
	ImageId string `json:"image_id"`

	// 镜像类型，包含如下:   - private_image : 私有镜像仓库   - shared_image : 共享镜像仓库   - instance_image : 企业镜像   - cicd : cicd镜像   - harbor ：Harbor仓库镜像
	ImageType string `json:"image_type"`

	// 组织名称
	Namespace *string `json:"namespace,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 镜像版本名称
	ImageVersion *string `json:"image_version,omitempty"`

	// os名称
	OsName *string `json:"os_name,omitempty"`
}

func (o ListImageBasicImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageBasicImageRequest struct{}"
	}

	return strings.Join([]string{"ListImageBasicImageRequest", string(data)}, " ")
}
