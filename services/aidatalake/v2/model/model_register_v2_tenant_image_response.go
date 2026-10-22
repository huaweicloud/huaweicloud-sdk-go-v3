package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterV2TenantImageResponse Response Object
type RegisterV2TenantImageResponse struct {

	// **参数解释**：镜像ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	ImageId *string `json:"image_id,omitempty"`

	Category *V2ImageCategory `json:"category,omitempty"`

	// **参数解释**：镜像域名。 **取值范围**：支持英文字母、数字、小数点及连字符，长度0~255个字符。
	ImageRegistry *string `json:"image_registry,omitempty"`

	// **参数解释**：镜像命名空间。 **取值范围**：由英文字母、数字、中划线和下划线组成，长度0~64个字符。
	ImageNamespace *string `json:"image_namespace,omitempty"`

	// **参数解释**：镜像名。 **取值范围**：由中文字符、英文字母、数字、中划线和下划线组成，长度0~128个字符。
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**：可见性。 **取值范围**：   - PRIVATE：私有。用户自己创建的。   - PUBLIC：公共。查询所有公共的，包括其他用户创建的。
	Visibility *string `json:"visibility,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterV2TenantImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterV2TenantImageResponse struct{}"
	}

	return strings.Join([]string{"RegisterV2TenantImageResponse", string(data)}, " ")
}
