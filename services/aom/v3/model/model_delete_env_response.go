package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnvResponse Response Object
type DeleteEnvResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnvResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnvResponse", string(data)}, " ")
}
