package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceJobResponse Response Object
type DeleteInstanceJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceJobResponse", string(data)}, " ")
}
