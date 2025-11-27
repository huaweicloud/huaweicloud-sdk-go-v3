package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVendorAccountResponse Response Object
type CreateVendorAccountResponse struct {

	// **参数解释：** CloudCMDB分配的uuid。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVendorAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVendorAccountResponse struct{}"
	}

	return strings.Join([]string{"CreateVendorAccountResponse", string(data)}, " ")
}
