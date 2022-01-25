# 0.0.76 2022-01-25

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateApiV2`新增请求参数 `content_type`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `publish_time`、`roma_app_name`、`ld_api_id`、`api_group_info`、`content_type`
  - 接口`UpdateApiV2`:
    - 新增请求参数 `content_type`
    - 新增响应参数 `publish_time`、`roma_app_name`、`ld_api_id`、`api_group_info`、`content_type`
  - 接口`ListApiRuntimeDefinitionV2`新增响应参数 `content_type`
  - 接口`ListApiVersionDetailV2`:
    - 新增响应参数 `roma_app_name`、`ld_api_id`、`api_group_info`、`content_type`
    - 响应参数`publish_time`类型调整 `date-time` -> `string`

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobs`:
    - 移除响应参数 `simple`
    - 响应参数`name`、`values`改为必填
  - 接口`UpdateJob`请求参数`name`、`values`改为必填
  - 接口`CreateAndStartRandomClusterJob`请求参数`name`、`values`改为必填
  - 接口`CreateJob`请求参数`name`、`values`改为必填
  - 接口`ListClusters`移除响应参数 `config_status`

### HuaweiCloud SDK CES

- _新增特性_
  - 支持接口`ListAlarmHistories`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`DeleteRuleset`、`SetDefaulTemplate`、`ShowTasklog`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTask`新增请求参数 `endpoint_id`
  - 接口`CreateRuleset`新增请求参数 `custom_attributes`
  - 接口`ListTemplateRules`:
    - 新增请求参数 `tags`
    - 新增响应参数 `rule_config_list`

### HuaweiCloud SDK DevStar

- _新增特性_
  - 支持接口`ShowRepositoryByCloudIde`、`ListTemplates`
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
  - 接口`CreateLoginToken`新增响应参数`session_user_id`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`ListEngineProducts`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `dr_enable`
  - 接口`ShowInstance`新增响应参数 `dr_enable`
  - 接口`ListProducts`:
    - 新增响应参数 `Hourly`、`Monthly`
    - 移除响应参数 `hourly`、`honthly`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSqlAlarmRule`移除请求参数 `language`
  - 接口`UpdateSqlAlarmRule`移除请求参数 `language`
  - 接口`ListSqlAlarmRules`:
    - 新增响应参数 `template_name`、`status`
    - 移除响应参数 `language`
  - 接口`CreateKeywordsAlarmRule`移除请求参数 `language`、`eps_id`、`eps_name`、`is_time_range_relative`
  - 接口`UpdateKeywordsAlarmRule`:
    - 移除请求参数 `language`、`eps_id`、`eps_name`、`is_time_range_relative`
    - 移除响应参数 `eps_id`、`eps_name`、`is_time_range_relative`
  - 接口`ListKeywordsAlarmRules`:
    - 新增响应参数 `template_name`、`status`
    - 移除响应参数 `language`、`eps_id`、`eps_name`、`is_time_range_relative`
  - 接口`ListAccessConfig`:
    - 新增请求参数 `access_config_tag_list`
    - 新增响应参数 `access_config_tag`
  - 接口`CreateAccessConfig`:
    - 新增请求参数 `access_config_tag`
    - 新增响应参数 `access_config_tag`
  - 接口`UpdateAccessConfig`:
    - 新增请求参数 `access_config_tag`
    - 新增响应参数 `access_config_tag`
  - 接口`DeleteAccessConfig`新增响应参数 `access_config_tag`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 支持接口`ListEngineProducts`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.75 2022-01-17

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 修复接口`CreateNodePool`的请求体结构错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`ListFunctionAsyncInvocations`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowGaussMySqlInstanceInfo`新增响应参数`used`
  - 接口`UpdateInstanceMonitor`新增请求参数`monitor_switch`
  - 接口`UpdateInstanceMonitor`的请求参数`period`类型修改： `string` -> `int32`
  - 接口`ShowGaussMySqlProjectQuotas`移除响应参数`mode`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListApiVersionNew`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`新增响应参数 `az_desc`

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddUserToApp`新增请求参数 `id`、`roles`
  - 接口`DownloadAssetArchive`移除响应参数 `apps`、`tasks`
  - 接口`ImportAsset`移除请求参数 `apps`、`tasks`
  - 接口`DeleteAsset`请求参数`tasks`改为必填
  - 接口`ShowMqsInstance`新增响应参数`access_user`

# 0.0.74 2022-01-10

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListOrderCouponsByOrderId`新增响应参数 `coupon_max_use_quantity`、`coupon_group`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`新增请求参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`ListClusters`新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`UpdateCluster`:
    - 新增请求参数 `customSan`、`containerNetwork`
    - 新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`ShowCluster`新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`DeleteCluster`新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`CreateNode`新增请求参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ListNodes`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`UpdateNode`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ShowNode`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`DeleteNode`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`CreateNodePool`新增请求参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ListNodePools`新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`UpdateNodePool`:
    - 新增请求参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
    - 新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ShowNodePool`新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`DeleteNodePool`新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`UpdateDomainFullConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ListStacks`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListStacksByTag`

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowPlanList`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowPlans`:
    - 请求参数`limit`类型调整 `int64` -> `int32`
    - 请求参数`offset`类型调整 `int64` -> `int32`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListMigrationTask`新增响应参数 `ecs_tenant_private_ip`
  - 接口`ShowMigrationTask`新增响应参数 `ecs_tenant_private_ip`
  - 接口`StopMigrationTask`新增响应参数 `ecs_tenant_private_ip`

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`移除响应参数`lb_ip_address`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateServers`新增请求参数 `delete_on_termination`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateFunctionConfig`:
    - 新增请求参数 `is_stateful_function`
    - 新增响应参数 `is_stateful_function`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `lb_ip_address`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`请求参数`type`新增枚举值`GaussDB(for openGauss)`, 移除枚举值`GaussDB(openGauss)`
  - 接口`ListInstances`移除响应参数 `related_instance`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`GlanceListImages`响应参数`active_at`改为非必填
  - 接口`GlanceShowImage`响应参数`active_at`改为非必填
  - 接口`GlanceUpdateImage`响应参数`active_at`改为非必填

### HuaweiCloud SDK IoTAnalytics

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDatasource`请求参数`site_id`改为必填
  - 接口`UpdateDataSource`请求参数`site_id`改为必填

### HuaweiCloud SDK KPS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListKeypairs`:
    - 新增请求参数 `limit`、`marker`
    - 新增响应参数 `page_info`
  - 接口`ListFailedTask`新增请求参数 `limit`、`offset`
  - 接口`ListRunningTask`新增请求参数 `limit`、`offset`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`StartInstanceSingleToHaAction`:
    - 新增请求参数 `ad_domain_info`
    - 移除请求参数 `password`

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowMqsInstance`移除响应参数 `connect_dn`
  - 接口`ListMqsInstanceTopics`移除响应参数 `policies`
  - 接口`ShowDetailsOfAppV2`新增响应参数 `roma_app_type`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEndpointService`新增响应参数 `pool_id`
  - 接口`ListEndpointService`新增响应参数 `domain_id`
  - 接口`UpdateEndpointService`新增响应参数 `pool_id`
  - 接口`ListEndpointInfoDetails`新增响应参数 `enable_status`、`specification_name`

### HuaweiCloud SDK VSS

