package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessPointResponse Response Object
type DeleteAccessPointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessPointResponse struct{}"
	}

	return strings.Join([]string{"DeleteAccessPointResponse", string(data)}, " ")
}
