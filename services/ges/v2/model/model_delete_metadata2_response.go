package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMetadata2Response Response Object
type DeleteMetadata2Response struct {
	Body           map[string]string `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteMetadata2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetadata2Response struct{}"
	}

	return strings.Join([]string{"DeleteMetadata2Response", string(data)}, " ")
}
