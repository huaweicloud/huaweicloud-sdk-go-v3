package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAppDetailResponse struct {
	App            *AppResp `json:"app,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowAppDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAppDetailResponse", string(data)}, " ")
}
