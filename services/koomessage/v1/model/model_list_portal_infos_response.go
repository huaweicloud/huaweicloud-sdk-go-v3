package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPortalInfosResponse struct {
	Data           *ListPortalInfosResponseModel `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPortalInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortalInfosResponse struct{}"
	}

	return strings.Join([]string{"ListPortalInfosResponse", string(data)}, " ")
}
