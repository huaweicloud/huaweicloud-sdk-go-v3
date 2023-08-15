package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBlackWhiteListUsingPutResponse Response Object
type UpdateBlackWhiteListUsingPutResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateBlackWhiteListUsingPutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListUsingPutResponse struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListUsingPutResponse", string(data)}, " ")
}
