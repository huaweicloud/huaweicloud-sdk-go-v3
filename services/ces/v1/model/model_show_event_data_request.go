package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEventDataRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 指标命名空间，如：弹性云服务器的命名空间为SYS.ECS，文档数据库的命名空间为SYS.DDS，各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Namespace string `json:"namespace"`
	// 指标的第一层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.0=key,value，如dim.0=mongodb_cluster_id,4270ff17-aba3-4138-89fa-820594c39755；key为指标的维度信息，如：文档数据库服务，则第一层维度为mongodb_cluster_id，value为文档数据库实例ID；各服务资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Dim0 string `json:"dim.0"`
	// 指标的第二层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.1=key,value，如dim.1=mongos_instance_id,c65d39d7-185c-4616-9aca-ad65703b15f9；key为指标的维度信息，如：文档数据库服务，则第二层维度为mongos_instance_id，value为文档数据库集群实例下的mongos节点ID；各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Dim1 *string `json:"dim.1,omitempty"`
	// 指标的第三层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.2=key,value，如dim.2=mongod_primary_instance_id,5f9498e9-36f8-4317-9ea1-ebe28cba99b4；key为指标的维度信息，如：文档数据库服务，则第三层维度为mongod_primary_instance_id，value为文档数据库实例下的主节点ID；各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Dim2 *string `json:"dim.2,omitempty"`
	// 指标的第四层维度，目前最大支持4个维度，维度编号从0开始；维度格式为dim.3=key,value，如dim.3=mongod_secondary_instance_id,b46fa2c7-aac6-4ae3-9337-f4ea97f885cb；key为指标的维度信息，如：文档数据库服务，则第四层维度为mongod_secondary_instance_id，value为文档数据库实例下的备节点ID；各资源的指标维度名称可查看：“[服务指标维度](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。

	Dim3 *string `json:"dim.3,omitempty"`
	// 事件类型，只允许字母、下划线、中划线，字母开头，长度不超过64，如instance_host_info。

	Type string `json:"type"`
	// 查询数据起始时间，UNIX时间戳，单位毫秒；如：1607146998177。

	From int64 `json:"from"`
	// 查询数据截止时间UNIX时间戳，单位毫秒。from必须小于to；如：1607150598177。

	To int64 `json:"to"`
}

func (o ShowEventDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDataRequest struct{}"
	}

	return strings.Join([]string{"ShowEventDataRequest", string(data)}, " ")
}
