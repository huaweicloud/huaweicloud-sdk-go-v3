package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCorpDigitalInfoListResponse Response Object
type SearchCorpDigitalInfoListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SearchCorpDigitalInfoListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCorpDigitalInfoListResponse struct{}"
	}

	return strings.Join([]string{"SearchCorpDigitalInfoListResponse", string(data)}, " ")
}
