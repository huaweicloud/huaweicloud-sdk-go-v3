package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaDeleteKeypairResponse Response Object
type NovaDeleteKeypairResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDeleteKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDeleteKeypairResponse struct{}"
	}

	return strings.Join([]string{"NovaDeleteKeypairResponse", string(data)}, " ")
}
