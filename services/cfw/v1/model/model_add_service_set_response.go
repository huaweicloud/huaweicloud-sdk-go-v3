package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddServiceSetResponse Response Object
type AddServiceSetResponse struct {
	Data           *CommonResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AddServiceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceSetResponse struct{}"
	}

	return strings.Join([]string{"AddServiceSetResponse", string(data)}, " ")
}
