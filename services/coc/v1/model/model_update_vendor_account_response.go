package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVendorAccountResponse Response Object
type UpdateVendorAccountResponse struct {

	// **参数解释：** CloudCMDB分配的uuid。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVendorAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVendorAccountResponse struct{}"
	}

	return strings.Join([]string{"UpdateVendorAccountResponse", string(data)}, " ")
}
