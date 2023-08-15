package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStorageClaimRequest Request Object
type DeleteStorageClaimRequest struct {

	// WKS存储ID
	StorageId string `json:"storage_id"`

	Body *DeleteStorageClaimReq `json:"body,omitempty"`
}

func (o DeleteStorageClaimRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStorageClaimRequest struct{}"
	}

	return strings.Join([]string{"DeleteStorageClaimRequest", string(data)}, " ")
}
