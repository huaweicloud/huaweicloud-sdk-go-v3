package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterRequest Request Object
type ChangeClusterRequest struct {

	// WKS存储ID。
	StorageId string `json:"storage_id"`

	Body *ChangeClusterReq `json:"body,omitempty"`
}

func (o ChangeClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterRequest struct{}"
	}

	return strings.Join([]string{"ChangeClusterRequest", string(data)}, " ")
}
