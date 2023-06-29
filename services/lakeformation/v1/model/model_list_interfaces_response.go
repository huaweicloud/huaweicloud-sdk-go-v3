package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInterfacesResponse Response Object
type ListInterfacesResponse struct {

	// lakecat策略信息
	Policies *[]LakeFormationPolicy `json:"policies,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListInterfacesResponse", string(data)}, " ")
}
