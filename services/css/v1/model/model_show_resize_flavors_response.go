package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResizeFlavorsResponse Response Object
type ShowResizeFlavorsResponse struct {

	// **参数解释**： 引擎类型id。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 引擎名称。 **取值范围**： 不涉及
	Dbname *string `json:"dbname,omitempty"`

	// **参数解释**： 版本列表。 **取值范围**： 不涉及
	Versions       *[]ResizeFlavorRspVersionBody `json:"versions,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowResizeFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResizeFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowResizeFlavorsResponse", string(data)}, " ")
}
