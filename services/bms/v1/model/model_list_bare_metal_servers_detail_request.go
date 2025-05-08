package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBareMetalServersDetailRequest Request Object
type ListBareMetalServersDetailRequest struct {

	// 裸金属服务器规格ID
	Flavor *string `json:"flavor,omitempty"`

	// 裸金属服务器名称
	Name *string `json:"name,omitempty"`

	// 裸金属服务器状态,只有管理员可以使用DELETED状态过滤查询已经删除的裸金属服务器。取值范围：ACTIVE、BUILD、ERROR、HARD_REBOOT、REBOOT、REBUILD、SHUTOFF
	Status *string `json:"status,omitempty"`

	// 每页返回裸金属服务器的条数，默认值是25，最大值为1000。limit为每页返回裸金属服务器详情的条数
	Limit *int32 `json:"limit,omitempty"`

	// 此接口为分页查询接口，offset为查询页码（起始页码为1），返回值包括总条数和裸金属服务器详情列表。传入offset：按limit值分页（limit默认为1000），返回第offset页裸金属服务器详情列表和总条数，总条数最大值为limit，不足按实际情况返回。不传入offset，传入limit：返回裸金属服务器详情列表和总条数，总条数最大值为limit，不足按实际情况返回。不传入offset，不传入limit：按25条分页，返回第1页裸金属服务器详情列表，总条数最大值为25，不足按实际情况返回。
	Offset *int32 `json:"offset,omitempty"`

	// 查询裸金属服务器结果的详细级别，级别越高，查询到的裸金属服务器信息越多，默认为4。可使用的级别为 1，2，3，4
	Detail *string `json:"detail,omitempty"`
}

func (o ListBareMetalServersDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBareMetalServersDetailRequest struct{}"
	}

	return strings.Join([]string{"ListBareMetalServersDetailRequest", string(data)}, " ")
}
