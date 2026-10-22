package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2TenantImageRequest Request Object
type ShowV2TenantImageRequest struct {

	// **参数解释**：镜像ID。可通过[查询自定义镜像列表]获取镜像ID(ListV2TenantImages.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`
}

func (o ShowV2TenantImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2TenantImageRequest struct{}"
	}

	return strings.Join([]string{"ShowV2TenantImageRequest", string(data)}, " ")
}
