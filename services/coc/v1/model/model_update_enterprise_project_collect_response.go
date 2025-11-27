package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseProjectCollectResponse Response Object
type UpdateEnterpriseProjectCollectResponse struct {

	// **参数解释：** CloudCMDB分配的uuid。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnterpriseProjectCollectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectCollectResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectCollectResponse", string(data)}, " ")
}
