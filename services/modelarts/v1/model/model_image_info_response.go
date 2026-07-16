package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageInfoResponse **参数解释：** 镜像配置。
type ImageInfoResponse struct {

	// **参数解释：** 镜像id。 **取值范围：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 镜像类别，标识镜像来源。 **取值范围：** - SWR：软件仓库服务。
	Source string `json:"source"`

	// **参数解释：** 镜像地址，source不同取值时，地址为不同值。 **取值范围：** 不涉及。
	SwrPath string `json:"swr_path"`

	// **参数解释：** 镜像支持的规格。 **取值范围：** - GPU：图形处理器。 - CPU：中央处理器。 - ASCEND：昇腾芯片。
	Category *string `json:"category,omitempty"`
}

func (o ImageInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageInfoResponse struct{}"
	}

	return strings.Join([]string{"ImageInfoResponse", string(data)}, " ")
}
