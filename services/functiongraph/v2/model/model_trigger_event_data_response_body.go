package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// TriggerEventDataResponseBody 触发器源事件。
type TriggerEventDataResponseBody struct {

	// 触发器名称
	Name *string `json:"name,omitempty"`

	// 定时触发类型（TIMER触发器参数）。 - Rate：指定固定频率（分钟、小时、天数）定期调用函数，单位为分钟时，输入值不能超过60；单位为小时时，输入值不能超过24；单位为天时，输入值不能超过30。 - Cron：指定Cron表达式定期调用函数
	ScheduleType *TriggerEventDataResponseBodyScheduleType `json:"schedule_type,omitempty"`

	// 定时触发规则（TIMER触发器参数）。 - 触发类型为Rate时对应定时规则 - 触发类型为Cron时对应Cron表达式
	Schedule *string `json:"schedule,omitempty"`

	// 附加信息（TIMER触发器参数）。 当Timer触发器触发函数执行时，执行事件（函数的event参数）为： {\"version\": \"v1.0\",   \"time\": \"2018-06-01T08:30:00+08:00\",   \"trigger_type\": \"TIMER\",   \"trigger_name\": \"Timer_001\",   \"user_event\": \"您输入的附加信息\"}
	UserEvent *string `json:"user_event,omitempty"`

	// APIG触发器id。（APIG触发器参数）
	Triggerid *string `json:"triggerid,omitempty"`

	// API接口类型（APIG触发器参数）。 - 1：公有API - 2：私有API
	Type *int32 `json:"type,omitempty"`

	// APIG接口PATH路径（APIG触发器参数）。
	Path *string `json:"path,omitempty"`

	// API的请求协议（APIG触发器参数）。
	Protocol *TriggerEventDataResponseBodyProtocol `json:"protocol,omitempty"`

	// API的请求方式（APIG触发器参数）。
	ReqMethod *TriggerEventDataResponseBodyReqMethod `json:"req_method,omitempty"`

	// API所属的分组编号（APIG触发器参数）。
	GroupId *string `json:"group_id,omitempty"`

	// API所属的分组名称（APIG触发器参数）。
	GroupName *string `json:"group_name,omitempty"`

	// API的匹配方式（APIG触发器参数）。 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配）
	MatchMode *TriggerEventDataResponseBodyMatchMode `json:"match_mode,omitempty"`

	// API的发布环境（APIG触发器参数）。
	EnvName *string `json:"env_name,omitempty"`

	// API的发布环境id（APIG触发器参数）。
	EnvId *string `json:"env_id,omitempty"`

	// API编号（APIG触发器参数）。
	ApiId *string `json:"api_id,omitempty"`

	// API名称（APIG触发器参数）。
	ApiName *string `json:"api_name,omitempty"`

	// API的认证方式（APIG触发器参数）。 - IAM：IAM认证，只允许IAM用户能访问，安全级别中等 - APP：采用Appkey&Appsecret认证，安全级别高，推荐使用 - NONE：无认证模式，所有用户均可访问，不推荐使用
	Auth *TriggerEventDataResponseBodyAuth `json:"auth,omitempty"`

	// API调用地址（APIG触发器参数）。
	InvokeUrl *string `json:"invoke_url,omitempty"`

	FuncInfo *ApigTriggerFuncInfo `json:"func_info,omitempty"`

	// APIG系统默认分配的子域名（APIG触发器参数）。
	SlDomain *string `json:"sl_domain,omitempty"`

	// API的后端类型（APIG触发器参数）。
	BackendType *TriggerEventDataResponseBodyBackendType `json:"backend_type,omitempty"`

	// 实例id。DDS、KAFKA、RABBITMQ触发器此参数必填。 - APIG触发器：apig实例id - DDS触发器：文档数据库实例id - KAFKA触发器：KAFKA实例id - RABBITMQ触发器：RABBITMQ实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// API归属的集成应用编号。（APIG触发器参数）
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 自定义操作（CTS触发器参数）。 CTS云审计服务类型和操作订阅所需要的事件通知，当CTS云审计服务获取已订阅的操作记录后，通过CTS触发器将采集到的操作记录作为参数传递来调用FunctionGraph函数。
	Operations *[]string `json:"operations,omitempty"`

	// 集合名称（DDS触发器参数）。
	CollectionName *string `json:"collection_name,omitempty"`

	// 文档数据库名称（DDS触发器参数）。
	DbName *string `json:"db_name,omitempty"`

	// 文档数据库密码（DDS触发器参数）。
	DbPassword *string `json:"db_password,omitempty"`

	// 文档数据库用户名（DDS触发器参数）。
	DbUser *string `json:"db_user,omitempty"`

	// 文档数据库实例地址（DDS触发器参数）。
	InstanceAddrs *[]string `json:"instance_addrs,omitempty"`

	// 文档数据库实例类型（DDS触发器参数）。 - Sharding：集群实例 - ReplicaSet：副本集实例 - Single：单节点实例
	Mode *string `json:"mode,omitempty"`

	// 批处理大小，单次函数执行处理的最大数据量。DIS、DDS、KAFKA、RABBITMQ触发器此参数必填。 - DDS触发器：批处理大小设置1-10,000的范围内 - DIS触发器：批处理大小设置1-10,000的范围内 - KAFKA触发器：批处理大小设置1-1,000的范围内 - RABBITMQ触发器：批处理大小设置1-1,000的范围内
	BatchSize *int32 `json:"batch_size,omitempty"`

	// 队列id（DMS触发器参数）。
	QueueId *string `json:"queue_id,omitempty"`

	// 消费组id（DMS触发器参数）。
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`

	// 拉取周期。
	PollingInterval *int32 `json:"polling_interval,omitempty"`

	// 通道名称（DIS触发器参数）。
	StreamName *string `json:"stream_name,omitempty"`

	// 起始位置（DIS触发器参数）。 - TRIM_HORIZON：从最早被存储至分区的有效记录开始读取。 - LATEST：从分区中的最新记录开始读取，此设置可以保证总是读到分区中最新记录。
	SharditeratorType *TriggerEventDataResponseBodySharditeratorType `json:"sharditerator_type,omitempty"`

	// 拉取周期单位（DIS触发器参数）。 - s：秒 - ms：毫秒
	PollingUnit *TriggerEventDataResponseBodyPollingUnit `json:"polling_unit,omitempty"`

	// 最大提取字节数（DIS触发器参数）。
	MaxFetchBytes *int32 `json:"max_fetch_bytes,omitempty"`

	// 串行处理数据（DIS触发器参数），如果开启该选项，取一次数据处理完之后才会取下一次数据；否则只要拉取周期到了就会取数据进行处理。
	IsSerial *TriggerEventDataResponseBodyIsSerial `json:"is_serial,omitempty"`

	// 日志组id（LTS触发器参数）。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流id（LTS触发器参数）。
	LogTopicId *string `json:"log_topic_id,omitempty"`

	// 桶名称（OBS触发器参数），用作事件源的OBS存储桶，不能和本用户已有桶重名；不能和其他用户已有的桶重名；创建成功后不支持修改。
	Bucket *string `json:"bucket,omitempty"`

	// 前缀（OBS触发器参数），输入一个可选性前缀来限制对以此关键字开头的对象的通知。
	Prefix *string `json:"prefix,omitempty"`

	// 后缀（OBS触发器参数），输入一个可选性后缀来限制对以此关键字结尾的对象的通知
	Suffix *string `json:"suffix,omitempty"`

	// 触发事件（OBS触发器参数）。 - ObjectCreated：表示所有创建对象的操作，包含Put、Post、Copy对象以及合并段 - Put：使用Put方法上传对象 - Post：使用Post方法上传对象 - Copy：使用copy方法复制对象 - CompleteMultipartUpload：表示合并分段任务 - ObjectRemoved：表示删除对象 - Delete：指定对象版本号删除对象 - DeleteMarkerCreated：不指定对象版本号删除对象
	Events *[]string `json:"events,omitempty"`

	// 主题URN（SMN触发器参数）。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// KAFKA主题id列表（KAFKA触发器参数）。
	TopicIds *[]string `json:"topic_ids,omitempty"`

	// KAFKA账户名（KAFKA触发器参数）。
	KafkaUser *string `json:"kafka_user,omitempty"`

	// KAFKA账户密码（KAFKA触发器参数）。
	KafkaPassword *string `json:"kafka_password,omitempty"`

	// KAFKA实例连接IP地址（KAFKA触发器参数）。
	KafkaConnectAddress *string `json:"kafka_connect_address,omitempty"`

	// KAFKA连接是否开启安全认证（KAFKA触发器参数）。
	KafkaSslEnable *bool `json:"kafka_ssl_enable,omitempty"`

	// RABBITMQ账户密码（RABBITMQ触发器参数）。
	AccessPassword *string `json:"access_password,omitempty"`

	// RABBITMQ账户名（RABBITMQ触发器参数）。
	AccessUser *string `json:"access_user,omitempty"`

	// 实例连接IP地址（RABBITMQ触发器参数）。
	ConnectAddress *string `json:"connect_address,omitempty"`

	// 交换机名称（RABBITMQ触发器参数）。
	ExchangeName *string `json:"exchange_name,omitempty"`

	// 虚拟机名称（RABBITMQ触发器参数）。
	Vhost *string `json:"vhost,omitempty"`

	// RABBITMQ连接是否开启安全认证（RABBITMQ触发器参数）。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// EG obs触发器是否对对象加密（EVENTGRID触发器参数）。
	KeyEncode *bool `json:"Key_encode,omitempty"`

	// 使用的代理（EVENTGRID触发器参数）。
	Agency *string `json:"agency,omitempty"`

	// 通道名称（EVENTGRID触发器参数）。
	ChannelName *string `json:"channel_name,omitempty"`

	// 通道id（EVENTGRID触发器参数）。
	ChannelId *string `json:"channel_id,omitempty"`

	// 事件源名称（EVENTGRID触发器参数）。
	SourceName *string `json:"source_name,omitempty"`

	// 创建时间（EVENTGRID触发器参数）。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 触发器状态（EVENTGRID触发器参数）。
	Status *TriggerEventDataResponseBodyStatus `json:"status,omitempty"`

	// 触发器名称（EVENTGRID触发器参数）。
	TriggerName *string `json:"trigger_name,omitempty"`

	// 事件类型（EVENTGRID触发器参数）。
	EventTypes *[]string `json:"event_types,omitempty"`
}

func (o TriggerEventDataResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerEventDataResponseBody struct{}"
	}

	return strings.Join([]string{"TriggerEventDataResponseBody", string(data)}, " ")
}

type TriggerEventDataResponseBodyScheduleType struct {
	value string
}

type TriggerEventDataResponseBodyScheduleTypeEnum struct {
	RATE TriggerEventDataResponseBodyScheduleType
	CRON TriggerEventDataResponseBodyScheduleType
}

func GetTriggerEventDataResponseBodyScheduleTypeEnum() TriggerEventDataResponseBodyScheduleTypeEnum {
	return TriggerEventDataResponseBodyScheduleTypeEnum{
		RATE: TriggerEventDataResponseBodyScheduleType{
			value: "Rate",
		},
		CRON: TriggerEventDataResponseBodyScheduleType{
			value: "Cron",
		},
	}
}

func (c TriggerEventDataResponseBodyScheduleType) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyScheduleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyScheduleType) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyProtocol struct {
	value string
}

type TriggerEventDataResponseBodyProtocolEnum struct {
	HTTP  TriggerEventDataResponseBodyProtocol
	HTTPS TriggerEventDataResponseBodyProtocol
}

func GetTriggerEventDataResponseBodyProtocolEnum() TriggerEventDataResponseBodyProtocolEnum {
	return TriggerEventDataResponseBodyProtocolEnum{
		HTTP: TriggerEventDataResponseBodyProtocol{
			value: "HTTP",
		},
		HTTPS: TriggerEventDataResponseBodyProtocol{
			value: "HTTPS",
		},
	}
}

func (c TriggerEventDataResponseBodyProtocol) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyProtocol) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyReqMethod struct {
	value string
}

type TriggerEventDataResponseBodyReqMethodEnum struct {
	GET     TriggerEventDataResponseBodyReqMethod
	POST    TriggerEventDataResponseBodyReqMethod
	PUT     TriggerEventDataResponseBodyReqMethod
	DELETE  TriggerEventDataResponseBodyReqMethod
	HEAD    TriggerEventDataResponseBodyReqMethod
	PATCH   TriggerEventDataResponseBodyReqMethod
	OPTIONS TriggerEventDataResponseBodyReqMethod
	ANY     TriggerEventDataResponseBodyReqMethod
}

func GetTriggerEventDataResponseBodyReqMethodEnum() TriggerEventDataResponseBodyReqMethodEnum {
	return TriggerEventDataResponseBodyReqMethodEnum{
		GET: TriggerEventDataResponseBodyReqMethod{
			value: "GET",
		},
		POST: TriggerEventDataResponseBodyReqMethod{
			value: "POST",
		},
		PUT: TriggerEventDataResponseBodyReqMethod{
			value: "PUT",
		},
		DELETE: TriggerEventDataResponseBodyReqMethod{
			value: "DELETE",
		},
		HEAD: TriggerEventDataResponseBodyReqMethod{
			value: "HEAD",
		},
		PATCH: TriggerEventDataResponseBodyReqMethod{
			value: "PATCH",
		},
		OPTIONS: TriggerEventDataResponseBodyReqMethod{
			value: "OPTIONS",
		},
		ANY: TriggerEventDataResponseBodyReqMethod{
			value: "ANY",
		},
	}
}

func (c TriggerEventDataResponseBodyReqMethod) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyReqMethod) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyMatchMode struct {
	value string
}

type TriggerEventDataResponseBodyMatchModeEnum struct {
	SWA    TriggerEventDataResponseBodyMatchMode
	NORMAL TriggerEventDataResponseBodyMatchMode
}

func GetTriggerEventDataResponseBodyMatchModeEnum() TriggerEventDataResponseBodyMatchModeEnum {
	return TriggerEventDataResponseBodyMatchModeEnum{
		SWA: TriggerEventDataResponseBodyMatchMode{
			value: "SWA",
		},
		NORMAL: TriggerEventDataResponseBodyMatchMode{
			value: "NORMAL",
		},
	}
}

func (c TriggerEventDataResponseBodyMatchMode) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyMatchMode) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyAuth struct {
	value string
}

type TriggerEventDataResponseBodyAuthEnum struct {
	IAM  TriggerEventDataResponseBodyAuth
	APP  TriggerEventDataResponseBodyAuth
	NONE TriggerEventDataResponseBodyAuth
}

func GetTriggerEventDataResponseBodyAuthEnum() TriggerEventDataResponseBodyAuthEnum {
	return TriggerEventDataResponseBodyAuthEnum{
		IAM: TriggerEventDataResponseBodyAuth{
			value: "IAM",
		},
		APP: TriggerEventDataResponseBodyAuth{
			value: "APP",
		},
		NONE: TriggerEventDataResponseBodyAuth{
			value: "NONE",
		},
	}
}

func (c TriggerEventDataResponseBodyAuth) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyAuth) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyAuth) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyBackendType struct {
	value string
}

type TriggerEventDataResponseBodyBackendTypeEnum struct {
	FUNCTION TriggerEventDataResponseBodyBackendType
}

func GetTriggerEventDataResponseBodyBackendTypeEnum() TriggerEventDataResponseBodyBackendTypeEnum {
	return TriggerEventDataResponseBodyBackendTypeEnum{
		FUNCTION: TriggerEventDataResponseBodyBackendType{
			value: "FUNCTION",
		},
	}
}

func (c TriggerEventDataResponseBodyBackendType) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyBackendType) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodySharditeratorType struct {
	value string
}

type TriggerEventDataResponseBodySharditeratorTypeEnum struct {
	TRIM_HORIZON TriggerEventDataResponseBodySharditeratorType
	LATEST       TriggerEventDataResponseBodySharditeratorType
}

func GetTriggerEventDataResponseBodySharditeratorTypeEnum() TriggerEventDataResponseBodySharditeratorTypeEnum {
	return TriggerEventDataResponseBodySharditeratorTypeEnum{
		TRIM_HORIZON: TriggerEventDataResponseBodySharditeratorType{
			value: "TRIM_HORIZON",
		},
		LATEST: TriggerEventDataResponseBodySharditeratorType{
			value: "LATEST",
		},
	}
}

func (c TriggerEventDataResponseBodySharditeratorType) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodySharditeratorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodySharditeratorType) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyPollingUnit struct {
	value string
}

type TriggerEventDataResponseBodyPollingUnitEnum struct {
	S  TriggerEventDataResponseBodyPollingUnit
	MS TriggerEventDataResponseBodyPollingUnit
}

func GetTriggerEventDataResponseBodyPollingUnitEnum() TriggerEventDataResponseBodyPollingUnitEnum {
	return TriggerEventDataResponseBodyPollingUnitEnum{
		S: TriggerEventDataResponseBodyPollingUnit{
			value: "s",
		},
		MS: TriggerEventDataResponseBodyPollingUnit{
			value: "ms",
		},
	}
}

func (c TriggerEventDataResponseBodyPollingUnit) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyPollingUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyPollingUnit) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyIsSerial struct {
	value string
}

type TriggerEventDataResponseBodyIsSerialEnum struct {
	TRUE  TriggerEventDataResponseBodyIsSerial
	FALSE TriggerEventDataResponseBodyIsSerial
}

func GetTriggerEventDataResponseBodyIsSerialEnum() TriggerEventDataResponseBodyIsSerialEnum {
	return TriggerEventDataResponseBodyIsSerialEnum{
		TRUE: TriggerEventDataResponseBodyIsSerial{
			value: "true",
		},
		FALSE: TriggerEventDataResponseBodyIsSerial{
			value: "false",
		},
	}
}

func (c TriggerEventDataResponseBodyIsSerial) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyIsSerial) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyIsSerial) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataResponseBodyStatus struct {
	value string
}

type TriggerEventDataResponseBodyStatusEnum struct {
	ACTIVE  TriggerEventDataResponseBodyStatus
	DISABLE TriggerEventDataResponseBodyStatus
}

func GetTriggerEventDataResponseBodyStatusEnum() TriggerEventDataResponseBodyStatusEnum {
	return TriggerEventDataResponseBodyStatusEnum{
		ACTIVE: TriggerEventDataResponseBodyStatus{
			value: "ACTIVE",
		},
		DISABLE: TriggerEventDataResponseBodyStatus{
			value: "DISABLE",
		},
	}
}

func (c TriggerEventDataResponseBodyStatus) Value() string {
	return c.value
}

func (c TriggerEventDataResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataResponseBodyStatus) UnmarshalJSON(b []byte) error {
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
