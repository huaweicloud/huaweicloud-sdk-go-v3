# 0.0.48 2021-06-21

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ChangeBaremetalServerName`新增响应参数`server_tags`、`enterprise_project_id`、`group`

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - [Issue 22](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/22): 修正接口`ListAddonInstances`的响应参数`status`可选值
- _特性变更_
    - 无

### HuaweiCloud SDK CDN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListDomains`移除请求参数`user_domain_id`
    - 接口名称调整：`ShowRefer` -> `ShowReferer`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowTemplateDetail`新增请求参数：
        - `template_url`
        - `create_time`
        - `last_modify_time`
        - `can_update`
        - `can_delete`
        - `need_hub`

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口:
        - `CreateRecordCallbackConfig`
        - `ShowRecordCallbackConfig`
        - `UpdateRecordCallbackConfig`
        - `DeleteRecordCallbackConfig`
        - `ListRecordCallbackConfigs`
        - `UpdateRecordRule`
        - `ShowRecordRule`
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `CreateRecordConfig` -> `CreateRecordRule`
        - `DeleteRecordConfig` -> `DeleteRecordRule`
        - `ListRecordConfigs` -> `ListRecordRules`
    - 移除接口：
        - `ShowTraffic`
        - `ShowBandwidth`
        - `ShowOnlineUsers`

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowGroups`的响应参数`partitions`类型调整： `array[string]` -> `array[integer]`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
    - 无
- _解决问题_
    - 修复编译失败的问题
- _特性变更_
    - 无

# 0.0.47 2021-06-10

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口`ListFreeResources`、`ListFreeResourceUsages`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CDN

- _新增特性_
    - 支持内容分发网络服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DRS

- _新增特性_
    - 支持数据复制服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 新增支持接口
        - `ImportFunction`
        - `ExportFunction`
        - `AsyncInvokeReservedFunction`
        - `DeleteReservedInstanceById`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK OSM

- _新增特性_
    - 支持工单管理服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持接口`SetBinlogClearPolicy`、`ShowBinlogClearPolicy`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListOffSiteInstances`新增请求参数`offset`、`limit`

# 0.0.46 2021-06-04

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - [Issue 20](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/20): 修复`extendParam`类型定义错误的问题
- _特性变更_
    - 接口`DeleteCluster`新增请求参数`tobedeleted`

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口`ShowQuotas`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 新增支持接口`ListComplexQueryDevice`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
    - 支持`GaussDBforNoSQL`服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持接口`ShowQuotas`
- _解决问题_
    - 无
- _特性变更_
    - 接口`StartInstanceRestartAction`的请求参数`restart`类型调整：string -> object

# 0.0.45 2021-05-25

### HuaweiCloud SDK AS

- _新增特性_
    - 新增支持接口：
        - `ListApiVersions`
        - `ShowApiVersion`
        - `BatchProtectScalingInstances`
        - `BatchRemoveScalingInstances`
        - `CreateScalingTagInfo`
        - `BatchResumeScalingPolicies`
        - `BatchPauseScalingPolicies`
        - `PauseScalingGroup`
        - `BatchSetScalingInstancesStandby`
        - `BatchUnsetScalingInstancesStandby`
        - `ResumeScalingPolicy`
        - `PauseScalingPolicy`
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `ExecuteScalingPolicies` -> `BatchDeleteScalingPolicies`
        - `EnableOrDisableScalingGroup` -> `ResumeScalingGroup`
        - `UpdateScalingGroupInstance` -> `BatchAddScalingInstances`
        - `CompleteLifecycleAction` -> `AttachCallbackInstanceLifeCycleHook`
    - 移除接口：
        - `DeleteScalingTags`
    - `ListScalingGroups` 接口新增参数 `enterprise_project_id`
    - `ListScalingActivityV2Logs` 接口新增参数 `log_id`

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：
        - 查询月度成本 `ListCustomerBillsMonthlyBreakDown`
        - 查询订单可用折扣 `ListOrderDiscounts`
- _解决问题_
    - 无
- _特性变更_
    - 查询客户消费记录接口 `ListSubCustomerResFeeRecords` 增加 query 参数：bill_date_begin 和 bill_date_end

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 新增支持接口：停止流水线 `StopPipelineNew`
- _解决问题_
    - 无
- _特性变更_
    - 移除接口 `StartPipeline`、`StopPipeline`

### HuaweiCloud SDK DMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：（查询项目标签）ShowProjectTags -> ShowQueueProjectTags

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListEnterpriseProject`的请求参数`offset`从必填改为非必填

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 新增支持接口：
        - `ListFunctionAsyncInvokeConfig`
        - `ShowFunctionAsyncInvokeConfig`
        - `DeleteFunctionAsyncInvokeConfig`
        - `UpdateFunctionAsyncInvokeConfig`
