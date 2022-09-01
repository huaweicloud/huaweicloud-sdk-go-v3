package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSpaceAnalysisResponse struct {

	// 记录总数
	Total *int64 `json:"total,omitempty" xml:"total"`

	// 数据库对象列表
	DbObjects *[]DbObjectSpaceInfo `json:"db_objects,omitempty" xml:"db_objects"`

	InstanceInfo   *InstanceSpaceInfo `json:"instance_info,omitempty" xml:"instance_info"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSpaceAnalysisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpaceAnalysisResponse struct{}"
	}

	return strings.Join([]string{"ListSpaceAnalysisResponse", string(data)}, " ")
}
