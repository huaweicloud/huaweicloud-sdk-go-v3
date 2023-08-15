package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePublicipTagResponse Response Object
type DeletePublicipTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePublicipTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePublicipTagResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicipTagResponse", string(data)}, " ")
}