- _新增特性_
  - 支持接口`AuthorizeDomains`、`ShowTasks`、`CreateTasks`、`ShowResults`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDomains`响应参数`auth_status`新增枚举值`skip`

# 0.0.73 2021-12-25

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLink`新增请求参数 `id`
  - 接口`ShowClusterDetail`新增响应参数 `endpointDomainName`、`isScheduleBootOff`、`failedReasons`、`components`、`createFrom`、`resourceId`、`flavorType`、`workSpaceId`、`trial`
  - 接口`UpdateJob`新增请求参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
  - 接口`ShowJobs`新增响应参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
  - 接口`CreateAndStartRandomClusterJob`:
    - 新增请求参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
    - 新增响应参数 `submissions`
    - 移除响应参数 `name`、`validation-result`
  - 接口`CreateJob`新增请求参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
  - 接口`StartJob`新增响应参数 `execute-date`
  - 接口`UpdateLink`新增请求参数 `id`
  - 接口`ShowLink`新增响应参数 `id`
  - 接口`ListClusters`:
    - 新增响应参数 `bakExpectedStartTime`、`bakKeepDay`、`createFrom`、`resourceId`、`flavorType`、`workSpaceId`、`trial`、`components`
    - 移除响应参数 `version`

### HuaweiCloud SDK CloudBuild

- _新增特性_
  - 支持接口`ShowHistoryDetails`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudTable

- _新增特性_
  - 支持表格存储服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowPlanJournals`、`ShowIssuesByPlanId`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`CheckParameters`、`ListTaskParameter`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRules`新增响应参数 `rule_desc`
  - 接口`ListRulesets`新增响应参数 `is_devcloud_project_default`、`is_default_template`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持代码托管服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CopyInstance`新增请求参数 `backup_format`

### HuaweiCloud SDK DevStar

- _新增特性_
  - 支持以下接口：
    - `ShowApplicationV3`
    - `UpdateApplication`
    - `ShowApplicationDependentResources`
    - `DeleteApplicationV4`
    - `ShowApplicationResDeleteStatus`
    - `ListApplicationsV6`
    - `ShowDeploymentJobs`
    - `CreateDeploymentJobs`
    - `ShowPipelineLastStatusV2`
    - `ListPipelineTemplates`
    - `StartPipeline`
    - `ListProjectsV4`
    - `ShowRepositoryStatisticalDataV2`
    - `CheckRepositoryDuplicateName`
    - `ShowApplicationReleaseRepositories`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTemplateV3`新增响应参数 `dependents`
  - 接口`ListTemplatesV2`新增响应参数 `dependents`、`dependent_services`
  - 接口`ShowJobDetail`新增响应参数 `show_type`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NovaListServerActions`移除响应参数`updated_at`

### HuaweiCloud SDK IEF

- _新增特性_
  - 支持智能边缘平台服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTAnalytics

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDatasource`:
    - 新增请求参数 `site_id`
    - 移除请求参数 `instance_id`
  - 接口`ShowAllDataSource`:
    - 新增响应参数 `site_id`
    - 移除响应参数 `instance_id`
  - 接口`UpdateDataSource`:
    - 新增请求参数 `site_id`
    - 新增响应参数 `site_id`
    - 移除请求参数 `instance_id`
    - 移除响应参数 `instance_id`
  - 接口`ShowDataSource`:
    - 新增响应参数 `site_id`
    - 移除响应参数 `instance_id`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`请求参数`partition_num`改为非必填
  - 接口`RestartManager`新增响应参数 `result`、`instance_id`
  - 接口`ListProducts`:
    - 新增响应参数 `hourly`、`honthly`
    - 移除响应参数 `Hourly`、`Monthly`

### HuaweiCloud SDK KPS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFailedTask`响应参数`task_time`类型调整 `int64` -> `string`
  - 接口`ListRunningTask`响应参数`task_time`类型调整 `int64` -> `string`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCharts`新增请求参数 `offset`、`limit`
  - 接口`ListNotificationTemplates`新增请求参数 `offset`、`limit`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTranscodingTask`:
    - 新增请求参数 `auto_volume_setting`
    - 请求参数`volume`新增枚举值`original`
  - 接口`ListTranscodingTask`新增响应参数 `av_parameters`
  - 接口`CreateWatermarkTemplate`新增请求参数 `template_id`
  - 接口`ListWatermarkTemplate`新增响应参数 `template_id`
  - 接口`UpdateWatermarkTemplate`新增请求参数 `template_id`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持以下接口：
    - `CreateScalingPolicy`
    - `ShowClusterDetails`
    - `UpdateClusterScaling`
    - `ListHosts`
    - `CreateAndExecuteJob`
    - `ListExecuteJob`
    - `ShowJobExes`
    - `DeleteJobExecution`
    - `CreateCluster`
    - `ShowAgencyMapping`
    - `UpdateAgencyMapping`
    - `ShowJobExeListNew`
    - `CreateExecuteJob`
    - `ShowSingleJobExe`
    - `StopJob`
    - `ShowSqlResultWithJob`
    - `BatchDeleteJobs`
    - `ExecuteSql`
    - `ShowSqlResult`
    - `CancelSql`
    - `ShowHdfsFileList`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTasks`新增响应参数 `group_type`、`success_record_error_reason`、`skip_record_error_reason`、`save_prefix`
  - 接口`ShowTask`新增响应参数 `group_type`、`success_record_error_reason`、`skip_record_error_reason`、`save_prefix`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListApiVersion`、`ShowApiVersion`、`SearchQueryScaleComputeFlavors`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ROMA

- _新增特性_
  - 支持应用与数据集成平台服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SA

- _新增特性_
  - 支持态势感知服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.72 2021-12-17

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowVersion`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`:
    - 新增响应参数 `bundle_url`、`visitor_id`、`visitor_name`、`visitor_domain_name`
    - 移除响应参数 `action_list`、`role`、`role_id`、`sub_org`
  - 接口`ListOrgInstances`:
    - 新增响应参数 `visitor_id`、`visitor_name`、`visitor_domain_name`
    - 移除响应参数 `action_list`、`role`、`role_id`、`sub_org`
  - 接口`ListInstances`:
    - 新增响应参数 `visitor_id`、`visitor_name`、`visitor_domain_name`
    - 移除响应参数 `action_list`、`role`、`role_id`、`sub_org`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 支持接口`ListRtcAbnormalEvents`、`ListRtcAbnormalEventDimension`
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
  - 接口`ListEvents`:
    - 响应参数`event_count`类型调整 `string` -> `int32`
    - 响应参数`latest_occur_time`类型调整 `string` -> `int64`
  - 接口`BatchListMetricData`响应参数`variance`类型调整 `string` -> `double`
  - 接口`ListResourceGroup`响应参数`type_statistics`类型调整 `string` -> `int32`
  - 接口`ListEventDetail`:
    - 响应参数`event_users`类型调整 `string` -> `array`
    - 响应参数`event_sources`类型调整 `string` -> `array`

### HuaweiCloud SDK IoTAnalytics

- _新增特性_
  - 支持物联网数据分析服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEditingJob`新增请求参数 `input`、`edit_settings`
  - 接口`ListEditingJob`新增响应参数 `input`、`edit_settings`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持以下接口：
    - `RecognizeThailandIdcard`
    - `RecognizeMyanmarIdcard`
    - `RecognizeMyanmarDriverLicense`
    - `RecognizeChileIdCard`
    - `RecognizeThailandLicensePlate`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.71 2021-12-10

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEvents`的请求参数`relation`更新枚举值为`[AND, OR, NOT]`

### HuaweiCloud SDK APM

- _新增特性_
  - 支持接口`ShowMasterAddress`、`ListAkSk`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateScalingGroup`新增请求参数 `allowed_address_pairs`
  - 接口`ListScalingGroups`新增响应参数 `allowed_address_pairs`
  - 接口`ShowScalingGroup`新增响应参数 `allowed_address_pairs`
  - 接口`UpdateScalingGroup`新增请求参数 `allowed_address_pairs`

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBlockchainDetail`响应参数`node_cnt`类型调整 `int32` -> `int64`
  - 接口`ShowBlockchainFlavors`新增请求参数 `limit`、`offset`

### HuaweiCloud SDK CampusGo

- _新增特性_
  - 支持园区智能体服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CGS

