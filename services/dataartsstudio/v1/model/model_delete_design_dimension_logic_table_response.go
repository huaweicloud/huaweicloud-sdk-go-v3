package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignDimensionLogicTableResponse Response Object
type DeleteDesignDimensionLogicTableResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignDimensionLogicTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignDimensionLogicTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignDimensionLogicTableResponse", string(data)}, " ")
}