- _解决问题_
    - 无
- _特性变更_
    - 接口`DeleteVersionAlias`、`UpdateVersionAlias`、`ShowVersionAlias`的请求参数名称调整：`name` -> `alias_name`
    - 接口`DeleteFunctionTrigger`、`UpdateTrigger`、`ShowFunctionTrigger`的请求参数名称调整：`triggerId` -> `trigger_id`

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口 `CreateUsers` 新增请求体参数和响应体参数 _access_mode_
    - 接口 `DeleteBindingDevice` 将请求体参数 _authentication_code_ 设置为必填参数

### HuaweiCloud SDK Kafka

- _新增特性_
    - 新增支持接口：
        - `CreateInstanceUser`
        - `BatchDeleteInstanceUsers`
        - `ShowInstanceUsers`
        - `ShowTopicAccessPolicy`
        - `UpdateTopicAccessPolicy`
        - `ShowKafkaTopicPartitionDiskusage`
        - `ShowInstanceMessages`
        - `ResetUserPassword`
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `ShowInstanceTags` -> `ShowKafkaTags`
        - `ShowProjectTags` -> `ShowKafkaProjectTags`
        - `BatchCreateOrDeleteInstanceTag` -> `BatchCreateOrDeleteKafkaTag`
    - 接口 `BatchCreateOrDeleteInstanceTag` 请求体名称调整：`BatchCreateOrDeleteInstanceTagRequestBody`
      -> `BatchCreateOrDeleteKafkaTagRequestBody`
    - 接口 `UpdateSinkTaskQuota` 请求体 `UpdateSinkTaskQuotaReq` 的请求参数 `sink_max_tasks` 数据类型调整：String → Integer

### HuaweiCloud SDK OMS

- _新增特性_
    - 支持对象存储迁移服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RabbitMQ

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `BatchCreateOrDeleteInstanceTag` -> `BatchCreateOrDeleteRabbitMqTag`
        - `ShowProjectTags` -> `ShowRabbitMqProjectTags`
        - `ShowInstanceTags` -> `ShowRabbitMqTags`
    - 接口 `BatchCreateOrDeleteInstanceTag` 请求体名称调整：BatchCreateOrDeleteInstanceTagRequestBody ->
      BatchCreateOrDeleteRabbitMqTagRequestBody

# 0.0.43-rc 2021-05-14

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 解决了使用接口`NovaShowKeypair`获取秘钥，结果解析异常的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListInstances`的响应字段`type`新增结果值: `CLOUDSSD`、`LOCALSSD`
    - 接口`ListBackups`新增可选的请求参数`complete_version`
    - 将接口`ListSlowlogStatistics`的请求参数`type`从非必填改为必填

# 0.0.42-rc 2021-05-10

### HuaweiCloud SDK BMS

- _新增特性_
    - 新增支持接口`BatchCreateBaremetalServerTags`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口`MigrateAz`、`ListAz2Migrate`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/17): 修复`EpDetailType`枚举定义错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK KPS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 19](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/19): 修复`ListKeypairs`响应体类型的问题
- _特性变更_
    - 无

### HuaweiCloud SDK MRS

- _新增特性_
    - 新增支持相关接口：
        - `BatchDeleteClusterTags`
        - `CreateClusterTag`
        - `DeleteClusterTag`
        - `ListClusterTags`
        - `ListAllTags`
        - `BatchCreateClusterTags`
        - `ListClustersByTags`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListOffSiteInstances`的响应体名称调整：`OffSiteInstanceListResponse` -> `OffSiteInstanceListResponseBody`
    - 接口`ListOffSiteInstances`的响应字段名称调整：`offsite_backup_instances` -> `offsite_backup_instance`

# 0.0.41-rc 2021-04-30

### HuaweiCloud SDK BCS

