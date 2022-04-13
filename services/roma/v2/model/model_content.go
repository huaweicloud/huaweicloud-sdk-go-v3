package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type Content struct {
	// gauss100的版本号 - V100R003C20 - V300R001C00

	Gauss100Version *ContentGauss100Version `json:"gauss100_version,omitempty"`
	// 主机IP地址 - 数据源为DWS、HANA、RABBITMQ、SAP、SNMP、IBMMQ类型时需要配置。 - 数据源为MYSQL、ORACLE、SQLSERVER、DB2、GAUSS200、GAUSS100、TAURUS、ARTEMISMQ、POSTGRESQL、HIVE类型且mode为default时需要配置。 - 数据源为HL7类型且作为目标端（position为target）时需要配置。 - 初始值为空，配置任务启动后生成host

	Host *string `json:"host,omitempty"`
	// 端口，端口号为0到65535 - 数据源为DWS、HANA、RABBITMQ、SAP、SNMP、IBMMQ，obs类型时需要配置， - 数据源为MYSQL、ORACLE、SQLSERVER、DB2、GAUSS200、GAUSS100、TAURUS、ARTEMISMQ、POSTGRESQL、HIVE类型且mode为default时需要配置， - 数据源为HL7类型且作为目标端（position为target）时需要配置

	Port *string `json:"port,omitempty"`
	// 数据库名称 - 数据源为DWS、HANA、RABBITMQ、SAP、SNMP、IBMMQ，obs类型时需要配置， - 数据源为MYSQL、SQLSERVER、DB2、GAUSS200、GAUSS100、TAURUS、ARTEMISMQ、POSTGRESQL、HIVE类型且mode为default时需要配置， - 数据源为ORACLE类型且mode为（default、multiAddress）时需要配置

	DatabaseName *string `json:"database_name,omitempty"`
	// REDIS数据源类型配置，数据库编号, 纯数字编码

	RedisDatabase *string `json:"redis_database,omitempty"`
	// 访问服务的用户名 - 数据源为MYSQL、DWS、FTP、ORACLE、MONGODB、HANA、SQLSERVER、DB2、GAUSS200、GAUSS100、TAURUS、ACTIVEMQ、ARTEMISMQ、POSTGRESQL、RABBITMQ、SAP、IBMMQ、HIVE类型时需要配置 - 数据源为WEBSOCKET类型，认证方式（basicauth）时需要配置 - 数据源为LDAP，安全认证类型（security_auth_type）为simple时需要配置

	UserName *string `json:"user_name,omitempty"`
	// 访问服务的密码 - 数据源为MYSQL、DWS、FTP、ORACLE、MONGODB、HANA、SQLSERVER、DB2、GAUSS200、GAUSS100、TAURUS、ACTIVEMQ、ARTEMISMQ、POSTGRESQL、RABBITMQ、SAP、IBMMQ、HIVE类型时需要配置 - 数据源为WEBSOCKET，且认证方式（basicauth）时需要配置 - 数据源为LDAP，且安全认证类型（security_auth_type）为simple时需要配置

	Password *string `json:"password,omitempty"`
	// 数据源连接模式 - 数据源为DWS、MONGODB、REDIS、HANA时配置默认， - 数据源为MYSQL、SQLSERVER、DB2、GAUSS200、GAUSS100、TAURUS、POSTGRESQL、HIVE时配置（default,professional）， - 数据源为ORACLE时配置专有的模式multiAddress - default (默认模式) - professional (专业模式) - multiAddress (多地址)

	Mode *ContentMode `json:"mode,omitempty"`
	// cdc模式，只有组合任务使用

	CdcMode *ContentCdcMode `json:"cdc_mode,omitempty"`
	// ORACLE集群地址，当mode为multiAddress时需要配置

	MultiOracleAddress *[]MultiOracleAddress `json:"multi_oracle_address,omitempty"`
	// ORACLE集群服务名

	OracleServiceName *string `json:"oracle_service_name,omitempty"`
	// 访问FTP服务的连接模式 - active (主动模式) - passive (被动模式)

	FtpConnectMode *ContentFtpConnectMode `json:"ftp_connect_mode,omitempty"`
	// 访问FTP服务协议类型 - sftp - ftp

	FtpProtocol *ContentFtpProtocol `json:"ftp_protocol,omitempty"`
	// 地址 - OBS (obs远端地址，obs数据源使用) - MONGODB (MONGODB数据源类型主机IP地址，多个IP:PORT, 使用\",\"隔开) - REDIS (redis服务地址，多个IP:PORT, 使用\",\"隔开)

	Address *string `json:"address,omitempty"`
	// Access Key ID - 数据源为OBS，DIS类型时需要配置

	Ak *string `json:"ak,omitempty"`
	// Secret Access Key - 数据源为OBS，DIS类型时需要配置

	Sk *string `json:"sk,omitempty"`
	// 桶名称，数据源为OBS时需要配置

	BucketName *string `json:"bucket_name,omitempty"`
	// 是否使用https, 数据源为OBS时需要配置，一般默认使用

	Https *bool `json:"https,omitempty"`
	// 连接字符串，访问url - 数据源为API、LDAP、WEBSOCKE类型时需要配置， - 数据源为MYSQL、ORACLE、DB2、GAUSS200、GAUSS100、TAURUS、POSTGRESQL，且mode配置为professional专业时需要配置

	Url *string `json:"url,omitempty"`
	// 访问API请求方式 - POST - PUT - DELETE - PATCH - GET

	ApiMethod *ContentApiMethod `json:"api_method,omitempty"`
	// 访问WEBSOCKET服务的认证方式 - none - basicauth

	AuthMethod *ContentAuthMethod `json:"auth_method,omitempty"`

	ApiAuthDetail *ApiAuthDetail `json:"api_auth_detail,omitempty"`
	// KAFKA、ACTIVEMQ的服务器地址，多个IP:PORT, 使用\",\"分隔

	Broker *string `json:"broker,omitempty"`
	// 是否开启SSL认证 - 数据源为KAFKA时需要配置，连接MQS内网地址时，若MQS同时开启了SSL与VPC内网明文访问，请选择“否” - 数据源为ARTEMISMQ、ACTIVEMQ、RABBITMQ、IBMMQ时需要配置， - 数据源为HL7时且作为源端时需要配置

	SslEnable *bool `json:"ssl_enable,omitempty"`
	// SSL用户名/应用Key - 数据源为KAFKA且开启SSL认证时需要配置

	SslUsername *string `json:"ssl_username,omitempty"`
	// SSL密码/应用Secret - 数据源为KAFKA且开启SSL认证时需要配置

	SslPassword *string `json:"ssl_password,omitempty"`
	// MONGODB认证源

	MongodbAuthSource *string `json:"mongodb_auth_source,omitempty"`
	// MONGODB集群模式 - true (集群模式) - false （非集群模式）

	MongodbClusterEnable *bool `json:"mongodb_cluster_enable,omitempty"`
	// MONGODB副本集 当MONGODB为非集群模式时配置

	MongodbReplicaSet *string `json:"mongodb_replica_set,omitempty"`
	// 编码格式 - 数据源为GAUSS200、GAUSS100、POSTGRESQL类型时配置\"big5\"， - 数据源为MYSQL、TAURUS类型且mode为default时配置

	Encoding *ContentEncoding `json:"encoding,omitempty"`
	// MYSQL连接超时时间（秒）

	MysqlTimeout *int32 `json:"mysql_timeout,omitempty"`
	// 公钥库密码 - 数据源类型为ACTIVEMQ、ARTEMISMQ、RABBITMQ、IBMMQ且开启SSL认证时需要配置 - 数据源类为HL7且HL7为目标端（position为target）时，并且开启SSL认证时需要配置

	TrustStorePassword *string `json:"trust_store_password,omitempty"`
	// - 数据源类型为ACTIVEMQ、ARTEMISMQ、RABBITMQ、IBMMQ且开启SSL认证时需要配置，公钥库密码 - 数据源类型为HL7且为目标端（position为target），并且开启SSL认证时需要配置，公钥库密码

	TrustStore *string `json:"trust_store,omitempty"`
	// - 数据源类型为ACTIVEMQ、ARTEMISMQ、RABBITMQ、IBMMQ且开启SSL认证时需要配置，公钥库密码 - 数据源类型为HL7且为目标端（position为target），并且开启SSL认证时需要配置，公钥库密码

	TrustStoreFileType *ContentTrustStoreFileType `json:"trust_store_file_type,omitempty"`
	// 数据源类型为ACTIVEMQ、ARTEMISMQ且开启SSL认证时需要配置 - one-way (单向认证) - two-way (双向认证)

	SslAuthMethod *ContentSslAuthMethod `json:"ssl_auth_method,omitempty"`
	// 私钥库文件内容， - 数据源类型为ACTIVEMQ、ARTEMISMQ，开启SSL认证并且认证方式是two-way时需要配置 - 数据源类型HL7且为源端（position为source），并且开启SSL认证时需要配置

	KeyStore *string `json:"key_store,omitempty"`
	// 私钥库文件类型 - 数据源类型为ACTIVEMQ、ARTEMISMQ，开启SSL认证并且认证方式是two-way时需要配置

	KeyStoreFileType *ContentKeyStoreFileType `json:"key_store_file_type,omitempty"`
	// 私钥库密码 - 数据源类型为ACTIVEMQ、ARTEMISMQ，开启SSL认证并且认证方式是two-way时需要配置 - 数据源为HL7类型，为源端（position为source）并且开启SSL认证时需要配置

	KeyStorePassword *string `json:"key_store_password,omitempty"`
	// 私钥库私钥密码 - 数据源类型为ACTIVEMQ、ARTEMISMQ，开启SSL认证并且认证方式是two-way时需要配置 - 数据源为HL7类型，为源端（position为source）并且开启SSL认证时需要配置

	KeyStoreKeyPassword *string `json:"key_store_key_password,omitempty"`
	// DIS通道名称

	DisTunnelName *string `json:"dis_tunnel_name,omitempty"`
	// DIS数据类别 - JSON

	DisDataType *ContentDisDataType `json:"dis_data_type,omitempty"`
	// DIS配置类别 - senior (高级) - basic (基础)

	DisSettingType *ContentDisSettingType `json:"dis_setting_type,omitempty"`
	// DIS Endpoint，当setting_type为senior时填写

	DisEndpoint *string `json:"dis_endpoint,omitempty"`
	// DIS Region，当setting_type为senior时填写

	DisRegion *string `json:"dis_region,omitempty"`
	// DIS源端项目id，当setting_type为senior时填写

	DisSourceProjectId *string `json:"dis_source_project_id,omitempty"`
	// HL7数据源方向 - source (源端) - target (目标端)

	Hl7Position *ContentHl7Position `json:"hl7_position,omitempty"`
	// HL7是否开启白名单设置

	Hl7WhitelistEnable *bool `json:"hl7_whitelist_enable,omitempty"`
	// HL7白名单 - 允许同步数据到源端HL7的服务器地址，当HL7为源端（position为source）并且开启白名单设置(open_whitelist为true)时填写

	Hl7Whitelist *string `json:"hl7_whitelist,omitempty"`
	// LDAP安全认证类型

	LdapSecurityAuthType *ContentLdapSecurityAuthType `json:"ldap_security_auth_type,omitempty"`
	// RABBITMQ虚拟主机

	RabbitmqVirtualHost *string `json:"rabbitmq_virtual_host,omitempty"`
	// RABBITMQ SSL认证协议 - TLS

	RabbitmqSslProtocol *ContentRabbitmqSslProtocol `json:"rabbitmq_ssl_protocol,omitempty"`
	// SAP客户端号

	SapClient *string `json:"sap_client,omitempty"`
	// SAP实例编号

	SapSysnr *string `json:"sap_sysnr,omitempty"`
	// SNMP网络协议 - udp - tcp

	SnmpNetworkProtocol *ContentSnmpNetworkProtocol `json:"snmp_network_protocol,omitempty"`
	// SNMP版本号

	SnmpVersion *ContentSnmpVersion `json:"snmp_version,omitempty"`
	// SNMP团体名，用于访问SNMP管理代理的身份认证，相当于访问密码

	SnmpCommunity *string `json:"snmp_community,omitempty"`
	// IBMMQ字符集标识

	IbmmqCcsId *string `json:"ibmmq_ccs_id,omitempty"`
	// IBMMQ队列管理器

	IbmmqQueueManager *string `json:"ibmmq_queue_manager,omitempty"`
	// IBMMQ通道名称

	IbmmqChannel *string `json:"ibmmq_channel,omitempty"`
	// IBMMQ密钥算法套件

	IbmmqCipherSuite *string `json:"ibmmq_cipher_suite,omitempty"`
	// HDFS URL - 数据源为MRSHIVE、MRSHDFS、FIHDFS、FIHIVE类型时配置 - fihadfs (/fdi/autotest)

	HdfsPath *string `json:"hdfs_path,omitempty"`
	// 机机交互用户名 - 数据源为MRSHIVE、MRSHDFS、MRSHBASE、MRSKAFKA、FIHDFS、FIHIVE、FIKAFKA类型时配置

	PrincipalName *string `json:"principal_name,omitempty"`
	// - 用户认证文件，文件获取方式参考《ROMA Connect API参考》的“附录>获取数据源配置文件”章节 - 将获取到的文件打包成zip文件，文件名配置在config_file_name中，内容以BASE64编码形式放到config_file_content。 - 数据源为MRSHIVE、MRSHDFS、MRSHBASE、MRSKAFKA、FIHDFS、FIHIVE、FIKAFKA类型时配置

	ConfigFileName *string `json:"config_file_name,omitempty"`
	// - 用户认证文件内容，config_file_name对应的文件内容BASE64编码 - 数据源为MRSHIVE、MRSHDFS、MRSHBASE、MRSKAFKA、FIHDFS、FIHIVE、FIKAFKA类型时配置

	ConfigFileContent *string `json:"config_file_content,omitempty"`
	// 连接器实例ID，连接器发布后对应的实例ID

	ConnectionInstanceId *string `json:"connection_instance_id,omitempty"`
	// 连接器对应的数据源参数，值按实际填写

	ConnectorParams *interface{} `json:"connector_params,omitempty"`
}

