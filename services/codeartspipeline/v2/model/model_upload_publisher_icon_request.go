package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadPublisherIconRequest Request Object
type UploadPublisherIconRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 发布商名称
	PublisherEnName string `json:"publisher_en_name"`

	Body *UploadPublisherIconRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadPublisherIconRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadPublisherIconRequest struct{}"
	}

	return strings.Join([]string{"UploadPublisherIconRequest", string(data)}, " ")
}
