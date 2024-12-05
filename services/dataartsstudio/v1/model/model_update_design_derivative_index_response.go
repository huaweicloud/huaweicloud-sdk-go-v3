package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignDerivativeIndexResponse Response Object
type UpdateDesignDerivativeIndexResponse struct {
	Data           *UpdateDesignDerivativeIndexResultData `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o UpdateDesignDerivativeIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignDerivativeIndexResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignDerivativeIndexResponse", string(data)}, " ")
}
