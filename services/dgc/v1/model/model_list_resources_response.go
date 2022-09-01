package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourcesResponse struct {
	Total *int32 `json:"total,omitempty" xml:"total"`

	Resources      *[]ResourceInfo `json:"resources,omitempty" xml:"resources"`
	HttpStatusCode int             `json:"-"`
}

func (o ListResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesResponse", string(data)}, " ")
}
