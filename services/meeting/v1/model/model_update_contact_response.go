package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateContactResponse Response Object
type UpdateContactResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateContactResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateContactResponse struct{}"
	}

	return strings.Join([]string{"UpdateContactResponse", string(data)}, " ")
}
