package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
