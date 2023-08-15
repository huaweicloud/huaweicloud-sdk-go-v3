package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountEipsResponse Response Object
type CountEipsResponse struct {
	Data           *EipCountRespData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CountEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountEipsResponse struct{}"
	}

	return strings.Join([]string{"CountEipsResponse", string(data)}, " ")
}
