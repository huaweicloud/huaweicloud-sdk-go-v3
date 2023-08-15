package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageTagResponse Response Object
type DeleteImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImageTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageTagResponse", string(data)}, " ")
}
