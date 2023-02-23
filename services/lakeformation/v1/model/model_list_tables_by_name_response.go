package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTablesByNameResponse struct {
	Body           *[]Table `json:"body,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListTablesByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesByNameResponse struct{}"
	}

	return strings.Join([]string{"ListTablesByNameResponse", string(data)}, " ")
}
