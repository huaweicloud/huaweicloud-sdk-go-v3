package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableModelByIdResponse Response Object
type ShowTableModelByIdResponse struct {
	Data           *CreateTableModelResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowTableModelByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableModelByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTableModelByIdResponse", string(data)}, " ")
}
