package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuntimeServiceInfo struct {

	// **参数解释** 运行时类型 **取值范围** - docker：docker运行时。 - containerd：containerd运行时。 - podman：podman运行时。 - isulad：isulad运行时。 - crio：crio运行时。 - unknown：未知运行时。
	RuntimeType *string `json:"runtime_type,omitempty"`

	// **参数解释** 运行时版本 **取值范围** 字符长度0-128
	RuntimeVersion *string `json:"runtime_version,omitempty"`

	// **参数解释** 运行时存储驱动数组 **取值范围** 数组范围0-20
	StorageDriver *[]string `json:"storage_driver,omitempty"`
}

func (o RuntimeServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeServiceInfo struct{}"
	}

	return strings.Join([]string{"RuntimeServiceInfo", string(data)}, " ")
}
