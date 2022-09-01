package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupRsp struct {

	// 快照id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 快照名称。
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o BackupRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRsp struct{}"
	}

	return strings.Join([]string{"BackupRsp", string(data)}, " ")
}
