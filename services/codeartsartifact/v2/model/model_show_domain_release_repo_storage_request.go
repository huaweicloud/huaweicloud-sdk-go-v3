package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainReleaseRepoStorageRequest Request Object
type ShowDomainReleaseRepoStorageRequest struct {
}

func (o ShowDomainReleaseRepoStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainReleaseRepoStorageRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainReleaseRepoStorageRequest", string(data)}, " ")
}
