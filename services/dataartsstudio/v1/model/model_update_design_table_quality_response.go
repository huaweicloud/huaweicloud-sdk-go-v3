package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignTableQualityResponse Response Object
type UpdateDesignTableQualityResponse struct {
	Data           *UpdateDesignTableQualityResultData `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o UpdateDesignTableQualityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignTableQualityResponse struct{}"
	}

	return strings.Join([]string{"UpdateDesignTableQualityResponse", string(data)}, " ")
}
