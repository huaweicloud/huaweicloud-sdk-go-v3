package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteConfigMapResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConfigMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigMapResponse struct{}"
	}

	return strings.Join([]string{"DeleteConfigMapResponse", string(data)}, " ")
}
