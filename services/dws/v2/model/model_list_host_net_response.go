package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostNetResponse Response Object
type ListHostNetResponse struct {
	Body           *[]NetResp `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListHostNetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostNetResponse struct{}"
	}

	return strings.Join([]string{"ListHostNetResponse", string(data)}, " ")
}
