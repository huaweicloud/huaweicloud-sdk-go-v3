package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceIdsResponse Response Object
type ListResourceIdsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ListResourceIdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceIdsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceIdsResponse", string(data)}, " ")
}
