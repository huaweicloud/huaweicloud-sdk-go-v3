package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVolumesResponse struct {
	// 查询到的云硬盘总数量，不受分页影响。

	Count *int32 `json:"count,omitempty"`
	// 云硬盘列表查询位置标记。如果本次查询只返回部分列表信息时，会返回查询到的当前磁盘mark标记的url，可以继续使用这个url查询剩余列表信息。

	VolumesLinks *[]Link `json:"volumes_links,omitempty"`
	// 查询请求返回的云硬盘列表。

	Volumes        *[]VolumeDetail `json:"volumes,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesResponse", string(data)}, " ")
}
