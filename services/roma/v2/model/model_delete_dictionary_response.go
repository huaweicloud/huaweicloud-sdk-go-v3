package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDictionaryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDictionaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDictionaryResponse struct{}"
	}

	return strings.Join([]string{"DeleteDictionaryResponse", string(data)}, " ")
}
