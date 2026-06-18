package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantEncryptedRepositoriesResponse Response Object
type ListTenantEncryptedRepositoriesResponse struct {
	Body           *[]RepoEncryptionDto `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListTenantEncryptedRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantEncryptedRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantEncryptedRepositoriesResponse", string(data)}, " ")
}
