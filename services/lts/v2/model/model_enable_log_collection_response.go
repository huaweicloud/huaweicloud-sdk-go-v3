package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLogCollectionResponse Response Object
type EnableLogCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableLogCollectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLogCollectionResponse struct{}"
	}

	return strings.Join([]string{"EnableLogCollectionResponse", string(data)}, " ")
}
