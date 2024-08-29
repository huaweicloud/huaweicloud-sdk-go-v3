package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// TriggerEventDataRequestBody 触发器源事件。
type TriggerEventDataRequestBody struct {

	// - TIMER触发器：触发器名称 - APIG触发器：API名称 - CTS触发器：通知名称 - OBS触发器：事件通知名称，默认值为触发器id
	Name *string `json:"name,omitempty"`

	// 定时触发类型（TIMER触发器参数）。TIMER触发器此参数必填 - Rate：指定固定频率（分钟、小时、天数）定期调用函数，单位为分钟时，输入值不能超过60；单位为小时时，输入值不能超过24；单位为天时，输入值不能超过30。 - Cron：指定Cron表达式定期调用函数
	ScheduleType *TriggerEventDataRequestBodyScheduleType `json:"schedule_type,omitempty"`

	// 定时触发规则（TIMER触发器参数）。TIMER触发器此参数必填。 - 触发类型为Rate时对应定时规则 - 触发类型为Cron时对应Cron表达式
	Schedule *string `json:"schedule,omitempty"`

	// 附加信息（TIMER触发器参数）。 当Timer触发器触发函数执行时，执行事件（函数的event参数）为： {\"version\": \"v1.0\",   \"time\": \"2018-06-01T08:30:00+08:00\",   \"trigger_type\": \"TIMER\",   \"trigger_name\": \"Timer_001\",   \"user_event\": \"您输入的附加信息\"}
	UserEvent *string `json:"user_event,omitempty"`

	// API接口类型（APIG触发器参数）。APIG触发器此参数必填。 - 1：公有API - 2：私有API
	Type *int32 `json:"type,omitempty"`

	// APIG接口PATH路径（APIG触发器参数）。APIG触发器此参数必填。
	Path *string `json:"path,omitempty"`

	// API的请求协议（APIG触发器参数）。APIG触发器此参数必填。
	Protocol *TriggerEventDataRequestBodyProtocol `json:"protocol,omitempty"`

	// API的请求方式（APIG触发器参数）。APIG触发器此参数必填。
	ReqMethod *TriggerEventDataRequestBodyReqMethod `json:"req_method,omitempty"`

	// API所属的分组编号（APIG触发器参数）。APIG触发器此参数必填。
	GroupId *string `json:"group_id,omitempty"`

	// API所属的分组名称
	GroupName *string `json:"group_name,omitempty"`

	// API的匹配方式（APIG触发器参数）。APIG触发器此参数必填。 - SWA：前缀匹配 - NORMAL：正常匹配（绝对匹配）
	MatchMode *TriggerEventDataRequestBodyMatchMode `json:"match_mode,omitempty"`

	// API的发布环境（APIG触发器参数）。APIG触发器此参数必填。
	EnvName *string `json:"env_name,omitempty"`

	// API的发布环境id（APIG触发器参数）。APIG触发器此参数必填。
	EnvId *string `json:"env_id,omitempty"`

	// API的认证方式（APIG触发器参数）。APIG触发器此参数必填。 - IAM：IAM认证，只允许IAM用户能访问，安全级别中等 - APP：采用Appkey&Appsecret认证，安全级别高，推荐使用 - NONE：无认证模式，所有用户均可访问，不推荐使用
	Auth *TriggerEventDataRequestBodyAuth `json:"auth,omitempty"`

	FuncInfo *ApigTriggerFuncInfo `json:"func_info,omitempty"`

	// APIG系统默认分配的子域名（APIG触发器参数）。
	SlDomain *string `json:"sl_domain,omitempty"`

	// API的后端类型（APIG触发器参数）。
	BackendType *TriggerEventDataRequestBodyBackendType `json:"backend_type,omitempty"`

	// 自定义操作（CTS触发器参数）。CTS触发器此参数必填。 CTS云审计服务类型和操作订阅所需要的事件通知，当CTS云审计服务获取已订阅的操作记录后，通过CTS触发器将采集到的操作记录作为参数传递来调用FunctionGraph函数。
	Operations *[]string `json:"operations,omitempty"`

	// 实例id。DDS、KAFKA、RABBITMQ触发器此参数必填。 - APIG触发器：apig实例id - DDS触发器：文档数据库实例id - KAFKA触发器：KAFKA实例id - RABBITMQ触发器：RABBITMQ实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 集合名称（DDS触发器参数）。DDS触发器此参数必填。
	CollectionName *string `json:"collection_name,omitempty"`

	// 文档数据库名称（DDS触发器参数）。DDS触发器此参数必填。
	DbName *string `json:"db_name,omitempty"`

	// 文档数据库密码（DDS触发器参数）。DDS触发器此参数必填。
	DbPassword *string `json:"db_password,omitempty"`

	// 批处理大小，单次函数执行处理的最大数据量。DIS、DDS、KAFKA、RABBITMQ触发器此参数必填。 - DDS触发器：批处理大小设置1-10,000的范围内 - DIS触发器：批处理大小设置1-10,000的范围内 - KAFKA触发器：批处理大小设置1-1,000的范围内 - RABBITMQ触发器：批处理大小设置1-1,000的范围内
	BatchSize *int32 `json:"batch_size,omitempty"`

	// 队列id（DMS触发器参数）。DMS触发器此参数必填。
	QueueId *string `json:"queue_id,omitempty"`

	// 消费组id（DMS触发器参数）。DMS触发器此参数必填。
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`

	// 拉取周期。DIS、DMS触发器此参数必填。
	PollingInterval *int32 `json:"polling_interval,omitempty"`

	// 通道名称（DIS触发器参数）。DIS触发器此参数必填。
	StreamName *string `json:"stream_name,omitempty"`

	// 起始位置（DIS触发器参数）。DIS触发器此参数必填。 - TRIM_HORIZON：从最早被存储至分区的有效记录开始读取。 - LATEST：从分区中的最新记录开始读取，此设置可以保证总是读到分区中最新记录。
	SharditeratorType *TriggerEventDataRequestBodySharditeratorType `json:"sharditerator_type,omitempty"`

	// 拉取周期单位（DIS触发器参数）。DIS触发器此参数必填。 - s：秒 - ms：毫秒
	PollingUnit *TriggerEventDataRequestBodyPollingUnit `json:"polling_unit,omitempty"`

	// 最大提取字节数（DIS触发器参数）。
	MaxFetchBytes *int32 `json:"max_fetch_bytes,omitempty"`

	// 串行处理数据（DIS触发器参数），如果开启该选项，取一次数据处理完之后才会取下一次数据；否则只要拉取周期到了就会取数据进行处理。DIS触发器此参数必填。
	IsSerial *TriggerEventDataRequestBodyIsSerial `json:"is_serial,omitempty"`

	// 日志组id（LTS触发器参数）。LTS触发器此参数必填。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流id（LTS触发器参数）。LTS触发器此参数必填。
	LogTopicId *string `json:"log_topic_id,omitempty"`

	// 桶名称（OBS触发器参数），用作事件源的OBS存储桶，不能和本用户已有桶重名；不能和其他用户已有的桶重名；创建成功后不支持修改。OBS触发器此参数必填。
	Bucket *string `json:"bucket,omitempty"`

	// 前缀（OBS触发器参数），输入一个可选性前缀来限制对以此关键字开头的对象的通知。
	Prefix *string `json:"prefix,omitempty"`

	// 后缀（OBS触发器参数），输入一个可选性后缀来限制对以此关键字结尾的对象的通知
	Suffix *string `json:"suffix,omitempty"`

	// 触发事件（OBS触发器参数）。OBS触发器此参数必填。 - ObjectCreated：表示所有创建对象的操作，包含Put、Post、Copy对象以及合并段 - Put：使用Put方法上传对象 - Post：使用Post方法上传对象 - Copy：使用copy方法复制对象 - CompleteMultipartUpload：表示合并分段任务 - ObjectRemoved：表示删除对象 - Delete：指定对象版本号删除对象 - DeleteMarkerCreated：不指定对象版本号删除对象
	Events *[]string `json:"events,omitempty"`

	// 主题URN（SMN触发器参数）。SMN触发器此参数必填。
	TopicUrn *string `json:"topic_urn,omitempty"`

	// KAFKA主题id列表（KAFKA触发器参数）。KAFKA触发器此参数必填。
	TopicIds *[]string `json:"topic_ids,omitempty"`

	// KAFKA账户名（KAFKA触发器参数）。
	KafkaUser *string `json:"kafka_user,omitempty"`

	// KAFKA账户密码（KAFKA触发器参数）。
	KafkaPassword *string `json:"kafka_password,omitempty"`

	// KAFKA实例连接IP地址（KAFKA触发器参数）。
	KafkaConnectAddress *string `json:"kafka_connect_address,omitempty"`

	// KAFKA连接是否开启安全认证（KAFKA触发器参数）。
	KafkaSslEnable *bool `json:"kafka_ssl_enable,omitempty"`

	// RABBITMQ账户密码（RABBITMQ触发器参数）。RABBITMQ触发器此参数必填。
	AccessPassword *string `json:"access_password,omitempty"`

	// RABBITMQ账户名（RABBITMQ触发器参数）。
	AccessUser *string `json:"access_user,omitempty"`

	// 实例连接IP地址（RABBITMQ触发器参数）。
	ConnectAddress *string `json:"connect_address,omitempty"`

	// 交换机名称（RABBITMQ触发器参数）。RABBITMQ触发器此参数必填。
	ExchangeName *string `json:"exchange_name,omitempty"`

	// 虚拟机名称（RABBITMQ触发器参数）。
	Vhost *string `json:"vhost,omitempty"`

	// RABBITMQ连接是否开启安全认证（RABBITMQ触发器参数）。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// EG obs触发器是否对对象加密。
	KeyEncode *bool `json:"Key_encode,omitempty"`

	// 使用的代理
	Agency *string `json:"agency,omitempty"`

	// 通道名称
	ChannelName *string `json:"channel_name,omitempty"`

	// 事件源名称
	SourceName *string `json:"source_name,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 触发器状态
	Status *string `json:"status,omitempty"`

	// 触发器名称
	TriggerName *string `json:"trigger_name,omitempty"`
}

