package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandResponse Response Object
type HandResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o HandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandResponse struct{}"
	}

	return strings.Join([]string{"HandResponse", string(data)}, " ")
}
