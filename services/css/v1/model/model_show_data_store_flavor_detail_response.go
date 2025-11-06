package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataStoreFlavorDetailResponse Response Object
type ShowDataStoreFlavorDetailResponse struct {

	// **参数解释**： 引擎类型id。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 引擎名称。 **取值范围**： 不涉及
	Dbname *string `json:"dbname,omitempty"`

	// **参数解释**： 引擎版本。 **取值范围**： 不涉及
	Versions       *[]FlavorRespVersionBody `json:"versions,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowDataStoreFlavorDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataStoreFlavorDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDataStoreFlavorDetailResponse", string(data)}, " ")
}
