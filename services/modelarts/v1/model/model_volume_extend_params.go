package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeExtendParams **参数解释**：磁盘自定义配置。
type VolumeExtendParams struct {

	// **参数解释**：磁盘分组名称，用于各个存储空间的划分。 **取值范围**：可选项如下： - vgpaas：容器盘。 - default：普通数据盘，以默认方式挂载。 - vguser：普通数据盘，指定挂载路径，不同路径的分组名称不同，如vguser1，vguser2。 - vg-everest-localvolume-persistent：普通数据盘，作为持久存储卷 - vg-everest-localvolume-ephemeral：普通数据盘，作为临时存储卷
	VolumeGroup string `json:"volumeGroup"`
}

func (o VolumeExtendParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeExtendParams struct{}"
	}

	return strings.Join([]string{"VolumeExtendParams", string(data)}, " ")
}
