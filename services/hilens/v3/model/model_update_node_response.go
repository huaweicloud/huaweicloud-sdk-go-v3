package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeResponse Response Object
type UpdateNodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeResponse", string(data)}, " ")
}
