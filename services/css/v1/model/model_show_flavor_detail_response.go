package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlavorDetailResponse Response Object
type ShowFlavorDetailResponse struct {

	// **参数解释**： 规格id。 **取值范围**： 不涉及
	StrId *string `json:"str_id,omitempty"`

	// **参数解释**： 规格名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 此参数是Region级配置，某个AZ没有在condOperationAz参数中配置时默认使用此参数的取值。 **取值范围**：   - normal：正常商用。   - sellout：售罄。
	CondOperationStatus *string `json:"condOperationStatus,omitempty"`

	// **参数解释**： 此参数是AZ级配置，某个AZ没有在此参数中配置时默认使用condOperationAz参数的取值。此参数的配置格式“az(xx)”。()内为某个AZ的flavor状态，()内必须要填有状态，不填为无效配置。状态的取值范围与condOperationStatus参数相同。 **取值范围**： 不涉及
	CondOperationAz *string `json:"condOperationAz,omitempty"`

	// **参数解释**： 规格属性信息。 **取值范围**： 不涉及
	FlavorDetail   *[]ShowFlavorDetailRspFlavorDetail `json:"flavor_detail,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ShowFlavorDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorDetailResponse", string(data)}, " ")
}
