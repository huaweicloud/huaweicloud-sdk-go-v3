package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowTrendResponse Response Object
type ShowFlowTrendResponse struct {
	Data           *ShowFlowTrendRespData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowFlowTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowTrendResponse", string(data)}, " ")
}
