package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyInternetResponse Response Object
type ApplyInternetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ApplyInternetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyInternetResponse struct{}"
	}

	return strings.Join([]string{"ApplyInternetResponse", string(data)}, " ")
}
