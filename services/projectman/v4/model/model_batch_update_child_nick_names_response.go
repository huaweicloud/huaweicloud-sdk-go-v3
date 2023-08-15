package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateChildNickNamesResponse Response Object
type BatchUpdateChildNickNamesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateChildNickNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateChildNickNamesResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateChildNickNamesResponse", string(data)}, " ")
}
