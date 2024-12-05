package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignDimensionResponse Response Object
type CreateDesignDimensionResponse struct {
	Data           *ShowDimensionByIdResultData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CreateDesignDimensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignDimensionResponse struct{}"
	}

	return strings.Join([]string{"CreateDesignDimensionResponse", string(data)}, " ")
}
