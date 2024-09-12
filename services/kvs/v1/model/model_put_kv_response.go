package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutKvResponse Response Object
type PutKvResponse struct {
	HttpStatusCode int `bson:"-"`
}

func (o PutKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutKvResponse struct{}"
	}

	return strings.Join([]string{"PutKvResponse", string(data)}, " ")
}
