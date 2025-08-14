package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProfilesResponse Response Object
type ListProfilesResponse struct {

	// 应用程序Profile
	ApplicationProfiles *[]ApplicationProfileDto `json:"applicationProfiles,omitempty"`
	HttpStatusCode      int                      `json:"-"`
}

func (o ListProfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProfilesResponse struct{}"
	}

	return strings.Join([]string{"ListProfilesResponse", string(data)}, " ")
}
