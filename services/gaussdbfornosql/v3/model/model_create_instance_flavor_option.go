package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例规格详情。
type CreateInstanceFlavorOption struct {

	// 节点数量。   - GaussDB(for Cassandra)实例的节点数量可取3~12。   - GaussDB(for Mongo)4.0版本副本集实例的节点数量可取3。   - GaussDB(for Influx)实例的节点数量可取3~16。
	Num string `json:"num" xml:"num"`

	// 磁盘类型。 取值为“ULTRAHIGH”，表示SSD盘。
	Storage string `json:"storage" xml:"storage"`

	// 磁盘大小。必须为10的整数倍。单位为GB。最小磁盘容量100GB，最大磁盘容量与实例的性能规格有关，详见数据库实例规格。
	Size string `json:"size" xml:"size"`

	// 资源规格编码。获取方法请参见查询所有实例规格信息中响应参数“spec_code”的值。
	SpecCode string `json:"spec_code" xml:"spec_code"`
}

func (o CreateInstanceFlavorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceFlavorOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceFlavorOption", string(data)}, " ")
}
