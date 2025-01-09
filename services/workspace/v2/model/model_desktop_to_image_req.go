package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopToImageReq 桌面转镜像请求。
type DesktopToImageReq struct {

	// 镜像名称。
	ImageName string `json:"image_name"`

	// 镜像描述信息。
	ImageDescription *string `json:"image_description,omitempty"`

	// 用于制作镜像的云桌面的InstanceID。
	DesktopId string `json:"desktop_id"`

	// 是否执行系统封装步骤。
	ExecuteSysprep *bool `json:"execute_sysprep,omitempty"`

	// 镜像标签列表。
	ImageTags *[]TagKeyValue `json:"image_tags,omitempty"`

	// 表示当前镜像所属的企业项目。取值为0或无该值，表示属于default企业项目。取值为UUID，表示属于该UUID对应的企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 表示镜像支持的最大内存，单位为MB。
	MaxRam *string `json:"max_ram,omitempty"`

	// 表示镜像支持的最小内存，单位为MB，默认为0，表示不受限制。
	MinRam *string `json:"min_ram,omitempty"`
}

func (o DesktopToImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopToImageReq struct{}"
	}

	return strings.Join([]string{"DesktopToImageReq", string(data)}, " ")
}
