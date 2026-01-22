package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteServiceSetsRespData struct {

	// **参数解释**： 服务组名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 服务组ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`
}

func (o BatchDeleteServiceSetsRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServiceSetsRespData struct{}"
	}

	return strings.Join([]string{"BatchDeleteServiceSetsRespData", string(data)}, " ")
}
