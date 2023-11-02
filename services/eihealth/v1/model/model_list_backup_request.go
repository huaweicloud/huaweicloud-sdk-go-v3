package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupRequest Request Object
type ListBackupRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为desc）
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段（支持type, end_time）
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupRequest struct{}"
	}

	return strings.Join([]string{"ListBackupRequest", string(data)}, " ")
}
