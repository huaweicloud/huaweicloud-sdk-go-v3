package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEipCountResponse Response Object
type ListEipCountResponse struct {
	Data           *EipCountRespData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListEipCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEipCountResponse struct{}"
	}

	return strings.Join([]string{"ListEipCountResponse", string(data)}, " ")
}
