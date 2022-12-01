package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 专属分布式存储池详情。
type DssPool struct {

	// 专属分布式存储池名称。
	Id string `json:"id"`

	// 专属分布式存储池ID。
	Name string `json:"name"`

	// 专属分布式存储池的存储类型。 - SSD：超高IO专属分布式存储池。
	Type string `json:"type"`

	// 专属分布式存储池归属的project_id。
	ProjectId string `json:"project_id"`

	// 专属分布式存储池所属可用区。
	AvailabilityZone string `json:"availability_zone"`

	// 申请的专属分布式存储容量，单位TB。
	Capacity int32 `json:"capacity"`

	// 专属分布式存储池的状态。 - available：专属分布式存储池处于可用状态。 - deploying：专属分布式存储池处于正在部署的过程中，不可使用。 - extending：专属分布式存储池处于正在扩容的过程中，可使用。
	Status string `json:"status"`

	// 专属分布式存储池的创建时间。 - 时间格式：UTC YYYY-MM-DDTHH:MM:SS
	CreatedAt string `json:"created_at"`
}

func (o DssPool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DssPool struct{}"
	}

	return strings.Join([]string{"DssPool", string(data)}, " ")
}
