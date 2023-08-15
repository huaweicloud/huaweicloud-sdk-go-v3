package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishAppResponse Response Object
type UnpublishAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnpublishAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishAppResponse struct{}"
	}

	return strings.Join([]string{"UnpublishAppResponse", string(data)}, " ")
}
