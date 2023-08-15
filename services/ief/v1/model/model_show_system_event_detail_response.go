package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSystemEventDetailResponse Response Object
type ShowSystemEventDetailResponse struct {
	Systemevent    *Event `json:"systemevent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSystemEventDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSystemEventDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSystemEventDetailResponse", string(data)}, " ")
}
