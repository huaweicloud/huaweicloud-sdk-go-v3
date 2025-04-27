package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePtrResponse Response Object
type DeletePtrResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePtrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePtrResponse struct{}"
	}

	return strings.Join([]string{"DeletePtrResponse", string(data)}, " ")
}
