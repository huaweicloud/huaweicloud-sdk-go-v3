package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGraph2Response Response Object
type ShowGraph2Response struct {
	Graph          *ShowGraphRespGraph `json:"graph,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGraph2Response struct{}"
	}

	return strings.Join([]string{"ShowGraph2Response", string(data)}, " ")
}
