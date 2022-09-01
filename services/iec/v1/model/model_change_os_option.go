package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 切换操作系统的参数
type ChangeOsOption struct {

	// 切换系统所使用的新镜像的ID。
	ImageId string `json:"image_id" xml:"image_id"`

	Metadata *ChangeOsMetadata `json:"metadata,omitempty" xml:"metadata"`

	// 密钥对名称。 如果需要使用SSH密钥方式登录边缘实例，请指定已创建密钥的名称。
	KeyName *string `json:"key_name,omitempty" xml:"key_name"`
}

func (o ChangeOsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOsOption struct{}"
	}

	return strings.Join([]string{"ChangeOsOption", string(data)}, " ")
}
