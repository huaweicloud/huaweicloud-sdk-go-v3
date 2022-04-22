package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规格信息。
type Flavor struct {

	// CPU个数。
	Vcpus string `json:"vcpus"`

	// 内存大小，单位为GB。
	Ram string `json:"ram"`

	// 资源规格编码。例如：gaussdb.opengauss.ee.dn.m6.4xlarge.8.in。
	SpecCode string `json:"spec_code"`

	// 可用az
	AvailabilityZone []string `json:"availability_zone"`

	// 其中key是可用区编号，value是规格所在az的状态，包含以下状态： - normal，在售。 - unsupported，暂不支持该规格。 - sellout，售罄。
	AzStatus map[string]string `json:"az_status"`

	// 该规格支持的数据库版本号
	Version string `json:"version"`

	// 数组库引擎名称
	Name string `json:"name"`

	// 性能规格，包含以下状态： - normal：通用增强型。 - normal2：通用增强Ⅱ型。 - armFlavors：鲲鹏通用增强型。 - dedicicatenormal ：x86独享型。 - armlocalssd：鲲鹏通用型。 - normallocalssd：x86通用型。 - general：通用型。 - dedicated：独享型，仅云盘SSD支持。 - rapid：独享型，仅极速型SSD支持。
	GroupType string `json:"group_type"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
