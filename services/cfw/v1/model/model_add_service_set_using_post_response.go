package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServiceSetUsingPostResponse Response Object
type AddServiceSetUsingPostResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddServiceSetUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceSetUsingPostResponse struct{}"
	}

	return strings.Join([]string{"AddServiceSetUsingPostResponse", string(data)}, " ")
}
