package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoJobResponse Response Object
type UpdateAutoJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAutoJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateAutoJobResponse", string(data)}, " ")
}
