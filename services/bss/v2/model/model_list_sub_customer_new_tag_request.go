package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerNewTagRequest Request Object
type ListSubCustomerNewTagRequest struct {
	Body *ListSubCustomerNewTagReq `json:"body,omitempty"`
}

func (o ListSubCustomerNewTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerNewTagRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerNewTagRequest", string(data)}, " ")
}
