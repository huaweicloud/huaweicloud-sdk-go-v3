package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVendorAccountResponse Response Object
type DeleteVendorAccountResponse struct {

	// **参数解释：** CloudCMDB分配的uuid。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteVendorAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVendorAccountResponse struct{}"
	}

	return strings.Join([]string{"DeleteVendorAccountResponse", string(data)}, " ")
}
