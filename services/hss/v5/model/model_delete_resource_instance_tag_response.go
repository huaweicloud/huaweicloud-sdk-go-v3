package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceInstanceTagResponse Response Object
type DeleteResourceInstanceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceInstanceTagResponse", string(data)}, " ")
}