- _新增特性_
  - 支持容器安全服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ShowExtensionAuthorization`、`CreateExtensionAuthorization`、`CheckInstanceAccess`、`UpdateInstanceActivity`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRtcClientQosDetails`新增请求参数`stream_id`

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`ShowTaskCmetrics`、`ListTemplateRules`、`ListTaskRuleset`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTaskDetail`新增响应参数 `is_access`、`trigger_type`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `readonly_domain_name`
  - 接口`ShowInstance`新增响应参数 `readonly_domain_name`

### HuaweiCloud SDK DDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `time_zone`
  - 接口`RestartInstance`新增响应参数 `instanceId`、`instanceName`、`jobId`

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持接口`ShowOpenApiCalledRecords`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDocWatermark`新增请求参数 `readonly_password`
  - 接口`ShowScanJobs`:
    - 新增请求参数 `offset`
    - 移除请求参数 `page`
  - 接口`ShowScanJobResults`:
    - 新增请求参数 `offset`
    - 移除请求参数 `page`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectFaceByFile`移除响应参数 `landmark`、`gender`、`yaw_angle`、`roll_angle`、`pitch_angle`、`headpose`、`smile`、`skin`、`ethnic`
  - 接口`DetectFaceByUrl`移除响应参数 `landmark`、`gender`、`yaw_angle`、`roll_angle`、`pitch_angle`、`headpose`、`smile`、`skin`、`ethnic`
  - 接口`DetectFaceByBase64`移除响应参数 `landmark`、`gender`、`yaw_angle`、`roll_angle`、`pitch_angle`、`headpose`、`smile`、`skin`、`ethnic`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口：
    - `ListInstanceTags`
    - `ListProjectTags`
    - `BatchTagAction`
    - `ShowInstanceMonitorExtend`
    - `UpdateInstanceMonitor`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListGaussMySqlInstances`:
    - 新增请求参数 `private_ip`、`tags`
    - 新增响应参数 `tags`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`请求参数`region`改为必填

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CheckMetadata`

### HuaweiCloud SDK HiLens

- _新增特性_
  - 支持华为HiLens服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK KMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ImportKeyMaterial`新增请求参数 `encrypted_privatekey`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordRule`的请求参数`default_record_config`改为必填

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `ListLogHistogram`
    - `ListHost`
    - `ListHostGroup`
    - `UpdateHostGroup`
    - `CreateHostGroup`
    - `DeleteHostGroup`
    - `ListAccessConfig`
    - `UpdateAccessConfig`
    - `CreateAccessConfig`
    - `DeleteAccessConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`RestoreExistInstance`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SWR

- _新增特性_
  - 支持接口`ListQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateNamespaceAuth`修改请求体名称`UpdateNamespaceAuthReq` -> `UpdateNamespaceAuthRequestBody`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListStatistics`请求参数`to`、`from`改为必填
  - 接口`ListQpsTimeline`请求参数`from`、`to`改为必填
  - 接口`ListBandwidthTimeline`请求参数`from`、`to`改为必填
  - 接口`ListTopAbnormal`请求参数`from`、`to`改为必填
  - 接口`ShowEvent`新增响应参数 `cookie`
  - 接口`CreatePremiumHost`:
    - 新增请求参数 `Content-Type`
    - 请求参数`type`类型调整 `string` -> `enum`
  - 接口`ListPremiumHost`新增请求参数 `Content-Type`
  - 接口`UpdatePremiumHost`:
    - 新增请求参数 `Content-Type`
    - 响应参数`type`类型调整 `string` -> `enum`
  - 接口`DeletePremiumHost`新增请求参数 `Content-Type`
  - 接口`ShowPremiumHost`:
    - 新增请求参数 `Content-Type`
    - 响应参数`type`类型调整 `string` -> `enum`
  - 接口`UpdatePremiumHostProtectStatus`新增请求参数 `Content-Type`

# 0.0.70 2021-11-29

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ResettingAppSecretV2`新增响应参数`roma_app_type`

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEntityMetric`响应参数`values `类型调整： `object` -> `array`

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBackups`新增响应参数`provider_id`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowCluster`新增响应参数`cidrs`

### HuaweiCloud SDK DBSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SwitchAgent`和`SwitchRiskRule`的响应参数名称调整： `status` -> `result`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateInstance`新增请求参数`port`

### HuaweiCloud SDK DSC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowScanJobs`新增响应参数`start_time`

### HuaweiCloud SDK EVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CinderExportToImage`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NovaShowServer`和`NovaListServersDetails`新增响应参数`os:scheduler_hints`

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持以下接口：
    - `ListApiVersions`
    - `ListSecurityPolicies`
    - `CreateSecurityPolicy`
    - `ShowSecurityPolicy`
    - `UpdateSecurityPolicy`
    - `DeleteSecurityPolicy`
    - `ListSystemSecurityPolicies`
    - `ListQuotaDetails`
    - `ChangeLoadbalancerChargeMode`
    - `BatchUpdatePoliciesPriority`
    - `UpdateIpList`
    - `BatchDeleteIpList`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowQuotaDefaults`
  - 接口`ListFlavors`和`ShowFlavor`新增响应参数`flavor_sold_out`和`lcu`, 移除响应参数`availability_zone_ids`
  - 接口`ShowQuota`新增响应参数`members_per_pool`
  - 接口`CreateCertificate`和`UpdateCertificate`新增请求参数`enc_certificate`和`enc_private_key`
  - 接口`ListCertificates`和`ShowCertificate`新增响应参数`enc_certificate`和`enc_private_key`
  - 接口`CreateLoadBalancer`新增请求参数`prepaid_options`、`autoscaling`、`id`
  - 接口`ListLoadBalancers`新增请求参数`elb_virsubnet_type`和`autoscaling`,新增响应参数`autoscaling`和`ip_version`
  - 接口`UpdateLoadBalancer`新增请求参数`prepaid_options`和`autoscaling`,新增响应参数`loadbalancer_id`、`order_id`、`autoscaling`、`ip_version`
  - 接口`ShowLoadBalancer`新增响应参数`autoscaling`和`ip_version`
  - 接口`ShowLoadBalancerStatus`新增响应参数`id`、`type`、`provisioning_status`
  - 接口`CreateListener`新增请求参数`security_policy_id`、`enhance_l7policy_enable`
  - 接口`ListListeners`新增请求参数`enhance_l7policy_enable`和`member_instance_id`，新增响应参数`security_policy_id`、`transparent_client_ip_enable`、`enhance_l7policy_enable`
  - 接口`UpdateListener`新增请求参数`security_policy_id`、`enhance_l7policy_enable`，新增响应参数`security_policy_id`、`transparent_client_ip_enable`、`enhance_l7policy_enable`
  - 接口`ShowListener`新增响应参数`security_policy_id`、`transparent_client_ip_enable`、`enhance_l7policy_enable`
  - 接口`ListPools`新增请求参数`listener_id`和`member_instance_id`
  - 接口`ListMembers`新增请求参数`ip_version`和`member_type`，新增响应参数`member_type`和`instance_id`
  - 接口`UpdateMember`、`ShowMember`、`ListAllMembers`新增响应参数`member_type`和`instance_id`
  - 接口`CreateL7Policy`新增请求参数`priority`、`redirect_url_config`、`fixed_response_config`、`conditions`
  - 接口`ListL7Policies`新增请求参数`priority`,新增响应参数`redirect_url_config`和`fixed_response_config`
  - 接口`UpdateL7Policy`新增请求参数`priority`、`redirect_url_config`、`fixed_response_config`、`rules`,新增响应参数`redirect_url_config`和`fixed_response_config`
  - 接口`ShowL7Policy`新增响应参数`redirect_url_config`和`fixed_response_config`
  - 接口`CreateL7Rule`新增请求参数`conditions`
  - 接口`ListL7Rules`和`ShowL7Rule`新增响应参数`conditions`
  - 接口`UpdateL7Rule`新增请求参数和响应参数`conditions`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListVersionAliases`新增响应参数`ListVersionAliases`
  - 接口`CreateDependency`和`UpdateDependency`的请求参数`name`改为必填
  - 接口`CreateEvent`的请求参数`name`和`content`改为必填

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数`enable_force_switch`

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateGraph`的请求参数`graphSizeTypeIndex`类型调整： `integer` -> `string`

### HuaweiCloud SDK KMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateKey`的请求参数`key_alias`改为必填
  - 接口`EnableKey`、`CancelKeyDeletion`、`ListKeys`、`ListKeyDetail`、`ShowPublicKey`、`ListGrants`、`DeleteImportedKeyMaterial`、`EnableKeyRotation`、`DisableKeyRotation`、`ShowKeyRotationStatus`的请求参数`key_id`改为必填
  - 接口`DeleteKey`的请求参数`key_id`和`pending_days`改为必填
  - 接口`ListKeys`新增请求参数`enterprise_project_id`
  - 接口`CreateRandom`的请求参数`random_data_length`改为必填
  - 接口`CreateDatakey`和`CreateDatakeyWithoutPlaintext`的请求参数`key_id`和`datakey_length`改为必填
  - 接口`EncryptDatakey`的请求参数`key_id`、`plain_text`、`datakey_plain_length`改为必填
  - 接口`DecryptDatakey`的请求参数`key_id`、`cipher_text`、`datakey_cipher_length`改为必填
  - 接口`UpdateKeyAlias`的请求参数`key_id`和`key_alias`改为必填
  - 接口`UpdateKeyDescription`的请求参数`key_id`和`key_description`改为必填
  - 接口`CreateGrant`的请求参数`key_id`、`grantee_principal`、`operations`改为必填
  - 接口`CancelGrant`和`CancelSelfGrant`的请求参数`key_id`和`grant_id`改为必填
  - 接口`EncryptData`的请求参数`key_id`和`plain_text`改为必填
  - 接口`DecryptData`的请求参数`cipher_text`改为必填
  - 接口`CreateParametersForImport`的请求参数`key_id`和`wrapping_algorithm`改为必填
  - 接口`ImportKeyMaterial`的请求参数`key_id`、`import_token`、`encrypted_key_material`改为必填
  - 接口`UpdateKeyRotationInterval`的请求参数`key_id`和`rotation_interval`改为必填
  - 接口`CreateKmsTag`新增请求参数`sequence`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持以下接口：
    - `ChangeProxyScale`
    - `SearchQueryScaleFlavors`
    - `ShowInformationAboutDatabaseProxy`
    - `StartDatabaseProxy`
    - `StopDatabaseProxy`
    - `UpdateReadWeight`
    - `ChangeTheDelayThreshold`
    - `ShowDrReplicaStatus`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPostgresqlDatabases`的响应参数`size`类型调整： `int32` -> `int64`

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 修复枚举值中包含中文描述导致参数错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIgnoreRule`新增响应参数`items`
  - 接口`ListEvent`新增请求参数`attacks`
  - 接口`ShowEvent`新增响应参数`host_id`
  - 接口`UpdateHost`新增响应参数`certificatename`

