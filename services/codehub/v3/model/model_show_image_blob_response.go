package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowImageBlobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowImageBlobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageBlobResponse struct{}"
	}

	return strings.Join([]string{"ShowImageBlobResponse", string(data)}, " ")
}
