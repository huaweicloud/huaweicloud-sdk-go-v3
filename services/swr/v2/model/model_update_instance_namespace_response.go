package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceNamespaceResponse Response Object
type UpdateInstanceNamespaceResponse struct {

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	Metadata *NamespaceMetadata `json:"metadata,omitempty"`

	CveAllowlist   *UpdateCveAllowlistRequest `json:"cve_allowlist,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o UpdateInstanceNamespaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNamespaceResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNamespaceResponse", string(data)}, " ")
}