# 0.0.69 2021-11-25

### HuaweiCloud SDK AOM

- _新增特性_
  - 支持应用维护管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CheckBackendConnectivity`

### HuaweiCloud SDK APM

- _新增特性_
  - 支持应用性能管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BCS

- _新增特性_
  - 支持接口`DeleteMemberInvite`、`CreateBlockchainCertByUserName`、`UnfreezeCert`、`FreezeCert`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecordDetails`新增请求参数`statistic_type`和响应参数`bill_date`

### HuaweiCloud SDK CBH

- _新增特性_
  - 支持云堡垒机服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowCluster`新增响应参数`platformVersion`
  - 接口`ListClusters`的请求参数`status`新增枚举值`RollingBack`和`RollbackFailed`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 支持以下接口：
    - `ListRtcRealtimeScaleDimension`
    - `ListRtcRealtimeQuality`
    - `ListRtcRealtimeNetwork`
    - `ListRtcHistoryUsage`
    - `ListRtcHistoryScale`
    - `ListRtcHistoryQuality`
    - `ListRtcClientQosDetails`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持代码检查服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DAS

- _新增特性_
  - 无
- _解决问题_
  - 修正部分接口的请求参数`X-Language`的枚举值
- _特性变更_
  - 无

### HuaweiCloud SDK DBSS

- _新增特性_
  - 支持数据库安全服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`、`ListInstances`新增请求参数`tags`

### HuaweiCloud SDK DeH

- _新增特性_
  - 支持专属主机服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePrePaidPublicip`、`CreatePublicip`新增请求和响应参数`alias`
  - 接口`ShowPublicip`、`UpdatePublicip`新增响应参数`alias`

### HuaweiCloud SDK GES

- _新增特性_
  - 支持接口`ResizeGraph`、`ExpandGraph`、`UploadFromObs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `CreateTransfer`
    - `DeleteTransfer`
    - `UpdateTransfer`
    - `ListTransfers`
    - `ListLogStreams`
    - `RegisterDmsKafkaInstance`
    - `CreateNotificationTemplate`
    - `UpdateNotificationTemplate`
    - `ListNotificationTemplates`
    - `DeleteNotificationTemplate`
    - `ShowNotificationTemplate`
    - `ListNotificationTemplate`
    - `UpdateAlarmRuleStatus`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK KMS

- _新增特性_
  - 支持V2版本的接口
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK NLP

- _新增特性_
  - 支持自然语言处理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeHandwriting`移除响应参数`extracted_data`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSlowlogStatistics`新增请求参数`sort`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 支持弹性文件服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK TMS

- _新增特性_
  - 支持接口`ShowTagQuota`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServicePublicDetails`、`ListServiceDetails`、`ListServiceConnections`、`AcceptOrRejectEndpoint`、`ListEndpoints`、`UpdateEndpointWhite`、`ListEndpointInfoDetails`的响应参数`created_at`和`updated_at`类型调整: `datetime` -> `string`
  - 接口`CreateEndpointService`的请求参数`vpc_id`和`port_id`改为必填
  - 接口`AcceptOrRejectEndpoint`移除响应参数`error`
  - 接口`ListVersionDetails`和`ListSpecifiedVersionDetails`的响应参数`updated`类型调整: `datetime` -> `string`
  - 接口`ListResourceInstances`和`BatchAddOrRemoveResourceInstance`的请求参数`action`改为必填

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEvent`新增请求参数`from`和`to`
  - 接口`ListWhiteblackipRule`新增请求参数`name`

