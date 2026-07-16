package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageInfo **参数解释：** 镜像配置。 **约束限制：** 不涉及。
type ImageInfo struct {

	// **参数解释：** 镜像ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 镜像类别，标识镜像来源。 **约束限制：** 不涉及。 **取值范围：** - SWR：软件仓库服务。 - [IMAGE：[通用镜像]。](tag:hws,hws_hk) **默认取值：** 不涉及。
	Source string `json:"source"`

	// **参数解释：** 镜像地址，source不同取值时，地址为不同值。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SwrPath string `json:"swr_path"`

	// **参数解释：** 镜像支持的规格。 **约束限制：** 不涉及。 **取值范围：** - GPU：图形处理器。 - CPU：中央处理器。 - ASCEND：昇腾芯片。 **默认取值：** CPU。
	Category *string `json:"category,omitempty"`
}

func (o ImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfo struct{}"
	}

	return strings.Join([]string{"ImageInfo", string(data)}, " ")
}
