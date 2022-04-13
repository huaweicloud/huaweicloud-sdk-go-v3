package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowReferResponse struct {
	Referer        *RefererRsp `json:"referer,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowReferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReferResponse struct{}"
	}

	return strings.Join([]string{"ShowReferResponse", string(data)}, " ")
}