# 0.0.68 2021-11-12

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复响应体类型为`array`或`map`时解析错误的问题
  - 修复请求体数据自动转义的问题
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateNode`新增请求参数`customSan`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateHttpsInfo`新增请求参数`force_redirect_https`

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 支持部署服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAuditLogs`的响应参数名调整： `total_count` -> `total_record`

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持接口`ShowScanJobs`和`ShowScanJobResults`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDocWatermark`新增请求参数`marked_file_password`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListGaussMySqlConfigurations`新增请求参数`offset`和`limit`
  - 接口`ShowGaussMySqlProxy`新增响应参数`id`和`spec_code`

### HuaweiCloud SDK GES

- _新增特性_
  - 支持图引擎服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProPricePlans`移除请求参数`sim_card_id`、`partner`、`package_type`、`sim_type`
  - 接口`ListSimCards`移除请求参数`tag_id`
  - 接口`ListSimPricePlans`移除请求参数`sim_price_plan_id`
  - 接口`StopSimCard`和`ResetSimCard`移除请求参数`price_plan_list`
  - 接口`ShowSimCard`和`ListSimCards`移除响应参数参数`sn`、`supply_code`、`bundle_id`、`test_type`
  - 接口`StopSimCard`移除响应参数`sim_price_plan_list`
  - 接口`ListSimPools`移除响应参数`order_id`
  - 接口`ListSimPricePlans`移除响应参数`partner`、`partner_pid`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJob`新增响应参数`results`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBandwidthDetail`、`ListUsersOfStream`新增请求参数`country`和`protocol`
  - 接口`ListDomainTrafficDetail`、`ListDomainBandwidthPeak`、`ListDomainTrafficSummary`新增请求参数`protocol`
  - 接口`ListTranscodeData`新增请求参数`stream`
  - 接口`ListHistoryStreams`新增请求参数`stream`、`start_time`、`end_time`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `ShowStructTemplate`
    - `CreateStructTemplate`
    - `UpdateStructTemplate`
    - `DeleteStructTemplate`
    - `ShowAomMappingRules`
    - `CreateAomMappingRules`
    - `UpdateAomMappingRules`
    - `DeleteAomMappingRules`
    - `ShowAomMappingRule`
    - `ListNotificationTopics`
    - `CreateSqlAlarmRule`
    - `UpdateSqlAlarmRule`
    - `ListSqlAlarmRules`
    - `DeleteSqlAlarmRule`
    - `CreateKeywordsAlarmRule`
    - `UpdateKeywordsAlarmRule`
    - `ListKeywordsAlarmRules`
    - `DeleteKeywordsAlarmRule`
    - `ListActiveOrHistoryAlarms`
    - `DeleteActiveAlarms`
    - `ListCharts`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEditingJob`请求参数名称调整： `const` -> `consts`
  - 接口`CreateEditingJob`移除请求参数`index`，新增请求参数`overlay_input`、`const`、`mix`
  - 接口`ListEditingJob`新增响应参数`output`
  - 接口`CreateTranscodingTask`新增请求参数`hls_init_count`和`hls_init_interval`, 请求参数`video_enhance`新增可选值`EFFICIENCY`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口（V2）`GetJobExeListNew`新增请求参数`job_id`、`user`、`queue`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeGeneralTable`新增响应参数`confidence`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`CreateSystemIssueV4`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIrs`新增响应参数`sequence`
  - 接口`BatchUpdateIrs`新增请求参数`status_id`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListSlowLogsNew`、`ListErrorLogsNew`、`UpgradeDbVersion`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VAS

- _新增特性_
  - 支持视频分析服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VPC

- _新增特性_
  - 支持接口（V3）: `AddVpcExtendCidr`、`RemoveVpcExtendCidr`、`ListVpcs`、`ShowVpc`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.67 2021-10-25

### HuaweiCloud SDK BCS

- _新增特性_
  - 支持接口`ShowBlockchainFlavors`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIndirectPartners`新增响应参数`account_manager_id`和`account_manager_name`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowHistoryTaskDetails`新增请求参数`create_time`和响应参数`task_type`
  - 接口`ShowHistoryTasks`移除请求参数`create_time`

### HuaweiCloud SDK DNS

- _新增特性_
  - 支持接口`ShowDomainQuota`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSharedBandwidth`新增请求参数`bandwidth_type`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddFacesByFile`、`AddFacesByBase64`、`AddFacesByUrl`新增请求参数`single`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`的请求参数和响应参数`num`、`size`类型调整： `integer` -> `string`

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持全球SIM联接服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ImageSearch

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunSearchPicture`的请求参数名称调整: `isCrop` -> `is_crop`
  - 接口`RunSearchPicture`新增请求参数`box`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJob`新增响应参数`current_task`、`image_name`、`process_percent`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDevices`新增请求参数`status`
  - 接口`CreateRuleAction`新增请求参数`file_path`

### HuaweiCloud SDK OCR

- _新增特性_
  - 新增接口`RecognizeInsurancePolicy`、`RecognizeFinancialStatement`、`RecognizeQualificationCertificate`
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
  - 接口`ShowBackupDownloadLink`新增响应参数`database_name`
  - 接口`ListInstances`新增响应参数`max_iops`和`expiration_time`

### HuaweiCloud SDK SDRS

- _新增特性_
  - 支持存储容灾服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.66 2021-10-19

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowMigrationTask`新增响应参数`backup_id`
  - 接口`ListMigrationTask`新增以下响应参数:
    - `source_instance_name`
    - `source_instance_id`
    - `target_instance_addrs`
    - `target_instance_id`
  - 接口`ListDiagnosisTasks`的响应参数`total_usec_sum`类型调整： `integer` -> `double`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持以下接口：
    - `ListCommonPools`
    - `ListPublicBorderGroups`
    - `ListPublicipPool`
    - `ShowPublicipPool`
    - `ListShareBandwidthTypes`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPublicips`新增请求参数`allow_share_bandwidth_type_any`
  - 接口`NeutronListFloatingIps`的请求参数`limit`的类型调整： `string` -> `integer`
  - 接口`NeutronUpdateFloatingIp`请求体的名称调整： `floatingip` -> `NeutronUpdateFloatingIpRequestBody`
  - 接口`ShowPublicip`新增响应参数`allow_share_bandwidth_types`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTranscodingTask`新增请求参数`hls_init_count`和`hls_init_interval`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 支持VPC终端节点服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.65 2021-10-11

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复响应体为文件流时没有响应状态码的问题
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持以下接口：
    - `CreateAppCodeV2`
    - `CreateAppCodeAutoV2`
    - `ListAppCodesV2`
    - `ShowDetailsOfAppCodeV2`
    - `DeleteAppCodeV2`
    - `DebugApiV2`
    - `BatchPublishOrOfflineApiV2`
    - `ListApiVersionsV2`
    - `ChangeApiVersionV2`
    - `ListApiRuntimeDefinitionV2`
    - `ListApiVersionDetailV2`
    - `DeleteApiByVersionIdV2`
    - `ListAclStrategiesV2`
    - `BatchDeleteAclV2`
    - `CreateAclStrategyV2`
    - `ShowDetailsOfAclPolicyV2`
    - `UpdateAclStrategyV2`
    - `DeleteAclV2`
    - `BatchDeleteApiAclBindingV2`
    - `CreateApiAclBindingV2`
    - `DeleteApiAclBindingV2`
    - `ListAclPolicyBindedToApiV2`
    - `ListApisBindedToAclPolicyV2`
    - `ListApisUnbindedToAclPolicyV2`
    - `ListCustomAuthorizersV2`
    - `CreateCustomAuthorizerV2`
    - `ShowDetailsOfCustomAuthorizersV2`
    - `UpdateCustomAuthorizerV2`
    - `DeleteCustomAuthorizerV2`
    - `ExportApiDefinitionsV2`
    - `ImportApiDefinitionsV2`
    - `CreateVpcChannelV2`
    - `ListVpcChannelsV2`
    - `ShowDetailsOfVpcChannelV2`
    - `DeleteVpcChannelV2`
    - `UpdateVpcChannelV2`
    - `AddingBackendInstancesV2`
    - `ListBackendInstancesV2`
    - `DeleteBackendInstanceV2`
    - `ListLatelyApiStatisticsV2`
    - `ListLatelyGroupStatisticsV2`
    - `CreateGatewayResponseV2`
    - `ListGatewayResponsesV2`
    - `ShowDetailsOfGatewayResponseV2`
    - `UpdateGatewayResponseV2`
    - `DeleteGatewayResponseV2`
    - `ShowDetailsOfGatewayResponseTypeV2`
    - `UpdateGatewayResponseTypeV2`
    - `DeleteGatewayResponseTypeV2`
    - `ListTagsV2`
    - `CreateFeatureV2`
    - `ListFeaturesV2`
    - `CreateInstanceV2`
    - `ListInstancesV2`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSkuInventories`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTask`新增以下响应参数:
    - `create_time`
    - `description`
    - `operate_mode`
    - `related_temp_running_data`
    - `run_status`
    - `update_time`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DisassociateServerVirtualIp`的请求参数`reverse_binding`改为非必填

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectFaceByFile`、`DetectFaceByBase64`、`DetectFaceByUrl`的请求参数`attributes`可选值调整为`2,4,6,7,8,11,12,13,21`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数`instance_mode`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 支持IoT边缘服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusters`新增响应参数`vpcId`，调整响应参数`start_time`类型: `string` -> `integer`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数`unchangeable_param`、`dry_run`、`count`
  - 接口`CreateRestoreInstance`新增响应参数`enable_ssl`

