package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterServerResponse Response Object
type RegisterServerResponse struct {

	// 源端ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerResponse struct{}"
	}

	return strings.Join([]string{"RegisterServerResponse", string(data)}, " ")
}
