package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModelStatisticVo struct {
	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	Level *ModelLevel `json:"level,omitempty"`

	// 数据库。
	Db *int32 `json:"db,omitempty"`

	// 数据表。
	Tb *int32 `json:"tb,omitempty"`

	// 已发布的数据表。
	TbPublished *int32 `json:"tb_published,omitempty"`

	// 字段。
	Fd *int32 `json:"fd,omitempty"`

	// 已发布的字段。
	FdPublished *int32 `json:"fd_published,omitempty"`

	// 标准覆盖率。
	St *float64 `json:"st,omitempty"`

	// 已发布的标准覆盖率。
	StPublished *float64 `json:"st_published,omitempty"`

	Model *WorkspaceVo `json:"model,omitempty"`
}

func (o ModelStatisticVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelStatisticVo struct{}"
	}

	return strings.Join([]string{"ModelStatisticVo", string(data)}, " ")
}
