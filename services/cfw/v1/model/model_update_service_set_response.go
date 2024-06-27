package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceSetResponse Response Object
type UpdateServiceSetResponse struct {
	Data           *CommonResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateServiceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSetResponse struct{}"
	}

	return strings.Join([]string{"UpdateServiceSetResponse", string(data)}, " ")
}
