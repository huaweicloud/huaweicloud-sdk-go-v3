package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableLogCollectionResponse Response Object
type DisableLogCollectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableLogCollectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableLogCollectionResponse struct{}"
	}

	return strings.Join([]string{"DisableLogCollectionResponse", string(data)}, " ")
}
