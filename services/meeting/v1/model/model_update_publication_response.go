package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePublicationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePublicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicationResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicationResponse", string(data)}, " ")
}
