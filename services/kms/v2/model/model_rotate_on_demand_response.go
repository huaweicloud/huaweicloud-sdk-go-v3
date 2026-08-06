package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateOnDemandResponse Response Object
type RotateOnDemandResponse struct {

	// **参数解释：** 密钥ID **取值范围：** 不涉及
	KeyId          *string `json:"key_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RotateOnDemandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateOnDemandResponse struct{}"
	}

	return strings.Join([]string{"RotateOnDemandResponse", string(data)}, " ")
}
