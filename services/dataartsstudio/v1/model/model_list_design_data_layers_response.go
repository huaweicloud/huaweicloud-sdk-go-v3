package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesignDataLayersResponse Response Object
type ListDesignDataLayersResponse struct {
	Data           *ListDesignDataLayersResultData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListDesignDataLayersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesignDataLayersResponse struct{}"
	}

	return strings.Join([]string{"ListDesignDataLayersResponse", string(data)}, " ")
}
