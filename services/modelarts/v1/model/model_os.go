package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Os **参数解释**：操作系统镜像信息。
type Os struct {

	// **参数解释**：操作系统名称和版本，如EulorOS 2.5。指定私有镜像时可不指定。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：操作系统镜像id。 **取值范围**：不涉及。
	ImageId *string `json:"imageId,omitempty"`

	// **参数解释**：操作系统镜像类型。设置私有镜像时必须指定。默认为预置镜像，无需指定该字段。 **取值范围**：可选值如下： - private：私有镜像 - \"\"：不指定类型即预置镜像。
	ImageType *string `json:"imageType,omitempty"`

	// **参数解释**：操作系统镜像自动匹配配置。当配置该参数时将会自动选择最优镜像，同时该参数会自动清空。 **取值范围**：操作系统名称和版本，如EulorOS 2.5。
	AutoMatch *string `json:"autoMatch,omitempty"`
}

func (o Os) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Os struct{}"
	}

	return strings.Join([]string{"Os", string(data)}, " ")
}