func (o TriggerEventDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerEventDataRequestBody struct{}"
	}

	return strings.Join([]string{"TriggerEventDataRequestBody", string(data)}, " ")
}

type TriggerEventDataRequestBodyScheduleType struct {
	value string
}

type TriggerEventDataRequestBodyScheduleTypeEnum struct {
	RATE TriggerEventDataRequestBodyScheduleType
	CRON TriggerEventDataRequestBodyScheduleType
}

func GetTriggerEventDataRequestBodyScheduleTypeEnum() TriggerEventDataRequestBodyScheduleTypeEnum {
	return TriggerEventDataRequestBodyScheduleTypeEnum{
		RATE: TriggerEventDataRequestBodyScheduleType{
			value: "Rate",
		},
		CRON: TriggerEventDataRequestBodyScheduleType{
			value: "Cron",
		},
	}
}

func (c TriggerEventDataRequestBodyScheduleType) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyScheduleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyScheduleType) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyProtocol struct {
	value string
}

type TriggerEventDataRequestBodyProtocolEnum struct {
	HTTP  TriggerEventDataRequestBodyProtocol
	HTTPS TriggerEventDataRequestBodyProtocol
}

func GetTriggerEventDataRequestBodyProtocolEnum() TriggerEventDataRequestBodyProtocolEnum {
	return TriggerEventDataRequestBodyProtocolEnum{
		HTTP: TriggerEventDataRequestBodyProtocol{
			value: "HTTP",
		},
		HTTPS: TriggerEventDataRequestBodyProtocol{
			value: "HTTPS",
		},
	}
}

