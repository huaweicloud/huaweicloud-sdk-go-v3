package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseAvailableVersionsResponse Response Object
type ListDatabaseAvailableVersionsResponse struct {

	// 可变更版本
	Versions *[]string `json:"versions,omitempty"`

	// 当前版本
	CurrentVersion *string `json:"current_version,omitempty"`

	// 最新优选版本
	LatestVersion *string `json:"latest_version,omitempty"`

	// 本系列优选版本，如3.0.8系列优选版本为3.0.8.5
	CurrentFavoredVersion *string `json:"current_favored_version,omitempty"`

	// 当前实例上一个版本
	PreviousVersion *string `json:"previous_version,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ListDatabaseAvailableVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseAvailableVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseAvailableVersionsResponse", string(data)}, " ")
}
