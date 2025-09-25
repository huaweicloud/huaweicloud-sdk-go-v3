package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddImageWhiteListsRequestBodyImageInfo struct {

	// **参数解释**： 仓库镜像id **约束限制**： 不涉及 **取值范围**： 最小值1，最大值9223372036854775807 **默认取值**： 不涉及
	Id *int64 `json:"id,omitempty"`

	// **参数解释**： 本地镜像id **约束限制**： 不涉及 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 镜像名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-256位 **默认取值**： 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 镜像类型 **约束限制**: 不涉及 **取值范围**: - local_image：本地镜像。 - private_image：私有镜像仓库。 - shared_image：共享镜像仓库。 - instance_image：企业仓库。 - harbor: harbor仓库。 - jfrog : jfrog仓库。  **默认取值**: 不涉及
	ImageType *string `json:"image_type,omitempty"`
}

func (o AddImageWhiteListsRequestBodyImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageWhiteListsRequestBodyImageInfo struct{}"
	}

	return strings.Join([]string{"AddImageWhiteListsRequestBodyImageInfo", string(data)}, " ")
}
