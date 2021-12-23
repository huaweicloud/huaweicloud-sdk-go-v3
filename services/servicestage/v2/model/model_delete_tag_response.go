package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteTagResponse", string(data)}, " ")
}
