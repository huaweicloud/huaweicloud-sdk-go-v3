package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantRepositoriesResponse Response Object
type ListTenantRepositoriesResponse struct {

	// 仓库列表
	Body *[]TenantRepositoryDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTenantRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantRepositoriesResponse", string(data)}, " ")
}
