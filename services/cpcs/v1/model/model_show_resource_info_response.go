package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceInfoResponse Response Object
type ShowResourceInfoResponse struct {
	CloudService *CloudServiceInfo `json:"cloud_service,omitempty"`

	CcspService *CcspServiceInfo `json:"ccsp_service,omitempty"`

	Vsm *VsmResourceInfo `json:"vsm,omitempty"`

	App *AppResourceInfo `json:"app,omitempty"`

	Kms            *KmsResourceInfo `json:"kms,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowResourceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceInfoResponse", string(data)}, " ")
}
