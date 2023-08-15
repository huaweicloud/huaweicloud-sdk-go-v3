package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlavorResponse Response Object
type UpdateFlavorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlavorResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlavorResponse", string(data)}, " ")
}
