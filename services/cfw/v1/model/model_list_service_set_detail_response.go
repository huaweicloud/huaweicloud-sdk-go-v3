package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSetDetailResponse Response Object
type ListServiceSetDetailResponse struct {
	Data           *ServiceSetDetailResponseDto `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListServiceSetDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetDetailResponse struct{}"
	}

	return strings.Join([]string{"ListServiceSetDetailResponse", string(data)}, " ")
}
