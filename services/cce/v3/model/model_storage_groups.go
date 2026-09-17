package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageGroups struct {

	// storageGroups的名字，作为虚拟存储组的名字，因此各个group名字不能重复。 > - 当cceManaged=ture时，name必须为：vgpass。 > - 当数据盘作为临时存储卷时：name必须为：vg-everest-localvolume-ephemeral。 > - 当数据盘作为持久存储卷时：name必须为：vg-everest-localvolume-persistent。
	Name string `json:"name"`

	// k8s及runtime所属存储空间。有且仅有一个group被设置为true，不填默认false。
	CceManaged *bool `json:"cceManaged,omitempty"`

	// **参数解释**： 对应storageSelectors中的name，一个group可选择多个selector；但一个selector只能被一个group选择。 **约束限制**： 系统组件无法分别存储于系统盘与数据盘中，因此选择selector的type为system时，group只能选择一个selector。 **取值范围**： 不涉及 **默认取值**： 不涉及
	SelectorNames []string `json:"selectorNames"`

	// **参数解释**： group中空间配置的详细管理。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	VirtualSpaces []VirtualSpace `json:"virtualSpaces"`
}

func (o StorageGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageGroups struct{}"
	}

	return strings.Join([]string{"StorageGroups", string(data)}, " ")
}
