package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flavor **参数解释**： 规格信息返回值。 **取值范围**： 不涉及
type Flavor struct {

	// **参数解释**： 实例的CPU核数。 **取值范围**： 不涉及
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**： 实例的内存大小。单位GB。 **取值范围**： 不涉及
	Ram *int32 `json:"ram,omitempty"`

	// **参数解释**： 规格名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 节点规格支持的Region。 **取值范围**： 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**： 节点类型名称。 **取值范围**： 不涉及
	Typename *string `json:"typename,omitempty"`

	// **参数解释**： 实例的硬盘容量范围，单位GB。 **取值范围**： 不涉及
	Diskrange *string `json:"diskrange,omitempty"`

	// **参数解释**： 此参数是Region级配置，某个AZ没有在condOperationAz参数中配置时默认使用此参数的取值。 **取值范围**：   - normal：正常商用。   - sellout：售罄。
	CondOperationStatus *string `json:"condOperationStatus,omitempty"`

	// **参数解释**： 此参数是AZ级配置，某个AZ没有在此参数中配置时默认使用condOperationAz参数的取值。此参数的配置格式“az(xx)”。()内为某个AZ的flavor状态，()内必须要填有状态，不填为无效配置。状态的取值范围与condOperationStatus参数相同。 **取值范围**： 不涉及
	CondOperationAz *string `json:"condOperationAz,omitempty"`

	// **参数解释**： 是否本地盘。 **取值范围**： - true：本地盘。 - false：非本地盘。
	Localdisk *string `json:"localdisk,omitempty"`

	// **参数解释**： 中文规格分类。 **取值范围**： 不涉及
	FlavorTypeCn *string `json:"flavorTypeCn,omitempty"`

	// **参数解释**： 英文规格分类。 **取值范围**： 不涉及
	FlavorTypeEn *string `json:"flavorTypeEn,omitempty"`

	// **参数解释**： 是否边缘区规格。 **取值范围**： - true：边缘区专属规格。 - false：通用规格。
	Edge *bool `json:"edge,omitempty"`

	// **参数解释**： 规格id。 **取值范围**： 不涉及
	StrId *string `json:"str_id,omitempty"`

	// **参数解释**： 节点类型是否支持HTTPS协议访问。 **取值范围**： 不涉及
	IsAllowHttps *bool `json:"isAllowHttps,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
