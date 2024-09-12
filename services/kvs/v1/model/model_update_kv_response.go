package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKvResponse Response Object
type UpdateKvResponse struct {
	HttpStatusCode int `bson:"-"`
}

func (o UpdateKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKvResponse struct{}"
	}

	return strings.Join([]string{"UpdateKvResponse", string(data)}, " ")
}
