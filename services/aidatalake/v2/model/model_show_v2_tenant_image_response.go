package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2TenantImageResponse Response Object
type ShowV2TenantImageResponse struct {

	// **参数解释**：镜像ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像名。 **取值范围**：由中文字符、英文字母、数字、中划线和下划线组成，长度0~128个字符。
	ImageName *string `json:"image_name,omitempty"`

	Category *V2ImageCategory `json:"category,omitempty"`

	// **参数解释**：镜像域名。 **取值范围**：支持英文字母、数字、小数点及连字符，长度0~255个字符。
	ImageRegistry *string `json:"image_registry,omitempty"`

	// **参数解释**：镜像命名空间。 **取值范围**：由英文字母、数字、中划线和下划线组成，长度0~64个字符。
	ImageNamespace *string `json:"image_namespace,omitempty"`

	// **参数解释**：镜像版本数量。 **取值范围**：大于等于0。
	VersionCount *int32 `json:"version_count,omitempty"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释**：资源更新时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// **参数解释**：可见性。 **取值范围**：   - PRIVATE：私有，用户自己创建的。   - PUBLIC：公共，查询所有公共的，包括其他用户创建的。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释**：镜像的描述。 **取值范围**：支持中文字符、英文字母、数字及常用符号（不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号），长度0~2000个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：公共镜像使用说明。 **取值范围**：支持中文字符、英文字母、数字及常用符号，长度0~20000个字符。
	Instruction    *string `json:"instruction,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowV2TenantImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2TenantImageResponse struct{}"
	}

	return strings.Join([]string{"ShowV2TenantImageResponse", string(data)}, " ")
}
