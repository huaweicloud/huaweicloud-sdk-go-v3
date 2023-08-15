package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSetDetailsResponse Response Object
type ListServiceSetDetailsResponse struct {
	Data           *ServiceSetDetailResponseDto `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListServiceSetDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceSetDetailsResponse", string(data)}, " ")
}