# 0.0.64 2021-09-29

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复RequestHandler添加顺序错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListSubCustomerBillDetail`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSubCustomerResFeeRecords`和`ListFreeResources`
  - 接口名调整: `ListParnterAdjustRecords` -> `ListPartnerAdjustRecords`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTemp`新增响应体

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTag`的响应参数`resource_detail`类型调整： `string` -> `object`
  - 接口`CreatePrivateZone`、`UpdatePublicZone`、`DeletePublicZone`的响应参数`masters`类型调整： `string` -> `array`
  - 接口`CreatePrivateZone`和`UpdatePublicZone`的请求参数`ttl`类型调整： `string` -> `integer`
  - 接口`ListRecordSets`、`ListRecordSetsWithLine`、`ListRecordSetsByZone`的请求参数`limit`和`offset`类型调整: `string` -> `integer`
  - 接口`CreatePrivateZone`、`ListRecordSetsByZone`、`ShowRecordSetWithLine`、`ListPtrRecords`、`ListPublicZones`新增响应参数`tags`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持以下接口:
    - `ListServerTags`
    - `BatchAttachSharableVolumes`
    - `ShowServerAutoRecovery`
    - `BatchResetServersPassword`
    - `ReinstallServerWithoutCloudInit`
    - `ChangeServerOsWithoutCloudInit`
    - `BatchUpdateServersName`
    - `ShowServerPassword`
    - `AssociateServerVirtualIp`
    - `MigrateServer`
    - `ShowServerBlockDevice`
    - `DisassociateServerVirtualIp`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 调整以下接口名:
    - `ShowMysqlEngineVersion` -> `ShowGaussMySqlEngineVersion`
    - `ShowMysqlFlavors` -> `ShowGaussMySqlFlavors`
    - `CreateMysqlInstance` -> `CreateGaussMySqlInstance`
    - `ShowMysqlInstanceList` -> `ListGaussMySqlInstances`
    - `DeleteMysqlInstance` -> `DeleteGaussMySqlInstance`
    - `ShowMysqlInstanceInfo` -> `ShowGaussMySqlInstanceInfo`
    - `CreateMysqlReadonlyNode` -> `CreateGaussMySqlReadonlyNode`
    - `DeleteMysqlReadonlyNode` -> `DeleteGaussMySqlReadonlyNode`
    - `ExpandMysqlInstanceVolume` -> `ExpandGaussMySqlInstanceVolume`
    - `UpdateMysqlBackupPolicy` -> `UpdateGaussMySqlBackupPolicy`
    - `UpdateMysqlInstanceName` -> `UpdateGaussMySqlInstanceName`
    - `ResetMysqlPassword` -> `ResetGaussMySqlPassword`
    - `ChangeMysqlInstanceSpecification` -> `ChangeGaussMySqlInstanceSpecification`
    - `ListDedicatedResources`  -> `ListGaussMySqlDedicatedResources`
    - `CreateMysqlProxy` -> `CreateGaussMySqlProxy`
    - `DeleteMysqlProxy` -> `DeleteGaussMySqlProxy`
    - `ShowMysqlProxy` -> `ShowGaussMySqlProxy`
    - `ShowMysqlProxyFlavors` -> `ShowGaussMySqlProxyFlavors`
    - `ExpandMysqlProxy` -> `ExpandGaussMySqlProxy`
    - `ListMysqlErrorLog` -> `ListGaussMySqlErrorLog`
    - `ListMysqlSlowLog` -> `ListGaussMySqlSlowLog`
    - `ShowMysqlProjectQuotas` -> `ShowGaussMySqlProjectQuotas`
    - `ShowMysqlQuotas` -> `ShowGaussMySqlQuotas`
    - `SetMysqlQuotas` -> `SetGaussMySqlQuotas`
    - `UpdateMysqlQuotas` -> `UpdateGaussMySqlQuotas`
    - `CreateMysqlBackup` -> `CreateGaussMySqlBackup`
    - `ShowMysqlBackupList` -> `ShowGaussMySqlBackupList`
    - `ShowMysqlBackupPolicy` -> `ShowGaussMySqlBackupPolicy`
    - `ListMysqlConfigurations` -> `ListGaussMySqlConfigurations`
    - `ShowMysqlJobInfo` -> `ShowGaussMySqlJobInfo`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIssueRecordsV4`新增响应参数`id`和`name`
  - 接口`ListProjectIterationsV4`新增响应参数`status`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`响应参数`group_type`新增可选值`bigmen`

# 0.0.63 2021-09-26

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`BatchSetPolicy`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 支持接口`CreateEditingJob`、`ListEditingJob`、`DeleteEditingJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持下列接口：
    - `ListIgnoreRule`
    - `ListStatistics`
    - `ListQpsTimeline`
    - `ListBandwidthTimeline`
    - `ListResponseCodeTimeline`
    - `ListTopAbnormal`
    - `ShowConsoleConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.62 2021-09-24

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 修复找不到接口`ListRecordContents`的问题
- _特性变更_
  - 无

# 0.0.61 2021-09-24

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListParnterAdjustRecords`和`ListFreeResourceInfos`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSubCustomerBillDetail`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Classroom

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ApplyJudgement`新增非必填的请求参数`testcases`

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTestCaseDetailV2`的请求参数`testcase_number`改为必填，移除请求参数`testcase_id`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持云数据库GaussDB(for openGauss)服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`ListRecordContents`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SWR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRepositoryTags`新增响应参数`domain_id`、`scanned`、`tag_type`

# 0.0.60 2021-09-16

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowCluster`新增响应参数`platformVersion`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowDomainStats`
- _解决问题_
  - 修复调用接口`ShowDomainItemLocationDetails`无响应数据的问题
- _特性变更_
  - 无

### HuaweiCloud SDK DDM

- _新增特性_
  - 支持接口`ListSlowLog`和`ListReadWriteRatio`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持数据安全中心服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持全球SIM联接服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持智能边缘云服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDataImage`新增非必填的请求参数`__support_amd`
  - 接口`ListImages`新增响应参数`__support_amd`

### HuaweiCloud SDK KMS

- _新增特性_
  - 支持接口`ShowPublicKey`、`Sign`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeInvoiceVerification`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.59 2021-09-10

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListSubCustomerBillDetail`、`ListResourceUsageSummary`、`ListResourceUsage`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListResourceUsages`

### HuaweiCloud SDK CBS

- _新增特性_
  - 支持接口`CreateTbSession`、`ExecuteTbSession`、`DeleteTbSession`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`AddNode`和`ResetNode`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDomain`新增以下响应参数:
    - `range_status`
    - `follow_status`
    - `origin_status`
    - `auto_refresh_preheat`
  - 接口`UpdateDomainMultiCertificates`新增必填请求参数`switch`和可选请求参数`redirect_type`
  - 接口`UpdateHttpsInfo`新增必填请求参数`switch`和可选请求参数`redirect_type`
  - 接口`ShowHistoryTasks`新增可选请求参数`create_time`

### HuaweiCloud SDK DAS

- _新增特性_
  - 支持数据管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobDetail`新增响应参数`status`和`fail_reason`

### HuaweiCloud SDK EVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateVolume`的请求参数`size`改为必填

### HuaweiCloud SDK IVS

- _新增特性_
  - 支持人证核身服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeInvoiceVerification`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeIdCard`新增可选请求参数`return_verification`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`UpdateDatabase`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数`alias`
  - 接口`CreateDatabase`新增可选请求参数`comment`

# 0.0.58 2021-08-31

