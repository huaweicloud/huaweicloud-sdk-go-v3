package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSparkJobRequestBody 创建批处理作业请求body体。
type CreateSparkJobRequestBody struct {

	// 参数解释:   用户已上传到DLI资源管理系统的类型为jar的资源包名 示例: batchTest/spark-examples_2.11-2.1.0.luxor.jar 约束限制:  无 取值范围: 无 默认取值: 无
	File string `json:"file"`

	// 参数解释:   批处理作业的Java/Spark主类 示例: org.apache.spark.examples.SparkPi 约束限制:  无 取值范围: 无 默认取值: 无
	ClassName string `json:"className"`

	// 参数解释:   用于指定队列，填写已创建DLI的队列名 示例: test 约束限制:  无 取值范围: 无 默认取值: 无
	ClusterName *string `json:"cluster_name,omitempty"`

	// 参数解释:   传入主类的参数 示例: -o result.txt 约束限制:  无 取值范围: 无 默认取值: 无
	Args *[]string `json:"args,omitempty"`

	// 参数解释:   计算资源类型，目前可接受参数A, B, C。如果不指定，则按最小类型创建。 示例: A 约束限制:  无 取值范围: A：物理资源：8核32G内存，driverCores：2；executorCores：1；driverMemory：7G；executorMemory：4G；numExecutor：6。 B：16核64G内存,2,2,7G,8G,7。 C：32核128G内存,4,2,15G,8G,14。 默认取值: 无
	ScType *string `json:"sc_type,omitempty"`

	// 参数解释:   用户已上传到DLI资源管理系统的类型为jar的程序包名。也支持指定OBS路径，例如：obs://桶名/包名 示例: obs://cwk/jars/com/example/demo 约束限制:  无 取值范围: 无 默认取值: 无
	Jars *[]string `json:"jars,omitempty"`

	// 参数解释:   用户已上传到DLI资源管理系统的类型为pyFile的资源包名。也支持指定OBS路径，例如：obs://桶名/包名 示例: obs://cwk/py/com/example/demo 约束限制:  无 取值范围: 无 默认取值: 无
	PyFiles *[]string `json:"pyFiles,omitempty"`

	// 参数解释:   用户已上传到DLI资源管理系统的类型为file的资源包名 示例: [\"count.txt\"] 约束限制:  无 取值范围: 无 默认取值: 无
	Files *[]string `json:"files,omitempty"`

	// 参数解释:   依赖的系统资源模块名，具体模块名可通过查询所有资源包接口查看。 DLI系统提供了用于执行跨源作业的依赖模块，各个不同的服务对应的模块列表如下： CloudTable/MRS HBase: sys.datasource.hbase CloudTable/MRS OpenTSDB: sys.datasource.opentsdb RDS MySQL: sys.datasource.rds RDS PostGre: 不需要选 DWS: 不需要选 CSS: sys.datasource.css 示例: [\"sys.datasource.hbase\",\"sys.datasource.rds\"] 约束限制:  无 取值范围: 无 默认取值: 无
	Modules *[]string `json:"modules,omitempty"`

	// JSON对象列表，填写用户已上传到队列的类型为JSON的资源包名和类型。
	Resources *[]JobResource `json:"resources,omitempty"`

	// JSON对象列表，填写用户组类型资源，格式详见请求示例。resources中的name未进行type校验，只要此分组中存在这个名字的包即可。
	Groups *[]JobResourcesGroup `json:"groups,omitempty"`

	// 参数解释:   batch配置项 约束限制:  无 取值范围: 无 默认取值: 无
	Conf map[string]interface{} `json:"conf,omitempty"`

	// 参数解释:   创建时用户指定的批处理名称 示例: TestDemo4 约束限制:  不超过128个字符的字符串 取值范围: 无 默认取值: 无
	Name *string `json:"name,omitempty"`

	// 参数解释:   Spark应用的Driver内存, 参数配置例如2G,2048M。该配置项会替换“sc_type”中对应的默认参数，使用时必需带单位，否则会启动失败 示例: 2G 约束限制:  无 取值范围: 无 默认取值: 无
	DriverMemory *string `json:"driverMemory,omitempty"`

	// 参数解释:   Spark应用Driver的CPU核数。该配置项会替换sc_type中对应的默认参数 示例: 8 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	DriverCores *int32 `json:"driverCores,omitempty"`

	// 参数解释:   Spark应用的Executor内存, 参数配置例如2G,2048M。该配置项会替换“sc_type”中对应的默认参数，使用时必需带单位，否则会启动失败。 示例: 2G 约束限制:  无 取值范围: 无 默认取值: 无
	ExecutorMemory *string `json:"executorMemory,omitempty"`

	// 参数解释:   Spark应用每个Executor的CPU核数。该配置项会替换sc_type中对应的默认参数 示例: 8 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	ExecutorCores *int32 `json:"executorCores,omitempty"`

	// 参数解释:   Spark应用Executor的个数。该配置项会替换sc_type中对应的默认参数 示例: 6 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	NumExecutors *int32 `json:"numExecutors,omitempty"`

	// 参数解释:   作业特性，作业运行在vm队列上支持basic，在container队列上支持basic、ai、custom，其中填写custom时需要同时填写image参数 示例: basic 约束限制:  无 取值范围: basic（基础型） ai（AI增强型） custom（自定义型） 默认取值: 无
	Feature *CreateSparkJobRequestBodyFeature `json:"feature,omitempty"`

	// 参数解释:   作业使用spark组件的版本号，在feature为“basic”或“ai”时填写，若不填写，则使用默认的spark组件版本号2.3.2 示例: 2.3.2 约束限制:  无 取值范围: 无 默认取值: 无
	SparkVersion *string `json:"spark_version,omitempty"`

	// 参数解释:   用于指定队列，填写已创建DLI的队列名 示例: gen_native 约束限制:  无 取值范围: 无 默认取值: 无
	Queue *string `json:"queue,omitempty"`

	// 参数解释:   是否开启重试功能，如果开启，可在Spark作业异常失败后自动重试 示例: false 约束限制:  无 取值范围: true,false 默认取值: false
	AutoRecovery *bool `json:"auto_recovery,omitempty"`

	// 参数解释:   最大重试次数 示例: 100 约束限制:  无 取值范围: 大于等于0的整数 默认取值: false
	MaxRetryTimes *int32 `json:"max_retry_times,omitempty"`

	// 参数解释:   自定义镜像。格式为：组织名/镜像名:镜像版本。当用户设置“feature”为“custom”时，该参数生效。用户可通过与“feature”参数配合使用，指定作业运行使用自定义的Spark镜像。关于如何使用自定义镜像，请参考《数据湖探索用户指南》 示例: ceshi/spark_general-x86_64:3.3.1-2.3.7.1720240718867424736954752.tensorflow  约束限制:  无 取值范围: 无 默认取值: 无
	Image *string `json:"image,omitempty"`

	// 参数解释:   授权给DLI的委托名。Spark3.3.1版本时支持配置该参数。 示例: agency 约束限制:  无 取值范围: 无 默认取值: 无
	ExecutionAgencyUrn *string `json:"execution_agency_urn,omitempty"`

	// 参数解释:   保存Spark作业的obs桶，需要保存作业时配置该参数 示例: rain3 约束限制:  无 取值范围: 无 默认取值: 无
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// 访问元数据时，需要将该参数配置为dli。
	CatalogName *string `json:"catalog_name,omitempty"`
}

func (o CreateSparkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSparkJobRequestBody", string(data)}, " ")
}

type CreateSparkJobRequestBodyFeature struct {
	value string
}

type CreateSparkJobRequestBodyFeatureEnum struct {
	BASIC  CreateSparkJobRequestBodyFeature
	AI     CreateSparkJobRequestBodyFeature
	CUSTOM CreateSparkJobRequestBodyFeature
}

func GetCreateSparkJobRequestBodyFeatureEnum() CreateSparkJobRequestBodyFeatureEnum {
	return CreateSparkJobRequestBodyFeatureEnum{
		BASIC: CreateSparkJobRequestBodyFeature{
			value: "basic",
		},
		AI: CreateSparkJobRequestBodyFeature{
			value: "ai",
		},
		CUSTOM: CreateSparkJobRequestBodyFeature{
			value: "custom",
		},
	}
}

func (c CreateSparkJobRequestBodyFeature) Value() string {
	return c.value
}

func (c CreateSparkJobRequestBodyFeature) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSparkJobRequestBodyFeature) UnmarshalJSON(b []byte) error {
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
