package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostGroupRequest Request Object
type ListHostGroupRequest struct {
	Body *GetHostGroupListRequestBody `json:"body,omitempty"`
}

func (o ListHostGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupRequest struct{}"
	}

	return strings.Join([]string{"ListHostGroupRequest", string(data)}, " ")
}
