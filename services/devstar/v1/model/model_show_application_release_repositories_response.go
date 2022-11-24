package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowApplicationReleaseRepositoriesResponse struct {

	// 软件包列表
	ReleaseRepositories *[]ReleaseRepository `json:"release_repositories,omitempty"`

	// 软件包总条数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowApplicationReleaseRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicationReleaseRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ShowApplicationReleaseRepositoriesResponse", string(data)}, " ")
}
