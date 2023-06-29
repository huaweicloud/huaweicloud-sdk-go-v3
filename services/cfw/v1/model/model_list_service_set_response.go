package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceSetResponse Response Object
type ListServiceSetResponse struct {
	Data           *ServiceSetRecords `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListServiceSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceSetResponse struct{}"
	}

	return strings.Join([]string{"ListServiceSetResponse", string(data)}, " ")
}
