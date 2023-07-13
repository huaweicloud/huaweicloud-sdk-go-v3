package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCorpResponse Response Object
type UpdateCorpResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCorpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCorpResponse struct{}"
	}

	return strings.Join([]string{"UpdateCorpResponse", string(data)}, " ")
}
