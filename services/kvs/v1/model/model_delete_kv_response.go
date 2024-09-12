package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKvResponse Response Object
type DeleteKvResponse struct {
	HttpStatusCode int `bson:"-"`
}

func (o DeleteKvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKvResponse struct{}"
	}

	return strings.Join([]string{"DeleteKvResponse", string(data)}, " ")
}
