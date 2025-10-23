package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrganizationPolicyResponse Response Object
type ShowOrganizationPolicyResponse struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	Detail         *BccOrganizationPolicyDetail `json:"detail,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowOrganizationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationPolicyResponse", string(data)}, " ")
}
