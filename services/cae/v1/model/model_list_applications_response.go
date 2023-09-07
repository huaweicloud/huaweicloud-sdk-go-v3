package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationsResponse Response Object
type ListApplicationsResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ApplicationKindObj `json:"kind,omitempty"`

	// 应用列表。
	Items          *[]ApplicationItem `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationsResponse", string(data)}, " ")
}
