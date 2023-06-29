package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandInstanceStorageResponse Response Object
type ExpandInstanceStorageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExpandInstanceStorageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceStorageResponse struct{}"
	}

	return strings.Join([]string{"ExpandInstanceStorageResponse", string(data)}, " ")
}
