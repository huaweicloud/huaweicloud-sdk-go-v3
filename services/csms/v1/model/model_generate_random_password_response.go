package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateRandomPasswordResponse Response Object
type GenerateRandomPasswordResponse struct {
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GenerateRandomPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateRandomPasswordResponse struct{}"
	}

	return strings.Join([]string{"GenerateRandomPasswordResponse", string(data)}, " ")
}