- _新增特性_
    - 新增支持查询异步操作结果的接口`ListOpRecord`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持的接口
        - 查询集群均衡设置 `ShowShardingBalancer`
        - 设置集群均衡开关 `SetBalancerSwitch`
        - 设置集群均衡活动时间窗 `SetBalancerWindow`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK HSS

- _新增特性_
    - 新增支持查询弹性云服务器状态列表的接口`ListHosts`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 查询账号配额接口`ShowDomainQuota`的请求参数`type`增加可选值：
        - `assigment_group_mp`
        - `assigment_agency_mp`
        - `assigment_group_ep`
        - `assigment_user_ep`

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 下线订阅管理相关接口：
        - `ListSubscriptions`
        - `CreateSubscription`
        - `UpdateSubscription`
        - `ShowSubscription`
        - `DeleteSubscription`

### HuaweiCloud SDK MPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateMpeCallBack`新增请求参数`language`、`sky_switch`
    - 接口`CreateTranscodingTask`的请求参数`subtitle_type`可选值调整为`0`、`1`、`2`

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 获取项目详情接口`ShowProjectInfoV4`的响应体新增字段`project_code`

# 0.0.40-rc 2021-04-15

### HuaweiCloud SDK Moderation

- _新增特性_
    - 支持Moderation内容审核服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Image

- _新增特性_
    - 支持Image图像识别服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持管理数据库的接口
        - `CreateSqlserverDatabase`
        - `DeleteSqlserverDatabase`
        - `ListSqlserverDatabases`
    - 新增支持管理用户的接口
        - `CreateSqlserverDbUser`
        - `ListSqlserverDbUsers`
        - `ListAuthorizedSqlserverDbUsers`
        - `DeleteSqlserverDbUser`
        - `AllowSqlserverDbUserPrivilege`
        - `RevokeSqlserverDbUserPrivilege`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口`DeleteDatabaseUser`、`DeleteDatabaseRole`、`ShowConnectionStatistics`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListIssuesV4`, `ListChildIssuesV4`响应体新增字段`closed_time` 、`id` 、`created_time`

### HuaweiCloud SDK VPC

- _新增特性_
    - 无
- _解决问题_
    - 修复问题，开放vpc和子网的标签
- _特性变更_
    - 无

# 0.0.39-rc 2021-03-30

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 修复查询消息接口没有时间戳字段的问题
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowIssueV4`、`UpdateIssueV4`的响应体 `IssueResponseV4`增加`name`属性
    - 将接口`ShowProjectWorkHours`、`ListProjectWorkHours`的响应体`ShowProjectWorkHoursResponseBody`下的属性`work_time`
      修改为`work_date`

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 将接口`PublishMessage`的请求参数`protocol`从必填改为非必填
    - 将接口 `PublishMessage` 请求体的 `PublishMessageRequestBody` 类属性 `subject` 由必填改为非必填

# 0.0.38-rc 2021-03-26

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决自动获取 domainId 的场景下报错校验 AK SK 校验失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 解决 ListLiveStreamsOnline 接口响应体反序列化失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 解决 ListSlowlogStatistics 接口响应体部分字段为空的问题
- _特性变更_
    - 无

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 将接口 `ListSlowlogStatistics` 下的 `property` 属性从必填调整为非必填

# 0.0.37-rc 2021-03-19

### HuaweiCloud SDK Core

- _新增特性_
    - 支持文件上传
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 解决 ListFlavors 接口响应体反序列化出错的问题
- _特性变更_
    - 无

# 0.0.36-rc 2021-03-16

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 申请包周期弹性公网IP接口增加 `enterprise_project_id` 字段

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 修复接口无法调用的问题
- _特性变更_
    - 无

# 0.0.35-rc 2021-03-15

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 当接口返回体带有 `\n` 等特殊字符时，反序列化失败问题修复
- _特性变更_
    - 当用户传入的 `endpoint` 未带协议前缀时，支持自动加上 `https` 前缀

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 使用资源分组创建告警规则时，维度字段调整为非必填，即 `CreateAlarmRequestBody` 类中的属性 `metric` 对应类调整：
      `MetricInfoForAlarm` → `MetricForAlarm`
    - 将 `MetricsDimension`类中的属性资源维度(`name`)字段调整为必填字段，涉及接口：批量查询监控数据、添加监控数据、创建资源分组、更新分组

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口：
        - 根据备份恢复新实例 `RestoreNewInstance`
        - 查询实例节点会话 `ListSessions`
        - 终结实例节点会话 `DeleteSession`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 新增支持接口：查询云服务器组详情 `ShowServerGroup`
