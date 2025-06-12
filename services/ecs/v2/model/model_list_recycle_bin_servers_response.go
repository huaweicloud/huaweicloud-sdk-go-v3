package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleBinServersResponse Response Object
type ListRecycleBinServersResponse struct {
	Servers        *[]ServerDetail `json:"servers,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListRecycleBinServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleBinServersResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleBinServersResponse", string(data)}, " ")
}
