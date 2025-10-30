package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSaleCyclesResponse Response Object
type ListSaleCyclesResponse struct {
	SaleCycles     *[]SaleCycleOption `json:"sale_cycles,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSaleCyclesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSaleCyclesResponse struct{}"
	}

	return strings.Join([]string{"ListSaleCyclesResponse", string(data)}, " ")
}
