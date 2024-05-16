package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapFlavorInfoFlavors struct {

	// 规格类型，取值为arm、x86和generalX86。  arm：独享型arm规格。  x86：独享型x86规格。  generalX86：通用型X86规格。
	Type *string `json:"type,omitempty"`

	// CPU大小。例如：1表示1U。
	Vcpus *string `json:"vcpus,omitempty"`

	// 内存大小，单位为GB。
	Ram *string `json:"ram,omitempty"`

	// 规格ID，该字段唯一。
	Id *string `json:"id,omitempty"`

	// 资源规格编码，.同创建指定的flavor_ref。例如：gaussdb.sr-be.xlarge.x86.4。  “gaussdb.sr”代表产品。  “xlarge” 代表计算规格为4U。  “x86” 代表CPU架构为x86。  “4” 表示vCPU和内存为1:4。
	SpecCode *string `json:"spec_code,omitempty"`

	// 数据库版本号。
	VersionName *string `json:"version_name,omitempty"`

	// 实例类型。目前仅支持Cluster、Single。
	InstanceMode *string `json:"instance_mode,omitempty"`

	// 规格所在AZ的状态，包含以下状态：  normal：在售。  unsupported：暂不支持该规格。  sellout：售罄。
	AzStatus map[string]string `json:"az_status,omitempty"`
}

func (o HtapFlavorInfoFlavors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapFlavorInfoFlavors struct{}"
	}

	return strings.Join([]string{"HtapFlavorInfoFlavors", string(data)}, " ")
}