- _解决问题_
    - 无
- _特性变更_
    - 云服务器获取密码接口名调整：`ShowWindowsServerPasswordResponseBody` → `ShowServerPasswordResponseBody`
    - 云服务器清除密码接口名调整：`DeleteWindowsServerPassword` → `DeleteServerPassword`

### HuaweiCloud SDK ELB

- _新增特性_
    - 新增支持接口：查询当前租户下的后端服务器列表 `ListAllMembers`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口获取依赖包列表 `ListDependencies` 响应体的属性 `size` 类型调整：string → int64

### HuaweiCloud SDK IAM

- _新增特性_
    - 新增支持接口：
        - 查询身份提供商详情 `KeystoneShowIdentityProvider`
        - 注册身份提供商 `KeystoneCreateIdentityProvider`
        - 更新身份提供商 `KeystoneUpdateIdentityProvider`
        - 删除身份提供商 `KeystoneDeleteIdentityProvider`
        - 获取联邦认证token(OpenId Connect Id token方式) `CreateTokenWithIdToken`
- _解决问题_
    - 无
- _特性变更_
    - 下线接口：获取联邦认证unscoped token `CreateUnscopeTokenByIdpInitiated`

### HuaweiCloud SDK IMS

- _新增特性_
    - 新增支持接口：
        - 按标签查询镜像 `ListImageByTags`
        - 查询租户所有镜像标签 `ListImagesTags`
        - 查询镜像标签 `ListImageTags`
        - 添加镜像标签 `AddImageTag`
        - 删除镜像标签 `DeleteImageTag`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 支持设备接入服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 新增支持接口：
        - 创建工作项类型自定义字段 `CreateCustomfields`
        - 查询人均bug `ShowBugsPerDeveloper`
        - 查询需求按时完成率 `ShowCompletionRate`
        - 查询缺陷密度 `ShowBugDensityV2`
        - 获取项目详情 `ShowProjectInfoV4`
- _解决问题_
    - 修改接口命名错误：`ShowtIssueCompletionRate` → `ShowIssueCompletionRate`
- _特性变更_
    - `ListProjectV4` 接口响应体中的 `created_time` 和 `updated_time` 属性类型调整：string → int64

### HuaweiCloud SDK RDS

- _新增特性_
    - 支持 Postgresql 相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK SCM

- _新增特性_
    - 支持 SSL 证书管理服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.34-rc 2021-02-27

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 优化 README 文档描述及 CHANGELOG 日志格式

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：`WindowsBaremetalServerCleanPwd` → `DeleteWindowsBareMetalServerPassword`
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 所有类型为 `float32` 的属性统一调整为 `float64`

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - 修复参数 `Content-Type` 必填且返回 "Unsupported Content Type" 错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 修复类 `CreateAlarmResponse` 出现循环依赖的问题
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口：获取慢日志下载链接 `DownloadSlowlog`、获取错误日志下载链接 `DownloadErrorlog`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DGC

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：
        - `ModifyJob` → `UpdateJob`
        - `ModifyScript` → `UpdateScript`
        - `ModifyResource` → `UpdateResource`
- _特性变更_
    - 无

### HuaweiCloud SDK DIS

- _新增特性_
    - 无
- _解决问题_
    - 修复类出现循环依赖的问题
        - `CommitCheckpointRequest`
        - `PutRecordsRequest`
        - `CreateAppRequest`
        - `UpdatePartitionCountRequest`
- _特性变更_
    - 无

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：`ListMenbers` → `ListMembers`
    - 修复接口 `ShowMember` 需要传 query 参数 `X-Auth-Token` 的问题
- _特性变更_
    - 无

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：`ModifyEnterpriseProject` → `UpdateEnterpriseProject`
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 修复 `KeystoneUserResult` 类中属性名的错误定义问题：`pwd_stength` → `pwd_strength`
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：
        - `DoManualBackup` → `CreateManualBackup`
        - `ModifyConfiguration` → `UpdateConfiguration`
        - `ModifyInstanceConfiguration` → `UpdateInstanceConfiguration`
    - 修复类 `CreateInstanceResponse` 和 `CreateConfigurationResponse` 出现循环依赖的问题
    - 修复接口 `CreateInstance` 不可用的问题
