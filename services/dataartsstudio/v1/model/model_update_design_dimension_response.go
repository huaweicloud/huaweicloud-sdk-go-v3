package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignDimensionResponse Response Object
type UpdateDesignDimensionResponse struct {
	Data           *ShowDimensionByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o UpdateDesignDimensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignDimensionResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignDimensionResponse", string(data)}, " ")
}
