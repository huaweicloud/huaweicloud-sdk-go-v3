package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGcbResourceTagResponse Response Object
type DeleteGcbResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGcbResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGcbResourceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteGcbResourceTagResponse", string(data)}, " ")
}
