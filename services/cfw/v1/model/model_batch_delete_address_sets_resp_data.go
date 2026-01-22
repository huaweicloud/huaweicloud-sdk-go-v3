package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAddressSetsRespData struct {

	// **参数解释**： 地址组名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 地址组ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o BatchDeleteAddressSetsRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAddressSetsRespData struct{}"
	}

	return strings.Join([]string{"BatchDeleteAddressSetsRespData", string(data)}, " ")
}
