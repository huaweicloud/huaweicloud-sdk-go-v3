package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNamespaceRequestBody struct {
	Metadata *NamespaceMetadata `json:"metadata"`
}

func (o UpdateNamespaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNamespaceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNamespaceRequestBody", string(data)}, " ")
}
