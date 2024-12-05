package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignDimensionResponse Response Object
type DeleteDesignDimensionResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignDimensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignDimensionResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignDimensionResponse", string(data)}, " ")
}
