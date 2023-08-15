package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagResponse Response Object
type DeleteTagResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteTagResponse", string(data)}, " ")
}
