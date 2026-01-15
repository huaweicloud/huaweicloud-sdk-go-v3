package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInstancesReq 分配用户请求。
type AttachInstancesReq struct {

	// 桌面信息列表。
	Desktops *[]AttachInstancesDesktopInfo `json:"desktops,omitempty"`

	// 镜像类型，涉及变更镜像时需传（可选）。
	ImageType *string `json:"image_type,omitempty"`

	// 模板ID，涉及变更镜像时需传（可选）。
	ImageId *string `json:"image_id,omitempty"`

	// 策略id，用于指定生成桌面名称策略，如果指定了桌面名称则优先使用指定的桌面名称。
	DesktopNamePolicyId *string `json:"desktop_name_policy_id,omitempty"`

	EncryptType *EncryptType `json:"encrypt_type,omitempty"`

	// 密钥ID，encrypt_type为ENCRYPTED时必传。
	KmsId *string `json:"kms_id,omitempty"`
}

func (o AttachInstancesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInstancesReq struct{}"
	}

	return strings.Join([]string{"AttachInstancesReq", string(data)}, " ")
}
