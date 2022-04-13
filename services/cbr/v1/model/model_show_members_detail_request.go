package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMembersDetailRequest struct {
	// 备份id

	BackupId string `json:"backup_id"`
	// 接受备份共享的项目id

	DestProjectId *string `json:"dest_project_id,omitempty"`
	// 接受的共享备份副本注册的镜像id

	ImageId *string `json:"image_id,omitempty"`
	// 备份共享状态

	Status *string `json:"status,omitempty"`
	// 目标端接受共享备份的存储库id，仅支持uuid

	VaultId *string `json:"vault_id,omitempty"`
	// 每页显示的条目数量，正整数

	Limit *int32 `json:"limit,omitempty"`
	// 上一次查询最后一条的id，仅支持uuid

	Marker *string `json:"marker,omitempty"`
	// 偏移值，正整数

	Offset *int32 `json:"offset,omitempty"`
	// sort的内容为一组由逗号分隔的属性及可选排序方向组成，形如<key1>[:<direction>],<key2>[:<direction>],其中direction的取值为asc (升序) 或 desc (降序),如没有传入direction参数，默认为降序，sort内容的长度限制为255个字符。

	Sort *string `json:"sort,omitempty"`
}

func (o ShowMembersDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMembersDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowMembersDetailRequest", string(data)}, " ")
}
