package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAuthorizeInterfaceResponse Response Object
type BatchAuthorizeInterfaceResponse struct {

	// lakecat策略信息
	Policies *[]LakeFormationPolicy `json:"policies,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o BatchAuthorizeInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAuthorizeInterfaceResponse struct{}"
	}

	return strings.Join([]string{"BatchAuthorizeInterfaceResponse", string(data)}, " ")
}