func (o Content) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Content struct{}"
	}

	return strings.Join([]string{"Content", string(data)}, " ")
}

type ContentGauss100Version struct {
	value string
}

type ContentGauss100VersionEnum struct {
	V100_R003_C20 ContentGauss100Version
	V300_R001_C00 ContentGauss100Version
}

func GetContentGauss100VersionEnum() ContentGauss100VersionEnum {
	return ContentGauss100VersionEnum{
		V100_R003_C20: ContentGauss100Version{
			value: "V100R003C20",
		},
		V300_R001_C00: ContentGauss100Version{
			value: "V300R001C00",
		},
	}
}

func (c ContentGauss100Version) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentGauss100Version) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentMode struct {
	value string
}

type ContentModeEnum struct {
	DEFAULT       ContentMode
	PROFESSIONAL  ContentMode
	MULTI_ADDRESS ContentMode
}

func GetContentModeEnum() ContentModeEnum {
	return ContentModeEnum{
		DEFAULT: ContentMode{
			value: "default",
		},
		PROFESSIONAL: ContentMode{
			value: "professional",
		},
		MULTI_ADDRESS: ContentMode{
			value: "multiAddress",
		},
	}
}

func (c ContentMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentCdcMode struct {
	value string
}

type ContentCdcModeEnum struct {
	XSTREAM  ContentCdcMode
	LOGMINER ContentCdcMode
}

func GetContentCdcModeEnum() ContentCdcModeEnum {
	return ContentCdcModeEnum{
		XSTREAM: ContentCdcMode{
			value: "xstream",
		},
		LOGMINER: ContentCdcMode{
			value: "logminer",
		},
	}
}

func (c ContentCdcMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentCdcMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentFtpConnectMode struct {
	value string
}

type ContentFtpConnectModeEnum struct {
	ACTIVE  ContentFtpConnectMode
	PASSIVE ContentFtpConnectMode
}

func GetContentFtpConnectModeEnum() ContentFtpConnectModeEnum {
	return ContentFtpConnectModeEnum{
		ACTIVE: ContentFtpConnectMode{
			value: "active",
		},
		PASSIVE: ContentFtpConnectMode{
			value: "passive",
		},
	}
}

func (c ContentFtpConnectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentFtpConnectMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentFtpProtocol struct {
	value string
}

type ContentFtpProtocolEnum struct {
	SFTP ContentFtpProtocol
	FTP  ContentFtpProtocol
}

func GetContentFtpProtocolEnum() ContentFtpProtocolEnum {
	return ContentFtpProtocolEnum{
		SFTP: ContentFtpProtocol{
			value: "sftp",
		},
		FTP: ContentFtpProtocol{
			value: "ftp",
		},
	}
}

func (c ContentFtpProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentFtpProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentApiMethod struct {
	value string
}

type ContentApiMethodEnum struct {
	POST   ContentApiMethod
	PUT    ContentApiMethod
	DELETE ContentApiMethod
	PATCH  ContentApiMethod
	GET    ContentApiMethod
}

func GetContentApiMethodEnum() ContentApiMethodEnum {
	return ContentApiMethodEnum{
		POST: ContentApiMethod{
			value: "POST",
		},
		PUT: ContentApiMethod{
			value: "PUT",
		},
		DELETE: ContentApiMethod{
			value: "DELETE",
		},
		PATCH: ContentApiMethod{
			value: "PATCH",
		},
		GET: ContentApiMethod{
			value: "GET",
		},
	}
}

func (c ContentApiMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentApiMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentAuthMethod struct {
	value string
}

type ContentAuthMethodEnum struct {
	NONE      ContentAuthMethod
	BASICAUTH ContentAuthMethod
}

func GetContentAuthMethodEnum() ContentAuthMethodEnum {
	return ContentAuthMethodEnum{
		NONE: ContentAuthMethod{
			value: "none",
		},
		BASICAUTH: ContentAuthMethod{
			value: "basicauth",
		},
	}
}

func (c ContentAuthMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentAuthMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentEncoding struct {
	value string
}

type ContentEncodingEnum struct {
	BIG5     ContentEncoding
	DEC8     ContentEncoding
	CP850    ContentEncoding
	HP8      ContentEncoding
	KOI8R    ContentEncoding
	LATIN1   ContentEncoding
	LATIN2   ContentEncoding
	SWE7     ContentEncoding
	ASCII    ContentEncoding
	UJIS     ContentEncoding
	SJIS     ContentEncoding
	HEBREW   ContentEncoding
	TIS620   ContentEncoding
	KOI8U    ContentEncoding
	GB2312   ContentEncoding
	GREEK    ContentEncoding
	CP1250   ContentEncoding
	GBK      ContentEncoding
	LATIN5   ContentEncoding
	ARMSCII8 ContentEncoding
	UTF8     ContentEncoding
	UCS2     ContentEncoding
	CP866    ContentEncoding
	KEYBCS2  ContentEncoding
	MACCE    ContentEncoding
	MACROMAN ContentEncoding
	CP852    ContentEncoding
	LATIN7   ContentEncoding
	UTF8MB4  ContentEncoding
	CP1251   ContentEncoding
	UTF16    ContentEncoding
	UTF16LE  ContentEncoding
	CP1256   ContentEncoding
	CP1257   ContentEncoding
	UTF32    ContentEncoding
	BINARY   ContentEncoding
	GEOSTD8  ContentEncoding
	CP932    ContentEncoding
	EUCJPMS  ContentEncoding
	GB18030  ContentEncoding
}

func GetContentEncodingEnum() ContentEncodingEnum {
	return ContentEncodingEnum{
		BIG5: ContentEncoding{
			value: "big5",
		},
		DEC8: ContentEncoding{
			value: "dec8",
		},
		CP850: ContentEncoding{
			value: "cp850",
		},
		HP8: ContentEncoding{
			value: "hp8",
		},
		KOI8R: ContentEncoding{
			value: "koi8r",
		},
		LATIN1: ContentEncoding{
			value: "latin1",
		},
		LATIN2: ContentEncoding{
			value: "latin2",
		},
		SWE7: ContentEncoding{
			value: "swe7",
		},
		ASCII: ContentEncoding{
			value: "ascii",
		},
		UJIS: ContentEncoding{
			value: "ujis",
		},
		SJIS: ContentEncoding{
			value: "sjis",
		},
		HEBREW: ContentEncoding{
			value: "hebrew",
		},
		TIS620: ContentEncoding{
			value: "tis620",
		},
		KOI8U: ContentEncoding{
			value: "koi8u",
		},
		GB2312: ContentEncoding{
			value: "gb2312",
		},
		GREEK: ContentEncoding{
			value: "greek",
		},
		CP1250: ContentEncoding{
			value: "cp1250",
		},
		GBK: ContentEncoding{
			value: "gbk",
		},
		LATIN5: ContentEncoding{
			value: "latin5",
		},
		ARMSCII8: ContentEncoding{
			value: "armscii8",
		},
		UTF8: ContentEncoding{
			value: "utf8",
		},
		UCS2: ContentEncoding{
			value: "ucs2",
		},
		CP866: ContentEncoding{
			value: "cp866",
		},
		KEYBCS2: ContentEncoding{
			value: "keybcs2",
		},
		MACCE: ContentEncoding{
			value: "macce",
		},
		MACROMAN: ContentEncoding{
			value: "macroman",
		},
		CP852: ContentEncoding{
			value: "cp852",
		},
		LATIN7: ContentEncoding{
			value: "latin7",
		},
		UTF8MB4: ContentEncoding{
			value: "utf8mb4",
		},
		CP1251: ContentEncoding{
			value: "cp1251",
		},
		UTF16: ContentEncoding{
			value: "utf16",
		},
		UTF16LE: ContentEncoding{
			value: "utf16le",
		},
		CP1256: ContentEncoding{
			value: "cp1256",
		},
		CP1257: ContentEncoding{
			value: "cp1257",
		},
		UTF32: ContentEncoding{
			value: "utf32",
		},
		BINARY: ContentEncoding{
			value: "binary",
		},
		GEOSTD8: ContentEncoding{
			value: "geostd8",
		},
		CP932: ContentEncoding{
			value: "cp932",
		},
		EUCJPMS: ContentEncoding{
			value: "eucjpms",
		},
		GB18030: ContentEncoding{
			value: "gb18030",
		},
	}
}

func (c ContentEncoding) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentEncoding) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentTrustStoreFileType struct {
	value string
}

type ContentTrustStoreFileTypeEnum struct {
	JKS ContentTrustStoreFileType
	TS  ContentTrustStoreFileType
}

func GetContentTrustStoreFileTypeEnum() ContentTrustStoreFileTypeEnum {
	return ContentTrustStoreFileTypeEnum{
		JKS: ContentTrustStoreFileType{
			value: "jks",
		},
		TS: ContentTrustStoreFileType{
			value: "ts",
		},
	}
}

func (c ContentTrustStoreFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentTrustStoreFileType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentSslAuthMethod struct {
	value string
}

type ContentSslAuthMethodEnum struct {
	ONE_WAY ContentSslAuthMethod
	TWO_WAY ContentSslAuthMethod
}

func GetContentSslAuthMethodEnum() ContentSslAuthMethodEnum {
	return ContentSslAuthMethodEnum{
		ONE_WAY: ContentSslAuthMethod{
			value: "one-way",
		},
		TWO_WAY: ContentSslAuthMethod{
			value: "two-way",
		},
	}
}

func (c ContentSslAuthMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentSslAuthMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentKeyStoreFileType struct {
	value string
}

type ContentKeyStoreFileTypeEnum struct {
	KS  ContentKeyStoreFileType
	JKS ContentKeyStoreFileType
}

func GetContentKeyStoreFileTypeEnum() ContentKeyStoreFileTypeEnum {
	return ContentKeyStoreFileTypeEnum{
		KS: ContentKeyStoreFileType{
			value: "ks",
		},
		JKS: ContentKeyStoreFileType{
			value: "jks",
		},
	}
}

func (c ContentKeyStoreFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentKeyStoreFileType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentDisDataType struct {
	value string
}

type ContentDisDataTypeEnum struct {
	JSON ContentDisDataType
}

func GetContentDisDataTypeEnum() ContentDisDataTypeEnum {
	return ContentDisDataTypeEnum{
		JSON: ContentDisDataType{
			value: "JSON",
		},
	}
}

func (c ContentDisDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentDisDataType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentDisSettingType struct {
	value string
}

type ContentDisSettingTypeEnum struct {
	BASIC  ContentDisSettingType
	SENIOR ContentDisSettingType
}

func GetContentDisSettingTypeEnum() ContentDisSettingTypeEnum {
	return ContentDisSettingTypeEnum{
		BASIC: ContentDisSettingType{
			value: "basic",
		},
		SENIOR: ContentDisSettingType{
			value: "senior",
		},
	}
}

func (c ContentDisSettingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentDisSettingType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentHl7Position struct {
	value string
}

type ContentHl7PositionEnum struct {
	SOURCE ContentHl7Position
	TARGET ContentHl7Position
}

func GetContentHl7PositionEnum() ContentHl7PositionEnum {
	return ContentHl7PositionEnum{
		SOURCE: ContentHl7Position{
			value: "source",
		},
		TARGET: ContentHl7Position{
			value: "target",
		},
	}
}

func (c ContentHl7Position) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentHl7Position) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentLdapSecurityAuthType struct {
	value string
}

type ContentLdapSecurityAuthTypeEnum struct {
	NO     ContentLdapSecurityAuthType
	SIMPLE ContentLdapSecurityAuthType
}

func GetContentLdapSecurityAuthTypeEnum() ContentLdapSecurityAuthTypeEnum {
	return ContentLdapSecurityAuthTypeEnum{
		NO: ContentLdapSecurityAuthType{
			value: "no",
		},
		SIMPLE: ContentLdapSecurityAuthType{
			value: "simple",
		},
	}
}

func (c ContentLdapSecurityAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentLdapSecurityAuthType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentRabbitmqSslProtocol struct {
	value string
}

type ContentRabbitmqSslProtocolEnum struct {
	TLS ContentRabbitmqSslProtocol
}

func GetContentRabbitmqSslProtocolEnum() ContentRabbitmqSslProtocolEnum {
	return ContentRabbitmqSslProtocolEnum{
		TLS: ContentRabbitmqSslProtocol{
			value: "TLS",
		},
	}
}

func (c ContentRabbitmqSslProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentRabbitmqSslProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentSnmpNetworkProtocol struct {
	value string
}

type ContentSnmpNetworkProtocolEnum struct {
	UDP ContentSnmpNetworkProtocol
	TCP ContentSnmpNetworkProtocol
}

func GetContentSnmpNetworkProtocolEnum() ContentSnmpNetworkProtocolEnum {
	return ContentSnmpNetworkProtocolEnum{
		UDP: ContentSnmpNetworkProtocol{
			value: "udp",
		},
		TCP: ContentSnmpNetworkProtocol{
			value: "tcp",
		},
	}
}

func (c ContentSnmpNetworkProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentSnmpNetworkProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ContentSnmpVersion struct {
	value int32
}

type ContentSnmpVersionEnum struct {
	E_0 ContentSnmpVersion
	E_1 ContentSnmpVersion
	E_3 ContentSnmpVersion
}

func GetContentSnmpVersionEnum() ContentSnmpVersionEnum {
	return ContentSnmpVersionEnum{
		E_0: ContentSnmpVersion{
			value: 0,
		}, E_1: ContentSnmpVersion{
			value: 1,
		}, E_3: ContentSnmpVersion{
			value: 3,
		},
	}
}

func (c ContentSnmpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ContentSnmpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
