package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Links struct {
	LinkConfigValues *LinksLinkconfigvalues `json:"link-config-values" xml:"link-config-values"`

	// 创建连接的用户
	CreationUser *string `json:"creation-user,omitempty" xml:"creation-user"`

	// 连接名称
	Name string `json:"name" xml:"name"`

	// 连接ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 创建连接的时间
	CreationDate *int64 `json:"creation-date,omitempty" xml:"creation-date"`

	// 连接器名称，对应的连接参数如下：generic-jdbc-connector：关系数据库连接。obs-connector：OBS连接、阿里云OSS连接。hdfs-connector：HDFS连接。hbase-connector：HBase连接、CloudTable连接。hive-connector：Hive连接。ftp-connector/sftp-connector：FTP/SFTP连接。mongodb-connector：MongoDB连接。redis-connector：Redis/DCS连接。nas-connector：NAS/SFS连接。kafka-connector：Kafka连接。dis-connector：DIS连接。elasticsearch-connector：Elasticsearch/云搜索服务连接。dli-connector：DLI连接。opentsdb-connector：CloudTable OpenTSDB连接。http-connector：HTTP/HTTPS连接，该连接暂无连接参数。thirdparty-obs-connector：七牛云KODO/腾讯云COS连接、亚马逊对象存储连接。dms-kafka-connector：DMS Kafka连接
	ConnectorName string `json:"connector-name" xml:"connector-name"`

	// 更新连接的时间
	UpdateDate *int64 `json:"update-date,omitempty" xml:"update-date"`

	// 是否激活连接，默认为“true”
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// 更新连接的用户
	UpdateUser *string `json:"update-user,omitempty" xml:"update-user"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