func (c TriggerEventDataRequestBodyProtocol) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyProtocol) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyReqMethod struct {
	value string
}

type TriggerEventDataRequestBodyReqMethodEnum struct {
	GET     TriggerEventDataRequestBodyReqMethod
	POST    TriggerEventDataRequestBodyReqMethod
	PUT     TriggerEventDataRequestBodyReqMethod
	DELETE  TriggerEventDataRequestBodyReqMethod
	HEAD    TriggerEventDataRequestBodyReqMethod
	PATCH   TriggerEventDataRequestBodyReqMethod
	OPTIONS TriggerEventDataRequestBodyReqMethod
	ANY     TriggerEventDataRequestBodyReqMethod
}

func GetTriggerEventDataRequestBodyReqMethodEnum() TriggerEventDataRequestBodyReqMethodEnum {
	return TriggerEventDataRequestBodyReqMethodEnum{
		GET: TriggerEventDataRequestBodyReqMethod{
			value: "GET",
		},
		POST: TriggerEventDataRequestBodyReqMethod{
			value: "POST",
		},
		PUT: TriggerEventDataRequestBodyReqMethod{
			value: "PUT",
		},
		DELETE: TriggerEventDataRequestBodyReqMethod{
			value: "DELETE",
		},
		HEAD: TriggerEventDataRequestBodyReqMethod{
			value: "HEAD",
		},
		PATCH: TriggerEventDataRequestBodyReqMethod{
			value: "PATCH",
		},
		OPTIONS: TriggerEventDataRequestBodyReqMethod{
			value: "OPTIONS",
		},
		ANY: TriggerEventDataRequestBodyReqMethod{
			value: "ANY",
		},
	}
}

func (c TriggerEventDataRequestBodyReqMethod) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyReqMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyReqMethod) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyMatchMode struct {
	value string
}

