package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableAccessPointResponse Response Object
type DisableAccessPointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAccessPointResponse struct{}"
	}

	return strings.Join([]string{"DisableAccessPointResponse", string(data)}, " ")
}
