package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowServiceDetailResponse struct {
	Service        *ServiceRespDetail `json:"service,omitempty" xml:"service"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowServiceDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowServiceDetailResponse", string(data)}, " ")
}
