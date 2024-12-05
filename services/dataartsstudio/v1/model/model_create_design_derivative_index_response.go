package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignDerivativeIndexResponse Response Object
type CreateDesignDerivativeIndexResponse struct {
	Data           *CreateDesignDerivativeIndexResultData `json:"data,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o CreateDesignDerivativeIndexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignDerivativeIndexResponse struct{}"
	}

	return strings.Join([]string{"CreateDesignDerivativeIndexResponse", string(data)}, " ")
}
