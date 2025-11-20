package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshPublicationSnapshotRequest Request Object
type RefreshPublicationSnapshotRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 发布ID
	PublicationId string `json:"publication_id"`
}

func (o RefreshPublicationSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshPublicationSnapshotRequest struct{}"
	}

	return strings.Join([]string{"RefreshPublicationSnapshotRequest", string(data)}, " ")
}
