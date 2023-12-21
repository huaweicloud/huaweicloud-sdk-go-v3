package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateClusterInstanceBody struct {

	// 集群节点类型，hbase有regionserver，hmaster，opentsdb等，doris有be，fe节点，clickhouse有server（计算节点），zookeeper
	InstanceType string `json:"instance_type"`

	// 节点个数，hbase取值：2<=num<=10],偶数 doris取值：be[3,100] fe只能是3或5， clickhouse取值：计算节点[2,10000],取偶数，zookeeper节点固定为3
	InstanceNum int32 `json:"instance_num"`

	// 节点规格，doris集群、clickhouse集群必选
	Flavor *string `json:"flavor,omitempty"`

	// 数据盘规格：COMMON、HIGH、ULTRAHIGH，NORMALHIGH, EXTREMEHIGH. doris集群、clickhouse集群必选
	VolumeType *string `json:"volume_type,omitempty"`

	// 数据盘大小，doris集群、clickhouse集群必选。 fe[200,2000] be[400,10000] server[500,10000] zookeeper[200,1000]
	VolumeSize *int32 `json:"volume_size,omitempty"`

	// 节点入参类型 0：flavor模式 ，1：cu模式，doris、hbase、clickhouse都是flavor模式
	FlavorType *string `json:"flavor_type,omitempty"`
}

func (o CreateClusterInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstanceBody struct{}"
	}

	return strings.Join([]string{"CreateClusterInstanceBody", string(data)}, " ")
}
