package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMembersDetailRequest struct {

	// 备份id
	BackupId string `json:"backup_id" xml:"backup_id"`

	// 接受备份共享的项目id
	DestProjectId *string `json:"dest_project_id,omitempty" xml:"dest_project_id"`

	// 接受的共享备份副本注册的镜像id
	ImageId *string `json:"image_id,omitempty" xml:"image_id"`

	// 备份共享状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 目标端接受共享备份的存储库id，仅支持uuid
	VaultId *string `json:"vault_id,omitempty" xml:"vault_id"`

	// 每页显示的条目数量，正整数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 上一次查询最后一条的id，仅支持uuid
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 偏移值，正整数
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// sort的内容为一组由逗号分隔的属性及可选排序方向组成，形如<key1>[:<direction>],<key2>[:<direction>],其中direction的取值为asc (升序) 或 desc (降序),如没有传入direction参数，默认为降序，sort内容的长度限制为255个字符。
	Sort *string `json:"sort,omitempty" xml:"sort"`
}

func (o ShowMembersDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMembersDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMembersDetailRequest", string(data)}, " ")
}