type TriggerEventDataRequestBodyMatchModeEnum struct {
	SWA    TriggerEventDataRequestBodyMatchMode
	NORMAL TriggerEventDataRequestBodyMatchMode
}

func GetTriggerEventDataRequestBodyMatchModeEnum() TriggerEventDataRequestBodyMatchModeEnum {
	return TriggerEventDataRequestBodyMatchModeEnum{
		SWA: TriggerEventDataRequestBodyMatchMode{
			value: "SWA",
		},
		NORMAL: TriggerEventDataRequestBodyMatchMode{
			value: "NORMAL",
		},
	}
}

func (c TriggerEventDataRequestBodyMatchMode) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyMatchMode) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyAuth struct {
	value string
}

type TriggerEventDataRequestBodyAuthEnum struct {
	IAM  TriggerEventDataRequestBodyAuth
	APP  TriggerEventDataRequestBodyAuth
	NONE TriggerEventDataRequestBodyAuth
}

func GetTriggerEventDataRequestBodyAuthEnum() TriggerEventDataRequestBodyAuthEnum {
	return TriggerEventDataRequestBodyAuthEnum{
		IAM: TriggerEventDataRequestBodyAuth{
			value: "IAM",
		},
		APP: TriggerEventDataRequestBodyAuth{
			value: "APP",
		},
		NONE: TriggerEventDataRequestBodyAuth{
			value: "NONE",
		},
	}
}

func (c TriggerEventDataRequestBodyAuth) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyAuth) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyAuth) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyBackendType struct {
	value string
}

type TriggerEventDataRequestBodyBackendTypeEnum struct {
	FUNCTION TriggerEventDataRequestBodyBackendType
}

func GetTriggerEventDataRequestBodyBackendTypeEnum() TriggerEventDataRequestBodyBackendTypeEnum {
	return TriggerEventDataRequestBodyBackendTypeEnum{
		FUNCTION: TriggerEventDataRequestBodyBackendType{
			value: "FUNCTION",
		},
	}
}

func (c TriggerEventDataRequestBodyBackendType) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyBackendType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyBackendType) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodySharditeratorType struct {
	value string
}

type TriggerEventDataRequestBodySharditeratorTypeEnum struct {
	TRIM_HORIZON TriggerEventDataRequestBodySharditeratorType
	LATEST       TriggerEventDataRequestBodySharditeratorType
}

func GetTriggerEventDataRequestBodySharditeratorTypeEnum() TriggerEventDataRequestBodySharditeratorTypeEnum {
	return TriggerEventDataRequestBodySharditeratorTypeEnum{
		TRIM_HORIZON: TriggerEventDataRequestBodySharditeratorType{
			value: "TRIM_HORIZON",
		},
		LATEST: TriggerEventDataRequestBodySharditeratorType{
			value: "LATEST",
		},
	}
}

func (c TriggerEventDataRequestBodySharditeratorType) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodySharditeratorType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodySharditeratorType) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyPollingUnit struct {
	value string
}

type TriggerEventDataRequestBodyPollingUnitEnum struct {
	S  TriggerEventDataRequestBodyPollingUnit
	MS TriggerEventDataRequestBodyPollingUnit
}

func GetTriggerEventDataRequestBodyPollingUnitEnum() TriggerEventDataRequestBodyPollingUnitEnum {
	return TriggerEventDataRequestBodyPollingUnitEnum{
		S: TriggerEventDataRequestBodyPollingUnit{
			value: "s",
		},
		MS: TriggerEventDataRequestBodyPollingUnit{
			value: "ms",
		},
	}
}

func (c TriggerEventDataRequestBodyPollingUnit) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyPollingUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyPollingUnit) UnmarshalJSON(b []byte) error {
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

type TriggerEventDataRequestBodyIsSerial struct {
	value string
}

type TriggerEventDataRequestBodyIsSerialEnum struct {
	TRUE  TriggerEventDataRequestBodyIsSerial
	FALSE TriggerEventDataRequestBodyIsSerial
}

func GetTriggerEventDataRequestBodyIsSerialEnum() TriggerEventDataRequestBodyIsSerialEnum {
	return TriggerEventDataRequestBodyIsSerialEnum{
		TRUE: TriggerEventDataRequestBodyIsSerial{
			value: "true",
		},
		FALSE: TriggerEventDataRequestBodyIsSerial{
			value: "false",
		},
	}
}

func (c TriggerEventDataRequestBodyIsSerial) Value() string {
	return c.value
}

func (c TriggerEventDataRequestBodyIsSerial) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataRequestBodyIsSerial) UnmarshalJSON(b []byte) error {
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
