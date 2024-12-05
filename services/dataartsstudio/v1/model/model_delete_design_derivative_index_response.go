package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesignDerivativeIndexResponse Response Object
type DeleteDesignDerivativeIndexResponse struct {
	Data           *DeleteResultData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteDesignDerivativeIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesignDerivativeIndexResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesignDerivativeIndexResponse", string(data)}, " ")
}
