package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipResourcesResponse Response Object
type ListEipResourcesResponse struct {
	Data           *EipResponseData `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListEipResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListEipResourcesResponse", string(data)}, " ")
}
