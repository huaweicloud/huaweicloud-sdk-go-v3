package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AvailableFlavorInfoResult 实例可变更规格信息。
type AvailableFlavorInfoResult struct {

	// 资源规格编码。
	SpecCode *string `json:"spec_code,omitempty"`

	// CPU核数。
	Vcpus *string `json:"vcpus,omitempty"`

	// 内存大小，单位：GB。
	Ram *string `json:"ram,omitempty"`

	// 其中key是可用区编号，value是规格所在az的状态。
	AzStatus map[string]string `json:"az_status,omitempty"`
}

func (o AvailableFlavorInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableFlavorInfoResult struct{}"
	}

	return strings.Join([]string{"AvailableFlavorInfoResult", string(data)}, " ")
}
