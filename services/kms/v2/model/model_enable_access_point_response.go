package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAccessPointResponse Response Object
type EnableAccessPointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAccessPointResponse struct{}"
	}

	return strings.Join([]string{"EnableAccessPointResponse", string(data)}, " ")
}
