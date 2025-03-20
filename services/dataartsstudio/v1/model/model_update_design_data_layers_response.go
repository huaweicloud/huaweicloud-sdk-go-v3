package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignDataLayersResponse Response Object
type UpdateDesignDataLayersResponse struct {
	Data           *ListDesignDataLayersResultData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o UpdateDesignDataLayersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignDataLayersResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignDataLayersResponse", string(data)}, " ")
}
