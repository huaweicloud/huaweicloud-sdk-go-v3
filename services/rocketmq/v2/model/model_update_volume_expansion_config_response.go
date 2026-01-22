package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVolumeExpansionConfigResponse Response Object
type UpdateVolumeExpansionConfigResponse struct {

	// **参数解释**： 修改结果。 **取值范围**： success。
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVolumeExpansionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeExpansionConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateVolumeExpansionConfigResponse", string(data)}, " ")
}
