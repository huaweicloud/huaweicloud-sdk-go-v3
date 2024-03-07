package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInstancesReq 分配用户请求。
type AttachInstancesReq struct {

	// 桌面信息列表。
	Desktops []AttachInstancesDesktopInfo `json:"desktops"`

	// 镜像类型，涉及变更镜像时需传（可选）
	ImageType *string `json:"image_type,omitempty"`

	// 模板ID，涉及变更镜像时需传（可选）
	ImageId *string `json:"image_id,omitempty"`

	// os类型，涉及变更镜像时需传（可选）
	OsType *string `json:"os_type,omitempty"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id，用于指定生成桌面名称策略，如果指定了桌面名称则优先使用指定的桌面名称。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`
}

func (o AttachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInstancesReq struct{}"
	}

	return strings.Join([]string{"AttachInstancesReq", string(data)}, " ")
}
