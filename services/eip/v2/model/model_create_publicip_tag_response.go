package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePublicipTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePublicipTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePublicipTagResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicipTagResponse", string(data)}, " ")
}
