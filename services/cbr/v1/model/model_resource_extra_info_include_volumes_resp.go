package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceExtraInfoIncludeVolumesResp struct {

	// **参数解释：** 备份时，需要备份的卷的ID **取值范围：** 格式为uuid
	Id string `json:"id"`

	// **参数解释：** 操作系统类型 **取值范围：** 不涉及
	OsVersion *string `json:"os_version,omitempty"`
}

func (o ResourceExtraInfoIncludeVolumesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfoIncludeVolumesResp struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfoIncludeVolumesResp", string(data)}, " ")
}
