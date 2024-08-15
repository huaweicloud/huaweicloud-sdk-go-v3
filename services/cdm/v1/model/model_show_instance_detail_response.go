package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDetailResponse Response Object
type ShowInstanceDetailResponse struct {
	Instance       *CdmQueryClusterInstanceDetail `json:"instance,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowInstanceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDetailResponse", string(data)}, " ")
}
