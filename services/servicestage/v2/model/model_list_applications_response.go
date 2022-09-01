package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicationsResponse struct {

	// 应用总数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 应用列表。
	Applications   *[]ApplicationView `json:"applications,omitempty" xml:"applications"`
	HttpStatusCode int                `json:"-"`
}

func (o ListApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationsResponse", string(data)}, " ")
}
