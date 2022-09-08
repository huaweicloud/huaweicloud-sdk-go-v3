package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例磁盘类型信息。
type Storage struct {

	// 磁盘类型名称，可能取值如下： - ULTRAHIGH，表示SSD。 - ESSD,表示急速云盘
	Name string `json:"name"`

	// 其中key是可用区编号，value是规格所在az的状态，包含以下状态： - normal，在售。 - unsupported，暂不支持该规格。 - sellout，售罄。
	AzStatus map[string]string `json:"az_status"`

	// 性能规格，包含以下状态： - normal：通用增强型。 - normal2：通用增强Ⅱ型。 - armFlavors：鲲鹏通用计算增强型。 - armFlavors2Shared：鲲鹏通用计算增强II型（共享型）。
	SupportComputeGroupType *[]string `json:"support_compute_group_type,omitempty"`
}

func (o Storage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Storage struct{}"
	}

	return strings.Join([]string{"Storage", string(data)}, " ")
}
