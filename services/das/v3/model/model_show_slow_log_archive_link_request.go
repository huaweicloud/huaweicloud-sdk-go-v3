package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogArchiveLinkRequest Request Object
type ShowSlowLogArchiveLinkRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 存档ID
	ArchiveId int64 `json:"archive_id"`
}

func (o ShowSlowLogArchiveLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogArchiveLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogArchiveLinkRequest", string(data)}, " ")
}