- _特性变更_
    - 接口 `StartInstanceAction` 请求中单机转主备场景增加 `is_auto_pay` 属性

# 0.0.33-rc 2021-02-07

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 修复请求体类型为string时请求失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 查询镜像支持的OS列表(ListOsVersions)接口返回体属性 `os_bit` 数据类型调整：string → int32

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口：获取直播播放日志(ListLiveSampleLogs)、创建直播域名(CreateDomain)、删除直播域名(DeleteDomain)、修改直播域名(UpdateDomain)、查询直播域名(
      ShowDomain)、创建直播域名映射关系(CreateDomainMapping)、删除直播域名映射关系(DeleteDomainMapping)
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK MRS

- _新增特性_
    - 支持MapReduce服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.32-rc 2021-01-30

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口：查询当前支持的API版本信息列表、查询指定API版本信息、设置审计日志策略、查询审计日志策略、获取审计日志列表
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整: UpdateAutoTerminateTimeServer → UpdateServerAutoTerminateTime

### HuaweiCloud SDK EVS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建云硬盘接口支持指定专属存储池ID
    - 查询配额相关接口属性 `allocated` 类型调整: string → int

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 查询IAM用户详情接口响应体增加属性`access_mode`

# 0.0.31-rc 2021-01-25

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 设置默认`DefaultTimeout`为120秒

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：查询订单可用折扣
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DNS

- _新增特性_
    - 支持云解析服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 新增支持接口：修改云服务器云主机销毁时间
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.30-rc 2021-01-15

### HuaweiCloud SDK Core

- _新增特性_
    - Region管理支持使用`ValueOf`方法获取`region`信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DGC

- _新增特性_
    - 新增支持接口：作业相关接口（增删改查）、脚本相关接口（增删改查）、资源相关接口（增删改查）
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建/查询用户接口响应字段 `is_domain_owner` 类型调整：string → boolean

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口：查询直播中的流信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持接口：查询跨区域备份策略、设置跨区域备份策略、获取跨区域备份列表、查询跨区域可恢复时间段、查询跨区域备份实例列表、查询API版本列表、查询指定的API版本信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK SWR

- _新增特性_
    - 支持容器镜像服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.29-beta 2020-12-31

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 修复查询裸金属服务器详情接口问题
    - 修复查询裸金属服务器详情列表接口问题
    - 完善查询裸金属服务器网卡信息接口字段
- _特性变更_
    - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
    - 新增支持接口：查询当前账号访问权限（ShowAccountStatus）
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 修正接口返回数据类型避免反序列化失败：
        - 查询所有实例列表接口：响应参数`Tags`类型调整 Tag → ResourceTag
        - 查询慢日志接口：响应参数`duration`类型调整 int32 → string
        - 查询大key分析详情接口：响应参数`db`类型调整 int32 → string
        - 查询热key分析详情接口：响应参数`db`类型调整 int32 → string
- _特性变更_
    - 无

### HuaweiCloud SDK DGC

- _新增特性_
    - 支持数据湖治理中心服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DIS

- _新增特性_
    - 支持数据接入服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建实例接口响应类型调整

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 修正消息发布接口请求参数：Object → Map<String, String>
- _特性变更_
    - 无

# 0.0.28-beta 2020-12-28

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 调用临时AK/SK返回异常问题修复
- _特性变更_
    - 无

### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 修改缓存接口port类型为integer
- _特性变更_
    - 无

# 0.0.27-beta 2020-12-25

### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 修复SDK在Mac操作系统上无法正常编译的问题
- _特性变更_
    - 接口ListInstances请求Query参数名称调整：id → instance_id

### HuaweiCloud SDK DDS

- _新增特性_
    - 支持文档数据库服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 修复SDK在Mac操作系统上无法正常编译的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RabbitMQ

- _新增特性_
    - 无
- _解决问题_
    - 修复SDK在Mac操作系统上无法正常编译的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RMS

- _新增特性_
    - 新增支持接口：资源查询相关接口、Region查询相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.26-beta 2020-12-23

### HuaweiCloud SDK Core

- _新增特性_
    - 支持Region管理，用户新建客户端时可以通过{Service}Region传入，无需自己配置endpoint。配置Region后，支持自动回填ProjectId/DomainId
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：查询用量单位列表
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - ShowMetricData接口字段更新

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口：查询播放画像信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK MPC

