package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPrivateCorpDigitalInfoResponse Response Object
type SearchPrivateCorpDigitalInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SearchPrivateCorpDigitalInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPrivateCorpDigitalInfoResponse struct{}"
	}

	return strings.Join([]string{"SearchPrivateCorpDigitalInfoResponse", string(data)}, " ")
}
