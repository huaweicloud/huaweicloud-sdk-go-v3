package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetadataRequest Request Object
type ShowMetadataRequest struct {

	// 备份ID
	BackupId string `json:"backup_id"`
}

func (o ShowMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowMetadataRequest", string(data)}, " ")
}
