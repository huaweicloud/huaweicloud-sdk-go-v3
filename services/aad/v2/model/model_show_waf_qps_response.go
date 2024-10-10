package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWafQpsResponse Response Object
type ShowWafQpsResponse struct {

	// qps
	Qps            *int32 `json:"qps,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWafQpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWafQpsResponse struct{}"
	}

	return strings.Join([]string{"ShowWafQpsResponse", string(data)}, " ")
}
