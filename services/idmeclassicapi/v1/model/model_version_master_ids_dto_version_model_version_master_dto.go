package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionMasterIdsDtoVersionModelVersionMasterDto struct {

	// 主对象集合。
	MasterIds *[]VersionModelMasterIdsDto `json:"masterIds,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionMasterIdsDtoVersionModelVersionMasterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionMasterIdsDtoVersionModelVersionMasterDto struct{}"
	}

	return strings.Join([]string{"VersionMasterIdsDtoVersionModelVersionMasterDto", string(data)}, " ")
}
