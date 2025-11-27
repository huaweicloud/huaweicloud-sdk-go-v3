package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CapacityWebListRequestProviderObj struct {

	// **参数解释：** 云服务名称。 **约束限制：** 不涉及。 **取值范围：** 字符串，长度1到64个字符。 **默认取值：** 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 资源类型名称。 **约束限制：** 不涉及。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o CapacityWebListRequestProviderObj) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityWebListRequestProviderObj struct{}"
	}

	return strings.Join([]string{"CapacityWebListRequestProviderObj", string(data)}, " ")
}
