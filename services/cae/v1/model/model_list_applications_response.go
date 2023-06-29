package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationsResponse Response Object
type ListApplicationsResponse struct {

	// API版本。
	ApiVersion *string `json:"api_version,omitempty"`

	// 资源种类。
	Kind *string `json:"kind,omitempty"`

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
