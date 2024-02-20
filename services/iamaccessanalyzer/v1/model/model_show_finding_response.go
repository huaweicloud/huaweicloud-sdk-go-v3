package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFindingResponse Response Object
type ShowFindingResponse struct {
	Finding        *Finding `json:"finding,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowFindingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFindingResponse struct{}"
	}

	return strings.Join([]string{"ShowFindingResponse", string(data)}, " ")
}
