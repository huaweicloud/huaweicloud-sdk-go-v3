package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocumentVersionVo struct {

	// 版本号，如v1
	Version *string `json:"version,omitempty"`

	// 版本id
	VersionUuid *string `json:"version_uuid,omitempty"`

	// 版本创建时间
	CreateTime *string `json:"create_time,omitempty"`
}

func (o DocumentVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentVersionVo struct{}"
	}

	return strings.Join([]string{"DocumentVersionVo", string(data)}, " ")
}
