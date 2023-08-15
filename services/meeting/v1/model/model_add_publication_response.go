package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPublicationResponse Response Object
type AddPublicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddPublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPublicationResponse struct{}"
	}

	return strings.Join([]string{"AddPublicationResponse", string(data)}, " ")
}
