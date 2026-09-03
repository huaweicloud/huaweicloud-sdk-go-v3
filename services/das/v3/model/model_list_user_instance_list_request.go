package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserInstanceListRequest Request Object
type ListUserInstanceListRequest struct {
	Body *ListUserInstanceListRequestBody `json:"body,omitempty"`
}

func (o ListUserInstanceListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserInstanceListRequest struct{}"
	}

	return strings.Join([]string{"ListUserInstanceListRequest", string(data)}, " ")
}
