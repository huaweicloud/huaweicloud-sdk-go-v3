package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyForAccessResponse Response Object
type ApplyForAccessResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ApplyForAccessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyForAccessResponse struct{}"
	}

	return strings.Join([]string{"ApplyForAccessResponse", string(data)}, " ")
}
