package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowGallerySubscription Gallery 订阅信息。
type WorkflowGallerySubscription struct {

	// 资产ID。
	ContentId *string `json:"content_id,omitempty"`

	// 版本ID。
	VersionId *string `json:"version_id,omitempty"`

	// 超期时间。
	ExpiredAt *string `json:"expired_at,omitempty"`
}

func (o WorkflowGallerySubscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowGallerySubscription struct{}"
	}

	return strings.Join([]string{"WorkflowGallerySubscription", string(data)}, " ")
}