### HuaweiCloud SDK CodeCraft

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCompetitionScore`、`UpdateCompetitionScore`的请求参数`score`类型调整： `string` -> `double`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 支持云性能测试服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FRS

- _新增特性_
  - 支持以下接口:
    - `DetectLiveByUrl`
    - `DetectLiveFaceByUrl`
    - `DetectLiveByFile`
    - `DetectLiveFaceByFile`
    - `DetectLiveByBase64`
    - `DetectLiveFaceByBase64`
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
  - 接口`ListLiveStreamsOnline`新增响应参数`video_frame_rate`、`audio_frame_rate`、`audio_bitrate`、`resolution`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ClusterScaling`的请求参数`data_volume_count`、`data_volume_size`、`data_volume_type`改为非必填
  - 接口`CreateCluster`新增请求参数`auto_create_default_security_group`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持文字识别服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ImportCertificate`的请求参数`name`、`certificate`、`certificate_chain`、`private_key`改为必填
  - 接口`PushCertificate`的请求参数`target_project`、`target_service`改为必填

### HuaweiCloud SDK SMN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTopics`新增请求参数`enterprise_project_id`、`name`、`fuzzy_name`
  - 接口`ListSubscriptions`新增请求参数`protocol`、`status`、`endpoint`

# 0.0.57 2021-08-25

### HuaweiCloud SDK CBS

- _新增特性_
  - 支持对话机器人服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCraft

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCompetitionScore`和`UpdateCompetitionScore`请求参数`score`类型调整: `float` -> `string`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `ShowJobDetail`
    - `SwitchSlowlogDesensitization`
    - `ListFlavorInfos`
    - `UpdateInstanceRemark`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RestoreInstance`、`CreateManualBackup`、`RestoreInstanceFromCollection`新增响应参数`job_id`
  - 接口`ListInstances`新增响应参数`remark`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServerInterfaces`新增以下响应参数:
    - `delete_on_termination`
    - `driver_mode`
    - `min_rate`
    - `multiqueue_num`
    - `pci_address`
  - 接口`ListServersDetails`新增响应参数`cpu_options`和`hypervisor`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateSharedBandwidths`新增请求参数和响应参数`public_border_group`
  - 接口`AddPublicipsIntoSharedBandwidth`新增响应参数`public_border_group`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口名调整： `AuthorizeFaceRecognitionService` -> `ShowSubscribes`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ExportFunction`的新增请求参数`function_urn`和`type`
  - 接口`ListStatistics`的请求参数`filter`可选值修改为`monitor_data`、`monthly_report`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持以下接口
    - `ListDedicatedResources`
    - `ListFlavorInfos`
    - `ListConfigurationTemplates`
    - `ListInstancesByResourceTags`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数`dedicated_resource_id`
  - 接口`ListInstances`新增响应参数`dedicated_resource_id`

### HuaweiCloud SDK ImageSearch

- _新增特性_
  - 支持图像搜索服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`RunRecord`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTransTemplate`移除请求参数`aspect_ratio`
  - 接口`CreateTranscodingTask`移除请求参数`GOP_structure`、`sr_factor`
  - 接口`CreateTemplateGroup`移除请求参数`GOP_structure`、`sr_factor`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListJobInfoDetail`的响应参数名调整: `taskDetail` -> `task_detail`
  - 接口`ListJobInfoDetail`新增响应参数`count`

# 0.0.56 2021-08-17

### HuaweiCloud SDK AntiDDoS

- _新增特性_
  - 支持流量清洗服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNodePools`新增以下响应参数：
    - `annotations`
    - `updateTimestamp`
    - `creationTimestamp`
    - `creatingNode`
    - `deletingNode`
    - `conditions`
    - `enterprise_project_id`
  - 接口`ListClusters`新增响应参数`orderID`和`upgradefrom`

### HuaweiCloud SDK CloudTest

- _新增特性_
  - 支持云测服务
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
  - 接口`ListServerResizeFlavors`新增响应参数`ecs:instance_architecture`
  - 接口`NovaCreateServers`新增请求参数`tenancy`和`dedicated_host_id`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 修复创建LB异常的问题
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordCallbackConfig`移除请求参数`key`
  - 接口`ListRecordCallbackConfigs`新增响应参数`sign_type`
  - 接口`ShowDomain`新增响应参数`status_describe`和`service_area`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AllowSqlserverDbUserPrivilege`和`RevokeSqlserverDbUserPrivilege`新增请求参数`readonly`

### HuaweiCloud SDK SMS

- _新增特性_
  - 支持主机迁移服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.55 2021-08-10

### HuaweiCloud SDK AS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改接口`ListResourceInstances`的请求参数`key`、`value`为必填参数

### HuaweiCloud SDK CloudBuild

- _新增特性_
    - 支持编译构建服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListBandwidths`、`ShowPublicip`的响应参数名称调整：`publicip_border_group` -> `public_border_group`

### HuaweiCloud SDK EVS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListVolumes`新增请求参数`server_id`

### HuaweiCloud SDK FRS

- _新增特性_
    - 支持人脸识别服务
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
    - 接口`CreateDeployment`移除响应参数`order_id`

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改接口`UpdateImage`的请求参数`value`为必填参数

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`UpdateInstanceTopic`新增请求参数`new_partition_numbers`

### HuaweiCloud SDK LTS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListLogs`新增请求参数`highlight`
    - 接口`ListLogs`的请求参数`search_type`新增可选值`backwards`

### HuaweiCloud SDK MPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListTranscodingTask`移除响应参数`dynamic_range`
    - 接口`CreateTransTemplate`移除请求参数`tenant_id`
    - 接口`CreateThumbnailsTask`移除请求参数`percent`、`frame_type`
    - 接口`CreateTranscodingTask`移除请求参数`black_enhance`

### HuaweiCloud SDK MRS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`DescribeCluster`新增响应参数`start_time`、`state`

### HuaweiCloud SDK SWR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowAccessDomain`新增以下响应参数：
        - `namespace`
        - `repository`
        - `access_domain`
        - `permit`
        - `deadline`
        - `description`
        - `creator_id`
        - `creator_name`
        - `created`
        - `updated`
        - `status`

### HuaweiCloud SDK VPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`NeutronListSubnets`新增请求参数`enable_dhcp`
    - 接口`NeutronListSecurityGroups`新增响应参数`security_groups_links`

# 0.0.54 2021-07-27

### HuaweiCloud SDK Classroom

- _新增特性_
    - 支持接口`ApplyJudgement`、`ShowJudgementDetail`、`ShowJudgementFile`
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.53 2021-07-26

### HuaweiCloud SDK CDN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowHistoryTasks`移除响应参数`urls`、`task_id`
    - 接口`ShowHistoryTaskDetails`移除响应参数`task_id`、`process_reason`,请求参数`process_reason`类型调整： `integer`->`string`
    - 接口`ShowTopUrl`移除请求参数`user_domain_id`、`task_id`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 支持接口`ShowPlans`
- _解决问题_
    - 无
- _特性变更_
    - 无
    
### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`UpdateConfigurations`新增请求参数`dcs_cluster_proxy2_node`

### HuaweiCloud SDK DDM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`UpdateUser`移除请求参数`extend_authority`

### HuaweiCloud SDK DDS

- _新增特性_
    - 支持接口`UpdateClientNetwork`
- _解决问题_
    - 无
