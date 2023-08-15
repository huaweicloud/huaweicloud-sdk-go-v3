package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelShareConnectionsResponse Response Object
type CancelShareConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelShareConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareConnectionsResponse struct{}"
	}

	return strings.Join([]string{"CancelShareConnectionsResponse", string(data)}, " ")
}
