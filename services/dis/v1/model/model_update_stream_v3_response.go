package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateStreamV3Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStreamV3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStreamV3Response struct{}"
	}

	return strings.Join([]string{"UpdateStreamV3Response", string(data)}, " ")
}