- _特性变更_
    - 接口`SetBalancerWindow`的请求参数`start_time`、`stop_time`改为非必填
    - 接口`CreateInstance`新增请求参数`port`，新增响应参数`port`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 支持接口`EnableLtsLogs`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowFunctionCode`新增响应参数`concurrent_num`、`id`、`encrypted_user_data`
    - 接口`ListFunctions`新增响应参数`func_vpc_id`、`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`，移除响应参数`version_description`、`last_modified_utc`、`dependencies`
    - 接口`UpdateVersionAlias`移除请求参数`name`、`last_modified`、`alias_urn`
    - 接口`ShowFunctionConfig`新增响应参数`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`，移除响应参数`version_description`、`concurrency`
    - 接口`UpdateFunctionConfig`移除请求参数`version_description`、`concurrency`、`depend_list`，新增请求参数`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`
    - 接口`ListFunctionVersions`移除响应参数`last_modified_utc`、`concurrency`，新增响应参数`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`
    - 接口`UpdateTrigger`的请求参数`size`类型调整： `string`->`integer`
    - 接口`ShowDependency`的响应参数`size`类型调整： `string`->`integer`
    - 接口`UpdateDependency`的响应参数`size`类型调整： `string`->`integer`

### HuaweiCloud SDK HSS

- _新增特性_
    - 支持接口`ListEvents`
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
    - 接口`ShowDomain`移除响应参数`domain_source`

### HuaweiCloud SDK MPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListTranscodeDetail`移除响应参数`language`
    - 接口`CreateThumbnailsTask`移除请求参数`project_id`、`tenant_project_id`、`domain_name`、`canonical_grant_id`
    - 接口`ListTranscodeDetail`新增响应参数`audit_report`
    - 接口`QueryTranscodingsTask`移除响应参数`output_url`
    - 接口`CreateTranscoding`新增请求参数`audit`，移除请求参数`special_effect`、`quality_enhance`、`template_extend`
    - 接口`ListWatermarkTemplate`移除响应参数`template_id`、`error`
    - 接口`CreateVodTranscodingTask`移除请求参数`multidrm`、`preview_duration`

### HuaweiCloud SDK VOD

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateAssetByFileUpload`的请求参数`auto_publish`类型调整： `string`->`integer`，并配置可选值为`0`、`1`

### HuaweiCloud SDK WAF

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 无
      接口`ListEvent`的响应参数`response_time`、`response_size`类型调整： `string`->`integer`

# 0.0.52 2021-07-16

### HuaweiCloud SDK AS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateScalingV2Policy`新增请求参数`description`
    - 接口`ShowScalingV2Policy`、`ShowScalingGroup`新增响应参数`description`

### HuaweiCloud SDK DCS

- _新增特性_
    - 支持更多接口：
        - `CreateDiagnosisTask`
        - `CreateRedislog`
        - `CreateRedislogDownloadLink`
        - `ListDiagnosisTasks`
        - `ListRedislog`
        - `ListSlowlog`
        - `ShowDiagnosisTaskDetails`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListInstances`新增请求参数`include_delete`

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 40](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/40): 修复响应参数`__lazyloading`类型定义错误的问题
- _特性变更_
    - 无

# 0.0.51 2021-07-09

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决了query参数名称存在大写字母的情况下鉴权失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口`ListBareMetalServers`的响应参数`addresses`数据结构定义错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListProtectable`新增响应参数`smn_notify`、`threshold`
    - 接口`AssociateVaultPolicy`新增请求参数`add_policy_ids`和响应参数`without_any_tag`、`smn_notify`、`threshold`

### HuaweiCloud SDK CCE

- _新增特性_
    - 支持接口`RemoveNode`、`MigrateNode`
- _解决问题_
    - 无
- _特性变更_
    - 接口`DeleteCluster`新增请求参数`tobedeleted`

### HuaweiCloud SDK CDN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowTopUrl`的请求参数`start_time`、`end_time`改为必填参数，`domain_name`新增可选值`outside_mainland_china`
    - 接口`ShowDomainItemDetails`新增请求参数`service_area`

### HuaweiCloud SDK DDM

- _新增特性_
    - 支持分布式数据库中间件服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DNS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreatePublicZone`的响应参数`masters`、`zones`类型调整：`string`->`array`

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateSharedBandwidth`、`ListBandwidths`新增响应参数`publicip_border_group`

### HuaweiCloud SDK GaussDB

- _新增特性_
    - 支持云数据库服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`GlanceCreateImageMetadata`新增响应参数`__root_origin`、`checksum`、`size`
    - 接口`GlanceAddImageMember`移除请求参数`deleted`、`deleted_at`,新增以下请求参数：
        - `__lazyloading`
        - `__os_feature_list`
        - `__root_origin`
        - `__sequence_num`
        - `__support_agent_list`
        - `__system__cmkid`
        - `active_at`
        - `hw_vif_multiqueue_enabled`
        - `max_ram`
        - `__image_location`
        - `__is_config_init`
        - `__account_code`

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListRules`新增响应参数`edge_node_ids`、`last_update_time`

### HuaweiCloud SDK LTS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListStructuredLogsWithTimeRange`响应参数`context`类型调整： `string`->`array`
    - 接口名称调整：
        - `UpdateLogContents`->`ListLogs`
        - `UpdateLogContents2`->`ListQueryStructuredLogs`
        - `UpdateLogContents3`->`ListStructuredLogsWithTimeRange`

### HuaweiCloud SDK MRS

- _新增特性_
    - 支持接口`ListClusters`
- _解决问题_
    - 无
- _特性变更_
    - 接口`DescribeCluster`的响应参数`clusterType`类型调整：`string`->`integer`

### HuaweiCloud SDK SWR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowRepository`新增响应参数`domain_id`、`priority`
    - 接口`CreateRetention`新增响应参数`template`

# 0.0.50 2021-06-29

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateNodePool`、`ShowNodePool`、`UpdateNodePool`、`DeleteNodePool`新增请求参数`storage`

### HuaweiCloud SDK DRS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`BatchUpdateUser`的参数`selected`类型调整： `string`->`boolean`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListInstances`的响应参数`port`类型调整： `string`->`integer`
    - 接口`ListInstances`的响应参数名称调整： `storage_engine`->`mode`
    - 接口`ListSlowLogs`移除响应参数`node_name`，新增响应参数`time`

### HuaweiCloud SDK MRS

- _新增特性_
    - 无
- _解决问题_
    - 修复部分参数定义错误问题
- _特性变更_
    - 接口`CreateCluster`移除请求参数`start_time`、`state`

### HuaweiCloud SDK NAT

- _新增特性_
    - 无
- _解决问题_
    - 修复接口`ListNatGateways`的请求参数`project_id`重复的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowInformationAboutDatabaseProxy`的响应参数`port`、`node_num`类型调整： `string`->`integer`

# 0.0.49 2021-06-25

### HuaweiCloud SDK APIG

- _新增特性_
    - 支持更多接口：
        - `ListGatewayResponsesV2`
        - `UpdateGatewayResponseV2`
        - `DeleteGatewayResponseV2`
        - `UpdateGatewayResponseTypeV2`
        - `DeleteGatewayResponseTypeV2`
        - `DeleteInstancesV2`
        - `UpdateInstanceV2`
        - `ListInstancesV2`
        - `RemoveEipV2`
        - `UpdateEngressEipV2`
        - `RemoveEngressEipV2`
        - `ListFeaturesV2`
        - `UpdateDomainV2`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BMS

- _新增特性_
    - 支持接口`ChangeBaremetalServerOs`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ChangeBaremetalServerName`的响应参数名称调整：`server_tags`->`sys_tags`

### HuaweiCloud SDK CDN

- _新增特性_
    - 支持接口`ShowQuota`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowHistoryTaskDetails`的请求参数`url`类型调整：`integer`->`string`

### HuaweiCloud SDK CloudRTC

- _新增特性_
    - 支持华为云实时音视频服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DRS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`BatchUpdateUser`参数`is_transfer`、`selected`类型调整：`string`->`boolean`

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`KeystoneListPermissions`新增请求参数`permission_type`、`display_name`、`catalog`、`type`

### HuaweiCloud SDK LTS

- _新增特性_
    - 支持云日志服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Meeting

- _新增特性_
    - 支持接口`InviteShare`
- _解决问题_
    - 无
- _特性变更_
    - 接口`SetMultiPicture`新增请求参数`multiPicSaveOnly`
    - 接口`SearchHisMeetings`新增响应参数`leftReason`

### HuaweiCloud SDK VOD

- _新增特性_
    - 支持视频点播服务
- _解决问题_
    - 无
- _特性变更_
    - 无

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