- _新增特性_
    - 新增支持接口：视频增强模板相关接口，声道合并任务相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 支持云数据库服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 消息模板名称中字段类型调整

# 0.0.25-beta 2020-12-15

### HuaweiCloud SDK CCE

- _新增特性_
    - 支持云容器引擎服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 创建监听器接口返回为空问题修复
    - 证书列表查询接口返回非列表问题修复
- _特性变更_
    - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 新增支持接口：更新函数预留实例个数
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK NAT

- _新增特性_
    - 无
- _解决问题_
    - 修复批量创建DNAT规则失败的问题
- _特性变更_
    - 无

# 0.0.24-beta 2020-12-04

### HuaweiCloud SDK SMN

- _新增特性_
    - 支持消息通知服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.23-beta 2020-11-30

### HuaweiCloud SDK BCS

- _新增特性_
    - 支持区块链服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BMS

- _新增特性_
    - 支持裸金属服务器
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：查询使用量列表，设置/取消包周期资源到期转按需
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 新增支持接口：迁移资源接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CES

- _新增特性_
    - 新增支持接口：
        - 查询事件监控列表
        - 查询某一个事件监控详情
        - 创建资源分组
        - 更新资源分组
        - 删除资源分组
        - 查询所有资源分组
        - 修改告警规则
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DCS

- _新增特性_
    - 支持分布式缓存服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - [Issue 21](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/21) 开放查询SSH密钥详情接口

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 支持函数工作流服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 新增支持接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - Live服务客户端名字修正：LiveAPIClient → LiveClient

# 0.0.22-beta 2020-11-17

### HuaweiCloud SDK AS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 8](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/8) Fix the problem that creating scaling
      policy failed.
- _特性变更_
    - 无

### HuaweiCloud SDK DMS

- _新增特性_
    - 支持分布式消息服务，提供Kafka专享版和RabbitMQ专享版
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建虚拟机接口（按需和包周期）增加 `dry_run` 属性，表示是否预检此次请求

### HuaweiCloud SDK NAT

- _新增特性_
    - 支持NAT网关服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK VPC

- _新增特性_
    - 支持网络ACL相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.21-beta 2020-11-11

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 同步[Pull requests #11](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/pull/11)修改
- _特性变更_
    - 无

# HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建存储库接口(CreateVault)新增存储库turbo类型
    - 修改策略接口(UpdatePolicy)删除多余字段

### HuaweiCloud SDK CES

- _新增特性_
    - 新增接口响应示例，调整字段描述
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 生成客户端文件的名字调整：devcloudpipeline_client → cloudPipeline_client
    - 生成Meta文件的名字调整：devcloudpipeline_meta → cloudPipeline_meta

### HuaweiCloud SDK DevStar

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改接口参数，调整示例代码

# 0.0.20-beta 2020-11-02

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建告警规则接口增加资源分组字段
    - 查询告警历史接口响应体metadata修改，只返回total字段

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 创建转发规则接口参数名错误问题修复
- _特性变更_
    - 无

# 0.0.19-beta 2020-10-31

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - Fix: query参数中包含枚举变量时请求失败
    - [Issue 7](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/7) 解决json.Marshal变成{}对象的问题
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 新增支持接口：TAG标签相关接口
- _解决问题_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/17) 修复ListBackups接口响应体问题
- _特性变更_
    - 无

### HuaweiCloud SDK CTS

- _新增特性_
    - 新增支持接口：ListQuotas（查询租户追踪器配额信息）
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整，原有的`*EP`接口展开为`*EnterpriseProject`，涉及接口：
        1. ListEP → ListEnterpriseProject
        2. CreateEP → CreateEnterpriseProject
        3. ShowEP → ShowEnterpriseProject
        4. ModifyEP → ModifyEnterpriseProject
        5. EnableEP → EnableEnterpriseProject
        6. ShowEPQuota → ShowEnterpriseProjectQuota
        7. ShowResourceBindEP → ShowResourceBindEnterpriseProject
        8. DisableEP → DisableEnterpriseProject

### HuaweiCloud SDK Iam

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整，涉及接口：
        1. KeystoneCreateUserTokenByPasswordAndMFA → KeystoneCreateUserTokenByPasswordAndMfa
        2. CreateUnscopeTokenByIDPInitiated → CreateUnscopeTokenByIdpInitiated

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 新增支持接口：迭代信息、用户信息、项目成员、项目信息、项目指标、项目统计等相关方法的接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.18-beta 2020-10-20

