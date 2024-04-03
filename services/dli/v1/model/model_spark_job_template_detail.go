package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SparkJobTemplateDetail 创建批处理作业请求body体。
type SparkJobTemplateDetail struct {

	// 用户已上传到DLI资源管理系统的类型为jar的资源包名。
	File string `json:"file"`

	// 批处理作业的Java/Spark主类。
	ClassName string `json:"className"`

	// 用于指定队列，填写已创建DLI的队列名。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 传入主类的参数。
	Args *[]string `json:"args,omitempty"`

	// 计算资源类型，目前可接受参数A, B, C。如果不指定，则按最小类型创建。 资源类型： A：物理资源：8核32G内存，driverCores：2；executorCores：1；driverMemory：7G；executorMemory：4G；numExecutor：6。 B：16核64G内存,2,2,7G,8G,7。 C：32核128G内存,4,2,15G,8G,14。
	ScType *string `json:"sc_type,omitempty"`

	// 用户已上传到DLI资源管理系统的类型为jar的资源包名。
	Jars *[]string `json:"jars,omitempty"`

	// 用户已上传到DLI资源管理系统的类型为pyFile的资源包名。
	PyFiles *[]string `json:"pyFiles,omitempty"`

	// 用户已上传到DLI资源管理系统的类型为file的资源包名。
	Files *[]string `json:"files,omitempty"`

	// 依赖的系统资源模块名，具体模块名可通过查询所有资源包接口查看。 DLI系统提供了用于执行跨源作业的依赖模块，各个不同的服务对应的模块列表如下： CloudTable/MRS HBase: sys.datasource.hbase CloudTable/MRS OpenTSDB: sys.datasource.opentsdb RDS MySQL: sys.datasource.rds RDS PostGre: 不需要选 DWS: 不需要选 CSS: sys.datasource.css
	Modules *[]string `json:"modules,omitempty"`

	// JSON对象列表，填写用户已上传到队列的类型为JSON的资源包名和类型。
	Resources *[]JobResource `json:"resources,omitempty"`

	// JSON对象列表，填写用户组类型资源，格式详见请求示例。resources中的name未进行type校验，只要此分组中存在这个名字的包即可。
	Groups *[]JobResourcesGroup `json:"groups,omitempty"`

	// batch配置项。
	Conf map[string]interface{} `json:"conf,omitempty"`

	// 创建时用户指定的批处理名称，不能超过128个字符。
	Name *string `json:"name,omitempty"`

	// Spark应用的Driver内存, 参数配置例如2G, 2048M。该配置项会替换“sc_type”中对应的默认参数，使用时必需带单位，否则会启动失败。
	DriverMemory *string `json:"driverMemory,omitempty"`

	// Spark应用Driver的CPU核数。该配置项会替换sc_type中对应的默认参数。
	DriverCores *int32 `json:"driverCores,omitempty"`

	// Spark应用的Executor内存, 参数配置例如2G, 2048M。该配置项会替换“sc_type”中对应的默认参数，使用时必需带单位，否则会启动失败。
	ExecutorMemory *string `json:"executorMemory,omitempty"`

	// Spark应用每个Executor的CPU核数。该配置项会替换sc_type中对应的默认参数。
	ExecutorCores *int32 `json:"executorCores,omitempty"`

	// Spark应用Executor的个数。该配置项会替换sc_type中对应的默认参数。
	NumExecutors *int32 `json:"numExecutors,omitempty"`

	// 作业特性，作业运行在vm队列上支持basic，在container队列上支持basic、ai、custom，其中填写custom时需要同时填写image参数。
	Feature *SparkJobTemplateDetailFeature `json:"feature,omitempty"`

	// 作业使用spark组件的版本号，在feature为“basic”或“ai”时填写，若不填写，则使用默认的spark组件版本号2.3.2。
	SparkVersion *string `json:"spark_version,omitempty"`

	// 用于指定队列，填写已创建DLI的队列名
	Queue *string `json:"queue,omitempty"`

	// 是否开启重试功能，如果开启，可在Spark作业异常失败后自动重试。默认值为“false”。
	AutoRecovery *bool `json:"auto_recovery,omitempty"`

	// 最大重试次数。最大值为“100”，默认值为“20”。
	MaxRetryTimes *int32 `json:"max_retry_times,omitempty"`

	// 自定义镜像。格式为：组织名/镜像名:镜像版本。当用户设置“feature”为“custom”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用自定义的Spark镜像。关于如何使用自定义镜像，请参考《数据湖探索用户指南》。
	Image *string `json:"image,omitempty"`

	// 保存Spark作业的obs桶，需要保存作业时配置该参数
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 访问元数据时，需要将该参数配置为dli。
	CatalogName *string `json:"catalog_name,omitempty"`
}

func (o SparkJobTemplateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobTemplateDetail struct{}"
	}

	return strings.Join([]string{"SparkJobTemplateDetail", string(data)}, " ")
}

type SparkJobTemplateDetailFeature struct {
	value string
}

type SparkJobTemplateDetailFeatureEnum struct {
	BASIC  SparkJobTemplateDetailFeature
	AI     SparkJobTemplateDetailFeature
	CUSTOM SparkJobTemplateDetailFeature
}

func GetSparkJobTemplateDetailFeatureEnum() SparkJobTemplateDetailFeatureEnum {
	return SparkJobTemplateDetailFeatureEnum{
		BASIC: SparkJobTemplateDetailFeature{
			value: "basic",
		},
		AI: SparkJobTemplateDetailFeature{
			value: "ai",
		},
		CUSTOM: SparkJobTemplateDetailFeature{
			value: "custom",
		},
	}
}

func (c SparkJobTemplateDetailFeature) Value() string {
	return c.value
}

func (c SparkJobTemplateDetailFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SparkJobTemplateDetailFeature) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
