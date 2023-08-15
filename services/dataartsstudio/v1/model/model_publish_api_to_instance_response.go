package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishApiToInstanceResponse Response Object
type PublishApiToInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PublishApiToInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApiToInstanceResponse struct{}"
	}

	return strings.Join([]string{"PublishApiToInstanceResponse", string(data)}, " ")
}
