package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTagResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTagResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteTagResourceResponse", string(data)}, " ")
}
