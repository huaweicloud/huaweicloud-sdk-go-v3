package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReferResponse Response Object
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
