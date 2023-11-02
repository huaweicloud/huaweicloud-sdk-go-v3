package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigHistoryDetailResponse Response Object
type ShowConfigHistoryDetailResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowConfigHistoryDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigHistoryDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigHistoryDetailResponse", string(data)}, " ")
}
