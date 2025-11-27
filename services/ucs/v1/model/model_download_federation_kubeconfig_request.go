package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadFederationKubeconfigRequest Request Object
type DownloadFederationKubeconfigRequest struct {

	// 舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *DownloadFederationKubeconfigRequestBody `json:"body,omitempty"`
}

func (o DownloadFederationKubeconfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFederationKubeconfigRequest struct{}"
	}

	return strings.Join([]string{"DownloadFederationKubeconfigRequest", string(data)}, " ")
}
