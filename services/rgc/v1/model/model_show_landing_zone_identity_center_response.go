package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLandingZoneIdentityCenterResponse Response Object
type ShowLandingZoneIdentityCenterResponse struct {

	// Identity Center目录ID。
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// Identity Center登录URL。
	UserPortalUrl *string `json:"user_portal_url,omitempty"`

	PermissionSets *[]PermissionSet `json:"permission_sets,omitempty"`

	Groups         *[]IdentityCenterGroup `json:"groups,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowLandingZoneIdentityCenterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLandingZoneIdentityCenterResponse struct{}"
	}

	return strings.Join([]string{"ShowLandingZoneIdentityCenterResponse", string(data)}, " ")
}