### HuaweiCloud SDK ELB

- _新增特性_
    - 增加v2版本接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.17-beta 2020-10-14

### HuaweiCloud SDK BSS

- _新增特性_
    - 伙伴中心支持导出产品目录价
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 新增直播V2接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.16-beta 2020-10-12

### HuaweiCloud SDK CTS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 删除v1版本废弃接口

### HuaweiCloud SDK DevStar

- _新增特性_
    - 无
- _解决问题_
    - 服务客户端认证类型调整为全局认证，即GlobalCredentials
- _特性变更_
    - 无

# 0.0.15-beta 2020-09-30

### HuaweiCloud SDK AS

- _新增特性_
    - 支持弹性伸缩服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.14-beta 2020-09-24

### HuaweiCloud SDK BSS

- _新增特性_
    - 无
- _解决问题_
    - 修复BssClient类无法加载的问题
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口ListPublicips调整，企业项目ID不支持多值查询

# 0.0.13-beta 2020-09-18

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 修正接口参数类型
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 运营能力开放接口更新

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 修复查询弹性公网IP不支持传入多个值的问题
- _特性变更_
    - 无

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 修复部分接口参数错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 调整接口描述

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改禁播参数描述

# 0.0.12.1-beta 2020-09-16

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口参数类型错误
- _特性变更_
    - 无

### HuaweiCloud SDK All

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 可选参数类型为list时，参数类型变更为指针，例如：[]string 将变更为 *[]string

# 0.0.12-beta 2020-09-11

### HuaweiCloud SDK KPS

- _新增特性_
    - 支持KPS服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EVS

- _新增特性_
    - 支持EVS服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 修复返回值类型定义错误的问题
- _特性变更_
    - 无

# 0.0.11-beta 2020-09-09

### HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 修复资源相关接口类型错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK VPC

- _新增特性_
    - 无
- _解决问题_
    - 修复安全组相关类型错误的问题
- _特性变更_
    - 无

# 0.0.10-beta 2020-09-04

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 修复yaml中没有定义format的整型枚举参数无法生成枚举代码的问题
- _特性变更_
    - 调整Http请求头的User-Agent信息

# 0.0.9-beta 2020-08-28

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 支持流水线服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 新增支持接口：弹性公网IP标签相关接口和共享带宽相关接口
- _解决问题_
    - 批量创建带宽接口，修改billing_info字段
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 新增支持接口：console一致性相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 支持项目管理服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK VPC

- _新增特性_
    - 新增支持接口：安全组、安全组规则和可用IP数
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.8-beta 2020-08-25

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 可选枚举变量类型变更为指针类型。

### HuaweiCloud SDK ELB

- _新增特性_
    - 支持弹性负载均衡服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.7-beta 2020-08-20

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决部分枚举类型变量无法正常读取的问题。
- _特性变更_
    - 以_开头的枚举类型变量名称添加前缀 'E'。

# 0.0.6-beta 2020-08-20

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决当结构体包含枚举类型变量场景下部分依赖丢失的问题。
- _特性变更_
    - 无

# 0.0.5-beta 2020-08-19

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决枚举类型反序列化失败的问题。
    - 解决部分场景下云服务异常响应反序列化失败问题。
- _特性变更_
    - 无

# 0.0.4-beta 2020-08-18

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - Go 基础类型默认值序列化丢失问题修复
- _特性变更_
    - 无

# 0.0.3-beta 2020-08-14

### HuaweiCloud SDK APIG

- _新增特性_
    - 支持API网关
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 支持能力开放服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.2-beta 2020-08-11

### HuaweiCloud SDK Core

- _新增特性_
    - 支持临时AK/SK认证模式
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 支持云备份服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
    - 支持云测服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 支持弹性云服务器服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 支持弹性公网IP服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EVS

- _新增特性_
    - 支持云硬盘服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 支持镜像服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 支持视频直播服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK MPC

- _新增特性_
    - 支持媒体转码服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 3.0.1-beta 2020-07-30

### 首次发布

- _已支持服务_
    - Classroom
    - 云审计服务（CTS）
    - 模板引擎（DevStar）
    - 企业管理服务（EPS）
    - 统一身份认证服务（IAM）
    - 标签管理服务（TMS）
    - 虚拟私有云（VPC）
