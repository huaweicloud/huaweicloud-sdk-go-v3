# 0.0.83 2022-04-07

### HuaweiCloud SDK BSS

- _Features_
  - Support the interface `ListStoredValueCards`
- _Bug Fix_
  - None
- _Change_
  - Remove the interfaces `ListSubCustomerDiscounts`, `BatchSetSubCustomerDiscount`
  - Add the response parameters `resource_type_name`, `service_type_name` to the interface `ShowRefundOrderDetails`
  - Add the response parameter `service_type_name` to the interface `ListCustomerOrders`
  - Add the response parameters `service_type_name`, `service_type_name` to the interface `ShowCustomerOrderDetails`
  - Add the response parameters `resource_type_name`, `service_type_name` to the interface `ListPayPerUseCustomerResources`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ListCustomerOnDemandResources`
  - Add the response parameters `cloud_service_type_name`, `resource_type_name` to the interface `ListSubcustomerMonthlyBills`
  - Add the response parameters `cloud_service_type_name`, `resource_type_name`, `period_type` to the interface `ListCustomerselfResourceRecordDetails`
  - Add the response parameters `cloud_service_type_name`, `resource_type_name` to the interface `ListCustomerselfResourceRecords`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ShowCustomerMonthlySum`
  - Changes of the interface `ListCustomerBillsFeeRecords`:
    - Add the request parameters `bill_date_begin`, `bill_date_end`
    - Add the response parameters `service_type_name`, `resource_type_name`
  - Add the response parameters `resource_type_name`, `service_type_name` to the interface `ListUsageTypes`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ListSubCustomerBillDetail`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ListCustomerBillsMonthlyBreakDown`
  - Add the response parameter `service_type_name` to the interface `ListFreeResourceInfos`
  - Add the response parameter `service_type_name` to the interface `ListIncentiveDiscountPolicies`

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameters `kind`, `apiVersion`, `status` from the interface `UpdateNodePool`

### HuaweiCloud SDK DSC

- _Features_
  - Support the following interfaces：
    - `CreateDocWatermarkByAddress`
    - `ShowDocWatermarkByAddress`
    - `ShowImageWatermarkWithImage`
    - `CreateImageWatermarkByAddress`
    - `ShowImageWatermarkByAddress`
    - `ShowImageWatermarkWithImageByAddress`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateImageWatermark`:
    - Add the request parameter `image_watermark`
    - The request parameter `blind_watermark` changed to required

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `price_plan_list` from the interface `StopSimCard`
  - Remove the request parameter `price_plan_list` from the interface `ResetSimCard`

### HuaweiCloud SDK Meeting

- _Features_
  - Support the interface `StartMeeting`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeWaybillElectronic`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `print_code` to the interface `RecognizeVatInvoice`
  - Changes of the interface `RecognizeVehicleLicense`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeTaxiInvoice`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Add the response parameters `type`, `accumulated_scores`, `status`, `generation_date`, `current_time` to the interface `RecognizeDriverLicense`
  - Changes of the interface `RecognizeTrainTicket`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeBankcard`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ApplyConfigurationAsync`, `UpdateInstanceConfigurationAsync`, `DeleteSqlserverDatabaseEx`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the interfaces `ListConfigurations`, `ListDatastores`, `ListFlavors`, `ListStorageTypes`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateInstance`:
    - Add the request parameter `solution`
    - Add the enum values `centralization_standard` to the request parameter `mode`
    - Add the enum values `ESSD` to the request parameter `type`

# 0.0.82 2022-03-31

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of concurrent read/write authentication information cache
- _Change_
  - None

### HuaweiCloud SDK CC

- _Features_
  - Support the service `Cloud Connect`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - Support the interface `BatchShowNodesInformation`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.81 2022-03-25

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `responseStatus` to the interface `DeleteserviceDiscoveryRules`
  - Add the response parameter `responseStatus` to the interface `AddOrUpdateServiceDiscoveryRules`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interfaces(v2):
    - `ShowDomainLocationStats`
    - `ShowDomainStats`
    - `ShowTopUrl`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DAS

- _Features_
  - Support the interface `ShowSqlExplain`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `group_name`, `replication_ip` to the interface `ListRedislog`

### HuaweiCloud SDK DNS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `routers` from the interface `ListPublicZones`
  - Add the request parameters `marker`, `limit`, `offset`, `line_id`, `tags`, `status`, `type`, `name`, `id`, `sort_key`, `sort_dir`, `search_mode` to the interface `ShowRecordSetByZone`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - Support the following interfaces：
    - `ListWorkflows`
    - `CreateWorkflow`
    - `BatchDeleteWorkflows`
    - `ListWorkflowExecutions`
    - `StartWorkflowExecution`
    - `ShowWorkflowExecution`
    - `ShowWorkFlow`
    - `UpdateWorkFlow`
    - `ShowTenantMetric`
    - `ShowWorkFlowMetric`
    - `RetryWorkFlow`
    - `StopWorkFlow`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSimCards`:
    - Add the request parameters `min_flow`, `max_flow`, `order_id`, `filter_downtime_period`
    - Modify the type `date` -> `date-time` of the response parameter `device_status_date`
    - Modify the type `date` -> `date-time` of the response parameter `expire_time`
  - Add the request parameter `price_plan_list` to the interface `StopSimCard`
  - Add the request parameter `price_plan_list` to the interface `ResetSimCard`
  - Changes of the interface `ShowSimCard`:
    - Modify the type `date` -> `date-time` of the response parameter `device_status_date`
    - Modify the type `date` -> `date-time` of the response parameter `expire_time`

### HuaweiCloud SDK IMS

- _Features_
  - Support the interfaces `ListVersions`, `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `os_type` changed to not required of the interface `CreateDataImage`

### HuaweiCloud SDK IoTDA

- _Features_
  - Support the interface `ResetFingerprint`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RecognizeVatInvoice`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeIdCard`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeDriverLicense`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`

### HuaweiCloud SDK VSS

- _Features_
  - Support the following interfaces：
    - `ShowDomainSettings`
    - `UpdateDomainSettings`
    - `ListTaskHistories`
    - `ListPortResults`
    - `ListBusinessRisks`
    - `UpdateFalsePositive`
    - `CancelTasks`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `domain_id` to the interface `ListDomains`
  - Add the response parameter `hit_details` to the interface `ShowResults`

# 0.0.80 2022-03-10

### HuaweiCloud SDK BCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `operation_id` to the interface `BatchCreateChannels`
  - The request parameter `fabric_version` changed to not required of the interface `CreateNewBlockchain`

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `cluster_id` changed to not required of the interface `DeleteAddonInstance`

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `X-Auth-Token` from the interface `ShowTopUrl`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainLocationStats`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainItemDetails`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainStats`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainItemLocationDetails`

### HuaweiCloud SDK DWS

- _Features_
  - Support the following interfaces：
    - `ListClusterDetails`
    - `DeleteCluster`
    - `ResetPassword`
    - `ListSnapshots`
    - `CreateSnapshot`
    - `RestartCluster`
    - `ListClusters`
    - `CreateCluster`
    - `RestoreCluster`
    - `ResizeCluster`
    - `ListNodeTypes`
    - `ListSnapshotDetails`
    - `DeleteSnapshot`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ELB

- _Features_
  - Support the following interfaces：
    - `ListLogtanks`
    - `CreateLogtank`
    - `ShowLogtank`
    - `UpdateLogtank`
    - `DeleteLogtank`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `https_cps` to the interface `ListFlavors`
  - Add the response parameter `https_cps` to the interface `ShowFlavor`
  - The request parameter `X-Auth-Token` changed to not required of the interface `ListLoadBalancers`
  - The request parameter `X-Auth-Token` changed to not required of the interface `CreateLoadBalancer`
  - The request parameter `X-Auth-Token` changed to not required of the interface `UpdateIpList`
  - The request parameter `X-Auth-Token` changed to not required of the interface `BatchDeleteIpList`

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `date` -> `date-time` of the response parameter `act_date` of the interface `ListSimCards`
  - Modify the type `date` -> `date-time` of the response parameter `act_date` of the interface `ShowSimCard`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `count` to the interface `ListLogs`

### HuaweiCloud SDK Meeting

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `sortLevel` to the interface `AddDepartment`
  - Add the request parameter `sortLevel` to the interface `UpdateDepartment`
  - Add the response parameters `sortLevel`, `sortLevel` to the interface `ShowDeptAndChildDept`
  - Add the response parameter `sortLevel` to the interface `SearchDepartmentByName`

# 0.0.79 2022-03-07

### HuaweiCloud SDK Core

- _Features_
  - Support SK derived authentication
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - Support the interfaces `UpdateClusterEip`, `ShowClusterEndpoints`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CES

- _Features_
  - Support the following interfaces (V2)：
    - `ListAlarms`
    - `CreateAlarm`
    - `DeleteAlarm`
    - `UpdateAlarmAction`
    - `ListAlarmResources`
    - `DeleteAlarmResources`
    - `AddAlarmResources`
    - `AddResourceGroupsResourcesBatch`
    - `DeleteResourceGroupsResourcesBatch`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `gaussdbv5`, `postgresql` to the request parameter `engine_type` to the interface `BatchCreateJobs`
  - Changes of the interface `BatchValidateConnections`:
    - Add the request parameter `kafka_security_config`
    - Add the enum values `postgresql` to the request parameter `db_type`
  - Changes of the interface `BatchUpdateUser`:
    - Add the request parameter `is_sync_object_privilege`
    - Add the response parameters `no_privileges`, `parent_account`, `no_parent_account`
  - Add the response parameters `no_privileges`, `parent_account`, `no_parent_account` to the interface `ListUsers`
  - Add the request parameters `topic_policy`, `topic`, `partition_policy`, `kafka_data_format`, `topic_name_format`, `partitions_num`, `replication_factor`, `is_fill_materialized_view`, `export_snapshot` to the interface `BatchSetPolicy`

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `ip_version` of the interface `CreatePrePaidPublicip`: `integer` -> `enum`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `order_id` to the interface `ShrinkInstanceNode`

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `error_code` to the interface `ListEditingJob`

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `request_id` from the interface `DownloadSlowlog`

# 0.0.78 2022-02-25

### HuaweiCloud SDK AS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `alarm_id` to the interface `ListAllScalingV2Policies`
  - Add the enum values `GPSSD` to the request parameter `volume_type` to the interface `CreateScalingConfig`
  - Add the response parameter `min` to the interface `ShowResourceQuota`
  - Add the response parameter `min` to the interface `ShowPolicyAndInstanceQuota`

### HuaweiCloud SDK BMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateBaremetalServerMetadata`:
    - Modify the type of the request body `MetaData` -> `UpdateBaremetalServerMetadataReq`
    - Remove the response parameter `key`

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The response parameter `from-connector-name`, `to-link-name` changed to required of the interface `ShowJobs`
  - The request parameter `from-connector-name`, `to-link-name` changed to required of the interface `UpdateJob`
  - Changes of the interface `CreateAndStartRandomClusterJob`:
    - The request parameter `from-connector-name`, `to-link-name` changed to required
    - Modify the type `int32` -> `float` of the response parameter `progress`
    - Modify the type `boolean` -> `string` of the response parameter `isStopingIncrement`
  - Add the response parameter `submissions` to the interface `StopJob`
  - The request parameter `from-connector-name`, `to-link-name` changed to required of the interface `CreateJob`
  - Changes of the interface `StartJob`:
    - Modify the type `int32` -> `float` of the response parameter `progress`
    - Modify the type `boolean` -> `string` of the response parameter `isStopingIncrement`
  - Modify the type `double` -> `float` of the response parameter `progress` of the interface `ShowJobStatus`
  - Modify the type `double` -> `float` of the response parameter `progress` of the interface `ShowSubmissions`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interfaces `ShowDomainLocationStats`, `ShowDomainFullConfig`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowDomainStats`:
    - Add the request parameter `service_area`
    - Remove the request parameters `X-Auth-Token`, `country`, `district`, `isp`
    - Remove the response parameters `start_time`, `end_time`, `interval`, `action`, `stat_type`, `group_by`
  - Add the request parameter `https` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CloudIDE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `instance_domain_id`, `instance_user_id` to the interface `CreateInstance`
  - The request parameter `instance_user_domain_name`, `instance_user_name` changed to not required of the interface `CreateInstanceBy3rd`

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interface `CheckRecord`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudRTC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `array` -> `string` of the request parameter `mid` of the interface `ListRtcClientQosDetails`

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the interface `ShowStatisticCommitV3`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListProjectSets`:
    - Add the response parameters `CreateTime`, `UpdateTime`, `external_params`, `variables_no_file`
    - Remove the response parameters `create_time`, `update_time`, `group`
  - The request parameter `name` changed to required of the interface `UpdateProject`
  - Changes of the interface `ShowTask`:
    - Add the response parameters `parallel`, `contents`, `sort`, `related_temp_running_data`
    - Remove the response parameter `content`
  - Changes of the interface `UpdateTask`:
    - Add the request parameters `contents`, `sort`, `related_temp_running_data`
    - Add the response parameters `parallel`, `contents`, `sort`, `related_temp_running_data`
    - Remove the request parameter `content`
    - Remove the response parameter `content`
    - The request parameter `name` changed to required
  - Changes of the interface `ShowReport`:
    - Add the response parameters `performance`, `minNetworkTraffic`, `avgNetworkTraffic`, `maxNetworkTraffic`, `branchId`, `branchName`, `projectId`, `serviceId`
    - Remove the response parameters `progressState`, `statusValue`
    - Modify the type `float` -> `double` of the response parameter `averageRespTime`
    - Modify the type `float` -> `double` of the response parameter `avgRecBytes`
    - Modify the type `int32` -> `double` of the response parameter `avgSentBytes`
    - Modify the type `string` -> `double` of the response parameter `avgTranRespTime`
    - Modify the type `int32` -> `double` of the response parameter `currentThreadNum`
    - Modify the type `int32` -> `double` of the response parameter `errorCount`
    - Modify the type `int32` -> `double` of the response parameter `errorEventsCount`
    - Modify the type `int32` -> `double` of the response parameter `failedAssert`
    - Modify the type `int32` -> `double` of the response parameter `failedOthers`
    - Modify the type `int32` -> `double` of the response parameter `failedParsed`
    - Modify the type `int32` -> `double` of the response parameter `failedRefused`
    - Modify the type `int32` -> `double` of the response parameter `failedTimeout`
    - Modify the type `int32` -> `double` of the response parameter `max`
    - Modify the type `int32` -> `double` of the response parameter `maxRecBytes`
    - Modify the type `int32` -> `double` of the response parameter `maxRespTime`
    - Modify the type `int32` -> `double` of the response parameter `maxSentBytes`
    - Modify the type `int32` -> `double` of the response parameter `maxTranRespTime`
    - Modify the type `int32` -> `double` of the response parameter `min`
    - Modify the type `int32` -> `double` of the response parameter `requests`
    - Modify the type `int32` -> `double` of the response parameter `result`
    - Modify the type `int32` -> `double` of the response parameter `status`
    - Modify the type `int32` -> `double` of the response parameter `successCount`
    - Modify the type `int32` -> `double` of the response parameter `successRate`
    - Modify the type `int32` -> `double` of the response parameter `sum1xx`
    - Modify the type `int32` -> `double` of the response parameter `sum2xx`
    - Modify the type `int32` -> `double` of the response parameter `sum3xx`
    - Modify the type `int32` -> `double` of the response parameter `sum4xx`
    - Modify the type `int32` -> `double` of the response parameter `sum5xx`
    - Modify the type `int32` -> `double` of the response parameter `taskStatus`
    - Modify the type `int32` -> `double` of the response parameter `tp50`
    - Modify the type `int32` -> `double` of the response parameter `tp75`
    - Modify the type `int32` -> `double` of the response parameter `tp90`
    - Modify the type `int32` -> `double` of the response parameter `tp95`
    - Modify the type `int32` -> `double` of the response parameter `tp99`
    - Modify the type `float` -> `double` of the response parameter `tps`
    - Modify the type `string` -> `double` of the response parameter `tranTPS`
    - Modify the type `string` -> `double` of the response parameter `transactionSuccess`
    - Modify the type `int32` -> `double` of the response parameter `transactionalSuccessRate`
    - Modify the type `int32` -> `double` of the response parameter `transactionalTps`
    - Modify the type `int32` -> `double` of the response parameter `transactionalTpsSuccess`
    - Modify the type `string` -> `double` of the response parameter `transactions`
    - Modify the type `int32` -> `double` of the response parameter `vum`
    - Modify the type `float` -> `double` of the response parameter `avgResponseTime`
    - Modify the type `int32` -> `double` of the response parameter `caseRetry`
    - Modify the type `int32` -> `double` of the response parameter `completeNum`
    - Modify the type `int32` -> `double` of the response parameter `duration`
    - Modify the type `int32` -> `double` of the response parameter `executedNum`
    - Modify the type `int32` -> `double` of the response parameter `kpiCaseCount`
    - Modify the type `int32` -> `double` of the response parameter `kpiCaseExecuteCount`
    - Modify the type `int32` -> `double` of the response parameter `kpiCasePassCount`
    - Modify the type `int32` -> `double` of the response parameter `maxUsers`
    - Modify the type `int32` -> `double` of the response parameter `passNum`
    - Modify the type `int32` -> `double` of the response parameter `stage`
    - Modify the type `int32` -> `double` of the response parameter `totalNum`
  - Changes of the interface `UpdateCase`:
    - Add the request parameters `contents`, `sort`
    - Remove the request parameter `content`
  - Add the request parameter `contents` to the interface `CreateTemp`
  - Changes of the interface `UpdateTemp`:
    - Modify the type `array` -> `string` of the request parameter `bodys`
    - The request parameter `name` changed to required
  - Add the request parameter `is_quoted` to the interface `CreateVariable`
  - Changes of the interface `ShowHistoryRunInfo`:
    - Modify the type `int32` -> `double` of the response parameter `run_id`
    - Modify the type `int32` -> `double` of the response parameter `run_type`
    - Modify the type `int32` -> `double` of the response parameter `continue_time`

### HuaweiCloud SDK CSS

- _Features_
  - Support the following interfaces：
    - `UpdateFlavor`
    - `UpdateFlavorByType`
    - `UpdateShrinkNodes`
    - `UpdateShrinkCluster`
    - `ListLogsJob`
    - `ShowClusterDetail`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `id`, `instances` from the interface `UpdateExtendCluster`
  - Remove the request parameter `status` from the interface `StartConnectivityTest`

### HuaweiCloud SDK DDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `group_id` to the interface `ExpandInstanceNodes`
  - Add the request parameter `group_id` to the interface `ShrinkInstanceNodes`
  - The request parameter `shard_unit` changed to not required of the interface `CreateDatabase`

### HuaweiCloud SDK DGC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListConnections`:
    - Add the response parameters `type`, `description`
    - Remove the response parameter `connectionType`
    - Modify the type `string` -> `int32` of the response parameter `total`
    - The response parameter `name` changed to required
  - Changes of the interface `CreateConnection`:
    - Add the request parameters `type`, `description`
    - Remove the request parameter `connectionType`
    - The request parameter `name` changed to required
  - Changes of the interface `ShowConnection`:
    - Add the response parameters `type`, `description`
    - Remove the response parameter `connectionType`
    - The response parameter `name` changed to required
  - Changes of the interface `UpdateConnection`:
    - Add the request parameters `type`, `description`
    - Remove the request parameter `connectionType`
    - The request parameter `name` changed to required
  - Changes of the interface `ExecuteScript`:
    - Add the response parameter `instanceId`
    - Remove the response parameter `jobId`
    - Modify the type `string` -> `object` of the request parameter `params`

### HuaweiCloud SDK ELB

- _Features_
  - Support the interfaces `BatchCreateMembers`, `BatchDeleteMembers`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ListFunctions`
  - Changes of the interface `CreateFunction`:
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the request parameter `runtime`
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ShowFunctionCode`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `UpdateFunctionCode`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ShowFunctionConfig`
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the request parameter `runtime`
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ListFunctionVersions`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `CreateFunctionVersion`
  - Add the enum values `Java11`, `Node.js14.18`, `Python3.9` to the request parameter `runtime` to the interface `CreateDependency`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `PHP 7.3` from the request parameter `runtime` to the interface `UpdateDependency`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ImportFunction`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interfaces `UpdateAuditLog`, `ShowAuditLog`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `ListSingleStreamDetail`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK LTS

- _Features_
  - Support the interfaces `UpdateStructConfig`, `CreateStructConfig`, `ListStructTemplate`, `ListBreifStructTemplate`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `tag` to the interface `ListLogGroups`
  - Add the response parameter `tag` to the interface `ListLogStream`

### HuaweiCloud SDK Meeting

- _Features_
  - Support the service `Meeting`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interfaces `BatchUpdateChildNickNames`, `ListIterationHistories`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `updated_time_interval`, `include_deleted` to the interface `ListProjectIterationsV4`
  - Add the request parameters `include_deleted`, `updated_time_interval` to the interface `ListIssuesV4`
  - Add the response parameters `description`, `order`, `accessories` to the interface `ShowIssueV4`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ListSlowLogFile`, `StopInstance`, `StartupInstance`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SCM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `sans`, `signature_algorithm`, `deploy_support` to the interface `ListCertificates`
  - Add the request parameters `enc_certificate`, `enc_private_key` to the interface `ImportCertificate`
  - Changes of the interface `ShowCertificate`:
    - Add the response parameter `signature_algorithm`
    - Remove the response parameter `signature_algrithm`
  - Add the response parameters `enc_certificate`, `enc_private_key` to the interface `ExportCertificate`

### HuaweiCloud SDK VOD

- _Features_
  - Support the interface `ListDomainLogs`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `delete_type` to the interface `DeleteAssets`

### HuaweiCloud SDK VPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `subnetpool_id` to the interface `NeutronListSubnets`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListWhiteblackipRule`:
    - Add the response parameter `addr`
    - Remove the response parameter `ip`

# 3.0.77 2022-02-10

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateAlarmRule`:
    - Modify the type `string` -> `enum` of the request parameter `statistic`
    - The request parameter `alarm_level`, `comparison_operator`, `evaluation_periods`, `metric_name`, `namespace`, `period`, `statistic`, `threshold`, `unit` changed to not required

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListRateOnPeriodDetail`:
    - Add the request parameter `fee_installment_mode`
    - Add the response parameters `installment_official_website_amount`、`installment_period_type`、`installment_official_discount_amount`、`installment_amount`

### HuaweiCloud SDK CBR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListProtectable`:
    - Modify the type `string` -> `boolean` of the response parameter `result`
    - Modify the type `string` -> `int32` of the response parameter `size`
  - Changes of the interface `ShowProtectable`:
    - Modify the type `string` -> `boolean` of the response parameter `result`
    - Modify the type `string` -> `int32` of the response parameter `size`

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `kind`, `apiVersion`, `metadata`, `spec`, `status` from the interface `CreateAddonInstance`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `ListNodes`
  - Add the request parameter `isStatic` to the interface `CreateNode`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `DeleteNode`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `ShowNode`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `UpdateNode`
  - Changes of the interface `RemoveNode`:
    - The request parameter `uid` changed to required
    - The response parameter `uid` changed to required
  - Changes of the interface `MigrateNode`:
    - The request parameter `uid` changed to required
    - The response parameter `uid` changed to required
  - Add the response parameter `isStatic` to the interface `ListNodePools`
  - Add the request parameter `isStatic` to the interface `CreateNodePool`
  - Add the response parameter `isStatic` to the interface `DeleteNodePool`
  - Add the response parameter `isStatic` to the interface `ShowNodePool`
  - Changes of the interface `UpdateNodePool`:
    - Add the request parameter `isStatic`
    - Add the response parameter `isStatic`

### HuaweiCloud SDK CSS

- _Features_
  - Support the following interfaces：
    - `UpdateOndemandClusterToPeriod`
    - `UpdateClusterName`
    - `ResetPassword`
    - `StartKibanaPublic`
    - `UpdateCloseKibana`
    - `UpdateAlterKibana`
    - `UpdatePublicKibanaWhitelist`
    - `StopPublicKibanaWhitelist`
    - `CreateCnf`
    - `UpdateCnf`
    - `StartPipeline`
    - `StopPipeline`
    - `AddFavorite`
    - `StartConnectivityTest`
    - `ListTemplates`
    - `ListConfs`
    - `ListPipelines`
    - `ListActions`
    - `ShowGetConfDetail`
    - `DeleteConf`
    - `DeleteTemplate`
    - `StartLogs`
    - `StopLogs`
    - `ShowGetLogSetting`
    - `UpdateLogSetting`
    - `StartLogAutoBackupPolicy`
    - `StopLogAutoBackupPolicy`
    - `CreateLogBackup`
    - `ShowLogBackup`
    - `CreateBindPublic`
    - `UpdateUnbindPublic`
    - `UpdatePublicBandWidth`
    - `StartPublicWhitelist`
    - `StopPublicWhitelist`
    - `StartVpecp`
    - `StopVpecp`
    - `ShowVpcepConnection`
    - `UpdateVpcepConnection`
    - `UpdateVpcepWhitelist`
    - `UpdateYmls`
    - `ListYmlsJob`
    - `ListYmls`
    - `ListClustersDetails`
    - `CreateCluster`
    - `DeleteCluster`
    - `RestartCluster`
    - `UpdateExtendCluster`
    - `UpdateExtendInstanceStorage`
    - `ListFlavors`
    - `ShowClusterTag`
    - `CreateClustersTags`
    - `ListClustersTags`
    - `UpdateBatchClustersTags`
    - `DeleteClustersTags`
    - `ShowIkThesaurus`
    - `CreateLoadIkThesaurus`
    - `DeleteIkThesaurus`
    - `StartAutoSetting`
    - `UpdateSnapshotSetting`
    - `ShowAutoCreatePolicy`
    - `CreateAutoCreatePolicy`
    - `CreateSnapshot`
    - `ListSnapshots`
    - `StopSnapshot`
    - `RestoreSnapshot`
    - `DeleteSnapshot`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - Support the interfaces `CreateOnlineMigrationTask`, `SetOnlineMigrationTask`, `BatchStopMigrationTasks`, `StopMigrationTaskSync`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DevStar

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowApplicationDependentResources`:
    - Add the request parameters `X-Auth-Token`, `limit`, `offset`
    - Add the response parameter `count`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateFunction`:
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the request parameter `runtime`
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ListFunctions`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `UpdateFunctionCode`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ShowFunctionCode`
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the request parameter `runtime`
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ShowFunctionConfig`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `CreateFunctionVersion`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ListFunctionVersions`
  - Add the enum values `Go1.x` to the request parameter `runtime` to the interface `CreateDependency`
  - Add the enum values `Go1.x` to the request parameter `runtime` to the interface `UpdateDependency`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ImportFunction`

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `alias` to the interface `ShowGaussMySqlInstanceInfo`
  - Add the response parameter `job_id` to the interface `CreateGaussMySqlBackup`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int32` -> `string` of the response parameter `port` of the interface `ListInstances`

### HuaweiCloud SDK Live

- _Features_
  - Support the interfaces `ListTranscodeTaskCount`, `ListAreaDetail`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `publish_domain` to the interface `ListRecordData`
  - Remove the request parameter `plan_record_time` from the interface `CreateRecordRule`
  - Remove the response parameter `plan_record_time` from the interface `ListRecordRules`
  - Remove the response parameter `plan_record_time` from the interface `ShowRecordRule`
  - Changes of the interface `UpdateRecordRule`:
    - Remove the request parameter `plan_record_time`
    - Remove the response parameter `plan_record_time`
  - Remove the request parameter `on_demand_callback_url` from the interface `CreateRecordCallbackConfig`
  - Remove the response parameter `on_demand_callback_url` from the interface `ListRecordCallbackConfigs`
  - Remove the response parameter `on_demand_callback_url` from the interface `ShowRecordCallbackConfig`
  - Remove the request parameter `on_demand_callback_url` from the interface `UpdateRecordCallbackConfig`
  
### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `restore_all_database` to the interface `RestoreToExistingInstance`
  - Remove the request parameter `is_open_recycle_policy` from the interface `StartRecyclePolicy`

### HuaweiCloud SDK SIS

- _Features_
  - Support the `Voice Interaction Service`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.76 2022-01-25

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `content_type` to the interface `CreateApiV2`
  - Add the response parameters `publish_time`, `roma_app_name`, `ld_api_id`, `api_group_info`, `content_type` to the interface `ShowDetailsOfApiV2`
  - Changes of the interface `UpdateApiV2`:
    - Add the request parameter `content_type`
    - Add the response parameters `publish_time`, `roma_app_name`, `ld_api_id`, `api_group_info`, `content_type`
  - Add the response parameter `content_type` to the interface `ListApiRuntimeDefinitionV2`
  - Changes of the interface `ListApiVersionDetailV2`:
    - Add the response parameters `roma_app_name`, `ld_api_id`, `api_group_info`, `content_type`
    - Modify the type `date-time` -> `string` of the response parameter `publish_time`

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowJobs`:
    - Remove the response parameter `simple`
    - The response parameter `name`, `values` changed to required
  - The request parameter `name`, `values` changed to required of the interface `UpdateJob`
  - The request parameter `name`, `values` changed to required of the interface `CreateAndStartRandomClusterJob`
  - The request parameter `name`, `values` changed to required of the interface `CreateJob`
  - Remove the response parameter `config_status` from the interface `ListClusters`

### HuaweiCloud SDK CES

- _Features_
  - Support the interface `ListAlarmHistories`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `DeleteRuleset`, `SetDefaulTemplate`, `ShowTasklog`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `endpoint_id` to the interface `CreateTask`
  - Add the request parameter `custom_attributes` to the interface `CreateRuleset`
  - Changes of the interface `ListTemplateRules`:
    - Add the request parameter `tags`
    - Add the response parameter `rule_config_list`

### HuaweiCloud SDK DevStar

- _Features_
  - Support the interfaces `ShowRepositoryByCloudIde`, `ListTemplates`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IAM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `session_user_id` to the interface `CreateLoginToken`

### HuaweiCloud SDK Kafka

- _Features_
  - Support the interface `ListEngineProducts`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `dr_enable` to the interface `ListInstances`
  - Add the response parameter `dr_enable` to the interface `ShowInstance`
  - Changes of the interface `ListProducts`:
    - Add the response parameters `Hourly`, `Monthly`
    - Remove the response parameters `hourly`, `honthly`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `language` from the interface `CreateSqlAlarmRule`
  - Remove the request parameter `language` from the interface `UpdateSqlAlarmRule`
  - Changes of the interface `ListSqlAlarmRules`:
    - Add the response parameters `template_name`, `status`
    - Remove the response parameter `language`
  - Remove the request parameters `language`, `eps_id`, `eps_name`, `is_time_range_relative` from the interface `CreateKeywordsAlarmRule`
  - Changes of the interface `UpdateKeywordsAlarmRule`:
    - Remove the request parameters `language`, `eps_id`, `eps_name`, `is_time_range_relative`
    - Remove the response parameters `eps_id`, `eps_name`, `is_time_range_relative`
  - Changes of the interface `ListKeywordsAlarmRules`:
    - Add the response parameters `template_name`, `status`
    - Remove the response parameters `language`, `eps_id`, `eps_name`, `is_time_range_relative`
  - Changes of the interface `ListAccessConfig`:
    - Add the request parameter `access_config_tag_list`
    - Add the response parameter `access_config_tag`
  - Changes of the interface `CreateAccessConfig`:
    - Add the request parameter `access_config_tag`
    - Add the response parameter `access_config_tag`
  - Changes of the interface `UpdateAccessConfig`:
    - Add the request parameter `access_config_tag`
    - Add the response parameter `access_config_tag`
  - Add the response parameter `access_config_tag` to the interface `DeleteAccessConfig`

### HuaweiCloud SDK RabbitMQ

- _Features_
  - Support the interface `ListEngineProducts`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.75 2022-01-17

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the request body structure of the interface `CreateNodePool` is incorrect
- _Change_
  - None

### HuaweiCloud SDK FunctionGraph

- _Features_
  - Support the interface `ListFunctionAsyncInvocations`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `used` to the interface `ShowGaussMySqlInstanceInfo`
  - Add the request parameter `monitor_switch` to the interface `UpdateInstanceMonitor`
  - Modify the type of the request parameter `period` of the interface `UpdateInstanceMonitor`: `string` -> `int32`
  - Remove the response parameter `mode` from the interface `ShowGaussMySqlProjectQuotas`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `ListApiVersionNew`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `az_desc` to the interface `ListFlavors`

### HuaweiCloud SDK ROMA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `id`, `roles` to the interface `AddUserToApp`
  - Remove the response parameters `apps`, `tasks` from the interface `DownloadAssetArchive`
  - Remove the request parameters `apps`, `tasks` from the interface `ImportAsset`
  - The request parameter `tasks` changed to required of the interface `DeleteAsset`
  - Add the response parameter `access_user` to the interface `ShowMqsInstance`

# 0.0.74 2022-01-10

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `coupon_max_use_quantity`, `coupon_group` to the interface `ListOrderCouponsByOrderId`

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowQuotas`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `CreateCluster`
  - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `ListClusters`
  - Changes of the interface `UpdateCluster`:
    - Add the request parameters `customSan`, `containerNetwork`
    - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain`
  - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `ShowCluster`
  - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `DeleteCluster`
  - Add the request parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `CreateNode`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `ListNodes`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `UpdateNode`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `ShowNode`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `DeleteNode`
  - Add the request parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `CreateNodePool`
  - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `ListNodePools`
  - Changes of the interface `UpdateNodePool`:
    - Add the request parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode`
    - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode`
  - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `ShowNodePool`
  - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `DeleteNodePool`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interface `UpdateDomainFullConfig`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudIDE

- _Features_
  - Support the interface `ListStacks`
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListStacksByTag`

### HuaweiCloud SDK Cloudtest

- _Features_
  - Support the interface `ShowPlanList`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowPlans`:
    - Modify the type `int64` -> `int32` of the request parameter `limit`
    - Modify the type `int64` -> `int32` of the request parameter `offset`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ecs_tenant_private_ip` to the interface `ListMigrationTask`
  - Add the response parameter `ecs_tenant_private_ip` to the interface `ShowMigrationTask`
  - Add the response parameter `ecs_tenant_private_ip` to the interface `StopMigrationTask`

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `lb_ip_address` from the interface `ListInstances`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `delete_on_termination` to the interface `CreateServers`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the request parameter `is_stateful_function`
    - Add the response parameter `is_stateful_function`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `lb_ip_address` to the interface `ListInstances`

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `GaussDB(for openGauss)`, Remove the enum values `GaussDB(openGauss)` from the request parameter `type` to the interface `CreateInstance`
  - Remove the response parameter `related_instance` from the interface `ListInstances`

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The response parameter `active_at` changed to not required of the interface `GlanceListImages`
  - The response parameter `active_at` changed to not required of the interface `GlanceShowImage`
  - The response parameter `active_at` changed to not required of the interface `GlanceUpdateImage`

### HuaweiCloud SDK IoTAnalytics

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `site_id` changed to required of the interface `CreateDatasource`
  - The request parameter `site_id` changed to required of the interface `UpdateDataSource`

### HuaweiCloud SDK KPS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListKeypairs`:
    - Add the request parameters `limit`, `marker`
    - Add the response parameter `page_info`
  - Add the request parameters `limit`, `offset` to the interface `ListFailedTask`
  - Add the request parameters `limit`, `offset` to the interface `ListRunningTask`

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `StartInstanceSingleToHaAction`:
    - Add the request parameter `ad_domain_info`
    - Remove the request parameter `password`

### HuaweiCloud SDK ROMA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `connect_dn` from the interface `ShowMqsInstance`
  - Remove the response parameter `policies` from the interface `ListMqsInstanceTopics`
  - Add the response parameter `roma_app_type` to the interface `ShowDetailsOfAppV2`

### HuaweiCloud SDK VPCEP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `pool_id` to the interface `CreateEndpointService`
  - Add the response parameter `domain_id` to the interface `ListEndpointService`
  - Add the response parameter `pool_id` to the interface `UpdateEndpointService`
  - Add the response parameters `enable_status`, `specification_name` to the interface `ListEndpointInfoDetails`

### HuaweiCloud SDK VSS

- _Features_
  - Support the interfaces `AuthorizeDomains`, `ShowTasks`, `CreateTasks`, `ShowResults`
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `skip` to the response parameter `auth_status` to the interface `ListDomains`

# 0.0.73 2021-12-25

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `id` to the interface `CreateLink`
  - Add the response parameters `endpointDomainName`, `isScheduleBootOff`, `failedReasons`, `components`, `createFrom`, `resourceId`, `flavorType`, `workSpaceId`, `trial` to the interface `ShowClusterDetail`
  - Add the request parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type` to the interface `UpdateJob`
  - Add the response parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type` to the interface `ShowJobs`
  - Changes of the interface `CreateAndStartRandomClusterJob`:
    - Add the request parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type`
    - Add the response parameter `submissions`
    - Remove the response parameters `name`, `validation-result`
  - Add the request parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type` to the interface `CreateJob`
  - Add the response parameter `execute-date` to the interface `StartJob`
  - Add the request parameter `id` to the interface `UpdateLink`
  - Add the response parameter `id` to the interface `ShowLink`
  - Changes of the interface `ListClusters`:
    - Add the response parameters `bakExpectedStartTime`, `bakKeepDay`, `createFrom`, `resourceId`, `flavorType`, `workSpaceId`, `trial`, `components`
    - Remove the response parameter `version`

### HuaweiCloud SDK CloudBuild

- _Features_
  - Support the interface `ShowHistoryDetails`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudTable

- _Features_
  - Support the service `CloudTable`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Cloudtest

- _Features_
  - Support the interfaces `ShowPlanJournals`, `ShowIssuesByPlanId`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `CheckParameters`, `ListTaskParameter`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `rule_desc` to the interface `ListRules`
  - Add the response parameters `is_devcloud_project_default`, `is_default_template` to the interface `ListRulesets`

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the service `CodeHub`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `backup_format` to the interface `CopyInstance`

### HuaweiCloud SDK DevStar

- _Features_
  - Support the following interfaces：
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
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `dependents` to the interface `ShowTemplateV3`
  - Add the response parameters `dependents`, `dependent_services` to the interface `ListTemplatesV2`
  - Add the response parameter `show_type` to the interface `ShowJobDetail`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `updated_at` from the interface `NovaListServerActions`

### HuaweiCloud SDK IEF

- _Features_
  - Support the service `Intelligent EdgeFabric`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IoTAnalytics

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateDatasource`:
    - Add the request parameter `site_id`
    - Remove the request parameter `instance_id`
  - Changes of the interface `ShowAllDataSource`:
    - Add the response parameter `site_id`
    - Remove the response parameter `instance_id`
  - Changes of the interface `UpdateDataSource`:
    - Add the request parameter `site_id`
    - Add the response parameter `site_id`
    - Remove the request parameter `instance_id`
    - Remove the response parameter `instance_id`
  - Changes of the interface `ShowDataSource`:
    - Add the response parameter `site_id`
    - Remove the response parameter `instance_id`

### HuaweiCloud SDK Kafka

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `partition_num` changed to not required of the interface `CreatePostPaidInstance`
  - Add the response parameters `result`, `instance_id` to the interface `RestartManager`
  - Changes of the interface `ListProducts`:
    - Add the response parameters `hourly`, `honthly`
    - Remove the response parameters `Hourly`, `Monthly`

### HuaweiCloud SDK KPS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int64` -> `string` of the response parameter `task_time` of the interface `ListFailedTask`
  - Modify the type `int64` -> `string` of the response parameter `task_time` of the interface `ListRunningTask`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `offset`, `limit` to the interface `ListCharts`
  - Add the request parameters `offset`, `limit` to the interface `ListNotificationTemplates`

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateTranscodingTask`:
    - Add the request parameter `auto_volume_setting`
    - Add the enum values `original` to the request parameter `volume`
  - Add the response parameter `av_parameters` to the interface `ListTranscodingTask`
  - Add the request parameter `template_id` to the interface `CreateWatermarkTemplate`
  - Add the response parameter `template_id` to the interface `ListWatermarkTemplate`
  - Add the request parameter `template_id` to the interface `UpdateWatermarkTemplate`

### HuaweiCloud SDK MRS

- _Features_
  - Support the following interfaces：
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
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `group_type`, `success_record_error_reason`, `skip_record_error_reason`, `save_prefix` to the interface `ListTasks`
  - Add the response parameters `group_type`, `success_record_error_reason`, `skip_record_error_reason`, `save_prefix` to the interface `ShowTask`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ListApiVersion`, `ShowApiVersion`, `SearchQueryScaleComputeFlavors`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ROMA

- _Features_
  - Support the service `ROMA Connect`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SA

- _Features_
  - Support the service `Situation Awareness`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.72 2021-12-17

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudIDE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowInstance`:
    - Add the response parameters `bundle_url`, `visitor_id`, `visitor_name`, `visitor_domain_name`
    - Remove the response parameters `action_list`, `role`, `role_id`, `sub_org`
  - Changes of the interface `ListOrgInstances`:
    - Add the response parameters `visitor_id`, `visitor_name`, `visitor_domain_name`
    - Remove the response parameters `action_list`, `role`, `role_id`, `sub_org`
  - Changes of the interface `ListInstances`:
    - Add the response parameters `visitor_id`, `visitor_name`, `visitor_domain_name`
    - Remove the response parameters `action_list`, `role`, `role_id`, `sub_org`

### HuaweiCloud SDK CloudRTC

- _Features_
  - Support the interfaces `ListRtcAbnormalEvents`, `ListRtcAbnormalEventDimension`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListEvents`:
    - Modify the type `string` -> `int32` of the response parameter `event_count`
    - Modify the type `string` -> `int64` of the response parameter `latest_occur_time`
  - Modify the type `string` -> `double` of the response parameter `variance` of the interface `BatchListMetricData`
  - Modify the type `string` -> `int32` of the response parameter `type_statistics` of the interface `ListResourceGroup`
  - Changes of the interface `ListEventDetail`:
    - Modify the type `string` -> `array` of the response parameter `event_users`
    - Modify the type `string` -> `array` of the response parameter `event_sources`

### HuaweiCloud SDK IoTAnalytics

- _Features_
  - Support the service `IoTAnalytics`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `input`, `edit_settings` to the interface `CreateEditingJob`
  - Add the response parameters `input`, `edit_settings` to the interface `ListEditingJob`

### HuaweiCloud SDK OCR

- _Features_
  - Support the following interfaces：
    - `RecognizeThailandIdcard`
    - `RecognizeMyanmarIdcard`
    - `RecognizeMyanmarDriverLicense`
    - `RecognizeChileIdCard`
    - `RecognizeThailandLicensePlate`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.71 2021-12-10

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Update the enum values to `[AND, OR, NOT]` of the request parameter `relation` of the interface `ListEvents`

### HuaweiCloud SDK APM

- _Features_
  - Support the interfaces `ShowMasterAddress`, `ListAkSk`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK AS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `allowed_address_pairs` to the interface `ListScalingGroups`
  - Add the request parameter `allowed_address_pairs` to the interface `CreateScalingGroup`
  - Add the request parameter `allowed_address_pairs` to the interface `UpdateScalingGroup`
  - Add the response parameter `allowed_address_pairs` to the interface `ShowScalingGroup`

### HuaweiCloud SDK BCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int32` -> `int64` of the response parameter `node_cnt` of the interface `ShowBlockchainDetail`
  - Add the request parameters `limit`, `offset` to the interface `ShowBlockchainFlavors`

### HuaweiCloud SDK CampusGo

- _Features_
  - Support the service `CampusGo`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CGS

- _Features_
  - Support the `Container Guard Service`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudIDE

- _Features_
  - Support the interfaces `ShowExtensionAuthorization`, `CreateExtensionAuthorization`, `CheckInstanceAccess`, `UpdateInstanceActivity`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudRTC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `stream_id` to the interface `ListRtcClientQosDetails`

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `ShowTaskCmetrics`, `ListTemplateRules`, `ListTaskRuleset`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `is_access`, `trigger_type` to the interface `ShowTaskDetail`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `readonly_domain_name` to the interface `ListInstances`
  - Add the response parameter `readonly_domain_name` to the interface `ShowInstance`

### HuaweiCloud SDK DDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `time_zone` to the interface `CreateInstance`
  - Add the response parameters `instanceId`, `instanceName`, `jobId` to the interface `RestartInstance`

### HuaweiCloud SDK DSC

- _Features_
  - Support the interface `ShowOpenApiCalledRecords`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `readonly_password` to the interface `CreateDocWatermark`
  - Changes of the interface `ShowScanJobs`:
    - Add the request parameter `offset`
    - Remove the request parameter `page`
  - Changes of the interface `ShowScanJobResults`:
    - Add the request parameter `offset`
    - Remove the request parameter `page`

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `landmark`, `gender`, `yaw_angle`, `roll_angle`, `pitch_angle`, `headpose`, `smile`, `skin`, `ethnic` from the interface `DetectFaceByFile`
  - Remove the response parameters `landmark`, `gender`, `yaw_angle`, `roll_angle`, `pitch_angle`, `headpose`, `smile`, `skin`, `ethnic` from the interface `DetectFaceByUrl`
  - Remove the response parameters `landmark`, `gender`, `yaw_angle`, `roll_angle`, `pitch_angle`, `headpose`, `smile`, `skin`, `ethnic` from the interface `DetectFaceByBase64`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the following interfaces：
    - `ListInstanceTags`
    - `ListProjectTags`
    - `BatchTagAction`
    - `ShowInstanceMonitorExtend`
    - `UpdateInstanceMonitor`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListGaussMySqlInstances`:
    - Add the request parameters `private_ip`, `tags`
    - Add the response parameter `tags`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `region` changed to required of the interface `ListFlavors`

### HuaweiCloud SDK GES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `CheckMetadata`

### HuaweiCloud SDK HiLens

- _Features_
  - Support the service `Huawei HiLens`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK KMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `encrypted_privatekey` to the interface `ImportKeyMaterial`

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `default_record_config` changed to required of the interface `CreateRecordRule`

### HuaweiCloud SDK LTS

- _Features_
  - Support the following interfaces：
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
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `RestoreExistInstance`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SWR

- _Features_
  - Support the interface `ListQuotas`
- _Bug Fix_
  - None
- _Change_
  - Modify the name `UpdateNamespaceAuthReq` -> `UpdateNamespaceAuthRequestBody` of the request body of the interface `UpdateNamespaceAuth`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `to`, `from` changed to required of the interface `ListStatistics`
  - The request parameter `from`, `to` changed to required of the interface `ListQpsTimeline`
  - The request parameter `from`, `to` changed to required of the interface `ListBandwidthTimeline`
  - The request parameter `from`, `to` changed to required of the interface `ListTopAbnormal`
  - Add the response parameter `cookie` to the interface `ShowEvent`
  - Changes of the interface `CreatePremiumHost`:
    - Add the request parameter `Content-Type`
    - Modify the type `string` -> `enum` of the request parameter `type`
  - Add the request parameter `Content-Type` to the interface `ListPremiumHost`
  - Changes of the interface `UpdatePremiumHost`:
    - Add the request parameter `Content-Type`
    - Modify the type `string` -> `enum` of the response parameter `type`
  - Add the request parameter `Content-Type` to the interface `DeletePremiumHost`
  - Changes of the interface `ShowPremiumHost`:
    - Add the request parameter `Content-Type`
    - Modify the type `string` -> `enum` of the response parameter `type`
  - Add the request parameter `Content-Type` to the interface `UpdatePremiumHostProtectStatus`

# 0.0.70 2021-11-29

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `roma_app_type` to the interface `ResettingAppSecretV2`.

### HuaweiCloud SDK BCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameter `values` of the interface `ListEntityMetric`: `object` -> `array`

### HuaweiCloud SDK CBR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `provider_id` to the interface `ListBackups`.

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `cidrs` to the interface `ShowCluster`.

### HuaweiCloud SDK DBSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the response parameter of the interfaces `SwitchAgent` and `SwitchRiskRule`: `status` -> `result`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `port` to the interface `UpdateInstance`.

### HuaweiCloud SDK DSC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `start_time` to the interface `ShowScanJobs`.

### HuaweiCloud SDK EVS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `CinderExportToImage`.

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `os:scheduler_hints` to the interfaces `NovaShowServer` and `NovaListServersDetails`.

### HuaweiCloud SDK ELB

- _Features_
  - Support the following interfaces:
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
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ShowQuotaDefaults`.
  - Add the response parameters `flavor_sold_out` and `lcu` to the interfaces `ListFlavors` and `ShowFlavor`, and remove the response parameter `availability_zone_ids`.
  - Add the response parameter `members_per_pool` to the interface `ShowQuota`.
  - Add the request parameters `enc_certificate` and `enc_private_key` to the interfaces `CreateCertificate` and `UpdateCertificate`.
  - Add the response parameters `enc_certificate` and `enc_private_key` to the interfaces `ListCertificates` and `ShowCertificate`.
  - Add the request parameters `prepaid_options`, `autoscaling` and `id` to the interface `CreateLoadBalancer`.
  - Add the request parameters `elb_virsubnet_type` and `autoscaling` and the response parameters `autoscaling` and `ip_version` to the interface `ListLoadBalancers`.
  - Add the request parameters `elb_virsubnet_type` and `autoscaling` and the response parameters `loadbalancer_id`, `order_id`, `autoscaling` and `ip_version` to the interface `UpdateLoadBalancer`.
  - Add the response parameters `autoscaling` and `ip_version` to the interface `ShowLoadBalancer`.
  - Add the response parameters `id`, `type` and `provisioning_status` to the interface `ShowLoadBalancerStatus`.
  - Add the request parameters `security_policy_id` and `enhance_l7policy_enable` to the interface `CreateListener`.
  - Add the request parameters `enhance_l7policy_enable` and `member_instance_id` and the response parameters `security_policy_id`, `transparent_client_ip_enable` and `enhance_l7policy_enable` to the interface `ListListeners`.
  - Add the request parameters `enhance_l7policy_enable` and `member_instance_id` and the response parameters `security_policy_id`, `transparent_client_ip_enable` and `enhance_l7policy_enable` to the interface `UpdateListener`.
  - Add the response parameters `security_policy_id`, `transparent_client_ip_enable` and `enhance_l7policy_enable` to the interface `ShowListener`.
  - Add the request parameters `listener_id` and `member_instance_id` to the interface `ListPools`.
  - Add the request parameters `ip_version` and `member_type` and the response parameters `member_type` and `instance_id` to the interface `ListMembers`.
  - Add the response parameters `member_type` and `instance_id` to the interfaces `UpdateMember`, `ShowMember` and `ListAllMembers`.
  - Add the request parameters `priority`, `redirect_url_config`, `fixed_response_config` and `conditions` to the interface `CreateL7Policy`.
  - Add the request parameter `priority` and the response parameters `redirect_url_config` and `fixed_response_config` to the interface `ListL7Policies`.
  - Add the request parameters `priority`, `redirect_url_config`, `fixed_response_config` and `rules` and the response parameters `member_type` and `instance_id` to the interface `UpdateL7Policy`.
  - Add the response parameters `redirect_url_config` and `fixed_response_config` to the interface `ShowL7Policy`.
  - Add the request parameter `conditions` to the interface `CreateL7Rule`.
  - Add the response parameter `conditions` to the interfaces `ListL7Rules` and `ShowL7Rule`.
  - Add the request and response parameter `conditions` to the interface `UpdateL7Rule`.

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ListVersionAliases` to the interface `ListVersionAliases`.
  - Update the request parameter `name` to required of the interfaces  `CreateDependency` and `UpdateDependency`.
  - Update the request parameters `name` and `content` to required of the interface `CreateEvent`.

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `enable_force_switch` to the interface `CreateInstance`.

### HuaweiCloud SDK GES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `graphSizeTypeIndex` of the interface `CreateGraph`: `integer` -> `string`

### HuaweiCloud SDK KMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Update the request parameter `key_alias` to required of the interface `CreateKey`.
  - Update the request parameter `key_id` to required of the interfaces `EnableKey`, `CancelKeyDeletion`, `ListKeys`, `ListKeyDetail`, `ShowPublicKey`, `ListGrants`, `DeleteImportedKeyMaterial`, `EnableKeyRotation`, `DisableKeyRotation` and `ShowKeyRotationStatus`.
  - Update the request parameters `key_id` and `pending_days` to required of the interface `DeleteKey`.
  - Add the request parameter `enterprise_project_id` to the interface `ListKeys`.
  - Update the request parameter `random_data_length` to required of the interface `CreateRandom`.
  - Update the request parameters `key_id` and `datakey_length` to required of the interfaces `CreateDatakey` and `CreateDatakeyWithoutPlaintext`.
  - Update the request parameters `key_id`, `plain_text` and `datakey_plain_length` to required of the interface `EncryptDatakey`.
  - Update the request parameters `key_id`, `cipher_text` and `datakey_cipher_length` to required of the interface `DecryptDatakey`.
  - Update the request parameters `key_id` and `key_alias` to required of the interface `UpdateKeyAlias`.
  - Update the request parameters `key_id` and `key_description` to required of the interface `UpdateKeyDescription`.
  - Update the request parameters `key_id`, `grantee_principal` and `operations` to required of the interface `CreateGrant`.
  - Update the request parameters `key_id` and `grant_id` to required of the interfaces `CancelGrant` and `CancelSelfGrant`.
  - Update the request parameters `key_id` and `plain_text` to required of the interface `EncryptData`.
  - Update the request parameter `cipher_text` to required of the interface `DecryptData`.
  - Update the request parameters `key_id` and `wrapping_algorithm` to required of the interface `CreateParametersForImport`.
  - Update the request parameters `key_id`, `import_token` and `encrypted_key_material` to required of the interface `ImportKeyMaterial`.
  - Update the request parameters `key_id` and `rotation_interval` to required of the interface `UpdateKeyRotationInterval`.
  - Add the request parameter `sequence` to the interface `CreateKmsTag`.

### HuaweiCloud SDK RDS

- _Features_
  - Support the following interfaces:
    - `ChangeProxyScale`
    - `SearchQueryScaleFlavors`
    - `ShowInformationAboutDatabaseProxy`
    - `StartDatabaseProxy`
    - `StopDatabaseProxy`
    - `UpdateReadWeight`
    - `ChangeTheDelayThreshold`
    - `ShowDrReplicaStatus`
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameter `size` of the interface `ListPostgresqlDatabases`: `int32` -> `int64`

### HuaweiCloud SDK SMS

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the enumeration value contains the Chinese description and causes the parameter error.
- _Change_
  - None

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `items` to the interface `ListIgnoreRule`.
  - Add the request parameter `attacks` to the interface `ListEvent`.
  - Add the response parameter `host_id` to the interface `ShowEvent`.
  - Add the response parameter `certificatename` to the interface `UpdateHost`.

# 0.0.69 2021-11-25

### HuaweiCloud SDK AOM

- _Features_
  - Support the service `Application Operations Management`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `CheckBackendConnectivity`.

### HuaweiCloud SDK APM

- _Features_
  - Support the service `Application Performance Management`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BCS

- _Features_
  - Support the service `Application Performance Management`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `statistic_type` and response parameter `bill_date` to the interface `ListCustomerselfResourceRecordDetails`.

### HuaweiCloud SDK CBH

- _Features_
  - Support the service `Cloud Bastion Host`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `platformVersion` to the interface `ShowCluster`.
  - Add the enumeration values `RollingBack` and `RollbackFailed` to the request parameter `status` of the interface `ListClusters`.

### HuaweiCloud SDK CloudRTC

- _Features_
  - Support the following interfaces:
    - `ListRtcRealtimeScaleDimension`
    - `ListRtcRealtimeQuality`
    - `ListRtcRealtimeNetwork`
    - `ListRtcHistoryUsage`
    - `ListRtcHistoryScale`
    - `ListRtcHistoryQuality`
    - `ListRtcClientQosDetails`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support CodeCheck service.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DAS

- _Features_
  - None
- _Bug Fix_
  - Correct the enumeration value of the request parameter `X-Language` of some interfaces.
- _Change_
  - None

### HuaweiCloud SDK DBSS

- _Features_
  - Support the `Database Security Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `tags` to the interfaces `CreateInstance` and `ListInstances`.

### HuaweiCloud SDK DeH

- _Features_
  - Support the service `Dedicated Host`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request and response parameter `alias` to the interfaces `CreatePrePaidPublicip` and `CreatePublicip`.
  - Add the response parameter `alias` to the interfaces `ShowPublicip` and `UpdatePublicip`.

### HuaweiCloud SDK GES

- _Features_
  - Support the interfaces `ResizeGraph`, `ExpandGraph` and `UploadFromObs`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK LTS

- _Features_
  - Support the following interfaces:
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
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK KMS

- _Features_
  - Support the interfaces of version `V2`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK NLP

- _Features_
  - Support the service `Natural Language Processing`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `extracted_data` of the interface `RecognizeHandwriting`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `sort` to the interface `ListSlowlogStatistics`.

### HuaweiCloud SDK SFSTurbo

- _Features_
  - Support the service `Scalable File Service - Turbo`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK TMS

- _Features_
  - Support the interface `ShowTagQuota`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VPCEP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameters `created_at` and `updated_at` of the interfaces `ListServicePublicDetails`, `ListServiceDetails`, `ListServiceConnections`, `AcceptOrRejectEndpoint`, `ListEndpoints`, `UpdateEndpointWhite` and `ListEndpointInfoDetails`.
  - Update the request parameters `vpc_id` and `port_id` to required of the interface `CreateEndpointService`.
  - Remove the response parameter `error` of the interface `AcceptOrRejectEndpoint`.
  - Modify the type of the response parameter `updated` of the interfaces `ListVersionDetails` and `ListSpecifiedVersionDetails`: `datetime` -> `string`
  - Update the request parameter `action` to required of the interfaces `ListResourceInstances` and `BatchAddOrRemoveResourceInstance`.

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `from` and `to` to the interface `ListEvent`.
  - Add the request parameter `name` to the interface `ListWhiteblackipRule`.

# 0.0.68 2021-11-12

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of parsing errors when the response body type is `array` or `map`.
  - Fix the problem of automatic escaping of request body.
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `customSan` to the interface `UpdateNode`.

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `force_redirect_https` to the interface `UpdateHttpsInfo`.

### HuaweiCloud SDK CloudDeploy

- _Features_
  - Support the service `CloudDeploy`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the response parameter of the interface `ListAuditLogs`: `total_count` -> `total_record`.

### HuaweiCloud SDK DSC

- _Features_
  - Support the interfaces `ShowScanJobs` and `ShowScanJobResults`.
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `marked_file_password` to the interface `CreateDocWatermark`.

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `offset` and `limit`.
  - Add the response parameters `id` and `spec_code`.

### HuaweiCloud SDK GES

- _Features_
  - Support `Graph Engine Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameters `sim_card_id`, `partner`, `package_type` and `sim_type` from the interface `ListProPricePlans`.
  - Remove the request parameter `tag_id` from the interface `ListSimCards`.
  - Remove the request parameter `sim_price_plan_id` from the interface `ListSimPricePlans`.
  - Remove the request parameter `price_plan_list` from the interfaces `StopSimCard` and `ResetSimCard`.
  - Remove the response parameters `sn`, `supply_code`, `bundle_id` and `test_type` from the interfaces `ShowSimCard` and `ListSimCards`.
  - Remove the response parameter `sim_price_plan_list` from the interface `StopSimCard`.
  - Remove the response parameter `order_id` from the interface `ListSimPools`.
  - Remove the response parameters `partner` and `partner_pid` from the interface `ListSimPricePlans`.

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `results` to the interface `ShowJob`.

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `country` and `protocol` to the interface `ListBandwidthDetail` and `ListUsersOfStream`.
  - Add the request parameter `protocol` to the interface `ListDomainTrafficDetail`, `ListDomainBandwidthPeak` and `ListDomainTrafficSummary`.
  - Add the request parameter `stream` to the interface `ListTranscodeData`.
  - Add the request parameters `stream`, `start_time` and `end_time` to the interface `ListHistoryStreams`.

### HuaweiCloud SDK LTS

- _Features_
  - Support the following interfaces:
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
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the request parameter of the interface `CreateEditingJob`: `const` -> `consts`
  - Remove the request parameter `index`, and add the request parameters `overlay_input`, `const` and `mix` to the interface `CreateEditingJob`.
  - Add the response parameter `output` to the interface `ListEditingJob`.
  - Add the response parameters `hls_init_count` and `hls_init_interval`, add the optional value `EFFICIENCY` to the request parameter `video_enhance` of the interface `CreateTranscodingTask`.

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `job_id`, `user` and `queue` to the interface(V2) `GetJobExeListNew`.

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `confidence` to the interface `RecognizeGeneralTable`.

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interface `CreateSystemIssueV4`.
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `sequence` to the interface `ListIrs`.
  - Add the request parameter `status_id` to the interface `BatchUpdateIrs`.

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ListSlowLogsNew`, `ListErrorLogsNew` and `UpgradeDbVersion`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VAS

- _Features_
  - Support `Video Analysis Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VPC

- _Features_
  - Support interfaces(V3): `AddVpcExtendCidr`、`RemoveVpcExtendCidr`、`ListVpcs`、`ShowVpc`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.67 2021-10-25

### HuaweiCloud SDK BCS

- _Features_
  - Support the interface `ShowBlockchainFlavors`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `account_manager_id` and `account_manager_name` to the interface `ListIndirectPartners`.

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `create_time` and response parameter `task_type` to the interface `ShowHistoryTaskDetails`.
  - Remove the request parameter `create_time` from the interface `ShowHistoryTasks`.

### HuaweiCloud SDK DNS

- _Features_
  - Support the interface `ShowDomainQuota`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `bandwidth_type` to the interface `CreateSharedBandwidth`.

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `single` to the interfaces `AddFacesByFile`, `AddFacesByBase64` and `AddFacesByUrl`.

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of request and response parameters `num` and `size` of the interface `CreateInstance`: `integer` -> `string`.

### HuaweiCloud SDK GSL

- _Features_
  - Support the service `Global SIM Link`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ImageSearch

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the request parameter of the interface `RunSearchPicture`: `isCrop` -> `is_crop`
  - Add the request parameter `box` to the interface `RunSearchPicture`.

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `current_task`, `image_name` and `process_percent` to the interface `ShowJob`.

### HuaweiCloud SDK IoTDA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `status` to the interface `ListDevices`.
  - Add the request parameter `file_path` to the interface `CreateRuleAction`.

### HuaweiCloud SDK OCR

- _Features_
  - Support the interfaces `RecognizeInsurancePolicy`, `RecognizeFinancialStatement` and `RecognizeQualificationCertificate`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `database_name` to the interface `ShowBackupDownloadLink`.
  - Add the response parameter `max_iops` and `expiration_time` to the interface `ListInstances`.

### HuaweiCloud SDK SDRS

- _Features_
  - Support the `Storage Disaster Recovery Service`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.66 2021-10-19

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `backup_id` to the interface `ShowMigrationTask`.
  - Add the following response parameters to the interface `ListMigrationTask`:
    - `source_instance_name`
    - `source_instance_id`
    - `target_instance_addrs`
    - `target_instance_id`
  - Modify the type of the response parameter `total_usec_sum` of the interface `ListDiagnosisTasks`: `integer` -> `double`

### HuaweiCloud SDK EIP

- _Features_
  - Support the following interfaces:
    - `ListCommonPools`
    - `ListPublicBorderGroups`
    - `ListPublicipPool`
    - `ShowPublicipPool`
    - `ListShareBandwidthTypes`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `allow_share_bandwidth_type_any` to the interface `ListPublicips`.
  - Modify the type of the request parameter `limit` of the interface `NeutronListFloatingIps`.
  - Modify the name of the request body of the interface `NeutronUpdateFloatingIp`: `floatingip` -> `NeutronUpdateFloatingIpRequestBody`
  - Add the response parameter `allow_share_bandwidth_types` to the interface `ShowPublicip`.

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `hls_init_count` and `hls_init_interval` to the interface `CreateTranscodingTask`.

### HuaweiCloud SDK VPCEP

- _Features_
  - Support the service `VPC Endpoint`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.65 2021-10-11

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the issue that there is no status code when the response body is a stream.
- _Change_
  - None

### HuaweiCloud SDK APIG

- _Features_
  - Support the interfaces:
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
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListSkuInventories`.

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `ShowTask`:
    - `create_time`
    - `description`
    - `operate_mode`
    - `related_temp_running_data`
    - `run_status`
    - `update_time`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameter `reverse_binding` to `optional` of the interface `DisassociateServerVirtualIp`.

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the optional value to `2,4,6,7,8,11,12,13,21` of the request parameter `attributes` of the interfaces `DetectFaceByFile`,`DetectFaceByBase64` and `DetectFaceByUrl`.

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `instance_mode` to the interface `ListInstances`.

### HuaweiCloud SDK IoTEdge

- _Features_
  - Support the service `IoTEdge`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `vpcId` to the interface `ListClusters`, and modify the type of the response parameter `start_time` of this interface: `string` -> `integer`

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `unchangeable_param`,`dry_run` and `count` to the interface `CreateInstance`.
  - Add the response parameter `enable_ssl` to the interface `CreateRestoreInstance`.

# 0.0.64 2021-09-29

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the issue that RequestHandler is added in the wrong order.
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - Support the interface `ListSubCustomerBillDetail`.
- _Bug Fix_
  - None
- _Change_
  - Remove the interfaces `ListSubCustomerResFeeRecords` and `ListFreeResources`.
  - Modify the name of the interface: `ListParnterAdjustRecords` -> `ListPartnerAdjustRecords`

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add request body to the interface `CreateTemp`.

### HuaweiCloud SDK DNS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameter `resource_detail` of the interface `ListTag`: `string` -> `object`
  - Modify the type of the response parameter `masters` of the interfaces `CreatePrivateZone`,`UpdatePublicZone` and `DeletePublicZone`: `string` -> `array`
  - Modify the type of the request parameter `ttl` of the interfaces `CreatePrivateZone` and `UpdatePublicZone`: `string` -> `integer`
  - Modify the type of the request parameters `limit` and `offset` of the interfaces `ListRecordSets`,`ListRecordSetsWithLine` and `ListRecordSetsByZone`: `string` -> `integer`
  - Add the response parameter `tags` to the interfaces `CreatePrivateZone`,`ListRecordSetsByZone`,`ShowRecordSetWithLine`,`ListPtrRecords` and `ListPublicZones`.

### HuaweiCloud SDK ECS

- _Features_
  - Support the following interfaces:
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
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the following interfaces:
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

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `id` and `name` to the interface `ListIssueRecordsV4`.
  - Add the response parameter `status` to the interface `ListProjectIterationsV4`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the optional value `bigmen` to the response parameter `group_type` of the interface `ListFlavors`.

# 0.0.63 2021-09-26

### HuaweiCloud SDK DRS

- _Features_
  - Support the interface `BatchSetPolicy`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - Support the interfaces `CreateEditingJob`,`ListEditingJob` and `DeleteEditingJob`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK WAF

- _Features_
  - Support the following interfaces:
    - `ListIgnoreRule`
    - `ListStatistics`
    - `ListQpsTimeline`
    - `ListBandwidthTimeline`
    - `ListResponseCodeTimeline`
    - `ListTopAbnormal`
    - `ShowConsoleConfig`
- _Bug Fix_
  - None
- _Change_
  - None

# 3.0.62 2021-09-24

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - Fix the issue that the interface `ListRecordContents` cannot be found.
- _Change_
  - None

# 0.0.61 2021-09-24

### HuaweiCloud SDK BSS

- _Features_
  - Support the interfaces `ListParnterAdjustRecords` and `ListFreeResourceInfos`.
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListSubCustomerBillDetail`.

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowQuotas`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Classroom

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the optional request parameter `testcases` to the interface `ApplyJudgement`.

### HuaweiCloud SDK Cloudtest

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameter `testcase_number` to `optional`, and remove the request parameter of the interface `ShowTestCaseDetailV2`.

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the service `GaussDB(for openGauss)`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `ListRecordContents`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SWR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `domain_id`, `scanned` and `tag_type` to the interface `ListRepositoryTags`.

# 0.0.60 2021-09-16

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `platformVersion` to the interface `ShowCluster`.

### HuaweiCloud SDK CDN

- _Features_
  - Support the interface `ShowDomainStats`.
- _Bug Fix_
  - Fix the issue of no response data when calling the interface `ShowDomainItemLocationDetails`.
- _Change_
  - None

### HuaweiCloud SDK DDM

- _Features_
  - Support the interfaces `ListSlowLog` and `ListReadWriteRatio`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DSC

- _Features_
  - Support the service `Data Security Center`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - Support the service `Global SIM Link`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IEC

- _Features_
  - Support the service `Intelligent EdgeCloud`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the optional request parameter `__support_amd` to the interface `CreateDataImage`.
  - Add the response parameter `__support_amd` to the interface `ListImages`.

### HuaweiCloud SDK KMS

- _Features_
  - Support the interfaces `ShowPublicKey` and `Sign`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeInvoiceVerification`.
- _Bug Fix_
  - None
- _Change_
  - None

# 3.0.59 2021-09-10

### HuaweiCloud SDK BSS

- _Features_
  - Support the interfaces `ListSubCustomerBillDetail`, `ListResourceUsageSummary` and `ListResourceUsage`.
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListResourceUsages`.

### HuaweiCloud SDK CBS

- _Features_
  - Support the interfaces `CreateTbSession`, `ExecuteTbSession` and `DeleteTbSession`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - Support the interfaces `AddNode` and `ResetNode`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `CreateDomain`:
    - `range_status`
    - `follow_status`
    - `origin_status`
    - `auto_refresh_preheat`
  - Add the required request parameter `switch` and optional request parameter `redirect_type` to the interface `UpdateDomainMultiCertificates`.
  - Add the required request parameter `switch` and optional request parameter `redirect_type` to the interface `UpdateHttpsInfo`.
  - Add the optional request parameter `create_time` to the interface `ShowHistoryTasks`.

### HuaweiCloud SDK DAS

- _Features_
  - Support the `Data Admin Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `status` and `fail_reason` to the interface `ShowJobDetail`.

### HuaweiCloud SDK EVS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameter `size` of the interface `CreateVolume` to `required`.

### HuaweiCloud SDK IVS

- _Features_
  - Support the service `Identity Verification Solution`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeInvoiceVerification`.
- _Bug Fix_
  - None
- _Change_
  - Add the optional request parameter `return_verification` to the interface `RecognizeIdCard`.

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `UpdateDatabase`.
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `alias` to the interface `ListInstances`.
  - Add the optional request parameter `comment` to the interface `CreateDatabase`.

# 3.0.58 2021-08-31

### HuaweiCloud SDK CodeCraft

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `score` of the interfaces `CreateCompetitionScore` and `UpdateCompetitionScore`: `string` -> `double`.

### HuaweiCloud SDK CPTS

- _Features_
  - Support the service `CPTS`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK FRS

- _Features_
  - Support the following interfaces:
    - `DetectLiveByUrl`
    - `DetectLiveFaceByUrl`
    - `DetectLiveByFile`
    - `DetectLiveFaceByFile`
    - `DetectLiveByBase64`
    - `DetectLiveFaceByBase64`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `video_frame_rate`,`audio_frame_rate`,`audio_bitrate` and `resolution`.

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameters `data_volume_count`, `data_volume_size` and `data_volume_type` to `optional` of the interface `ClusterScaling`.
  - Add the request parameter `auto_create_default_security_group` to the interface `CreateCluster`.

### HuaweiCloud SDK OCR

- _Features_
  - Support the service `Optical Character Recognition`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SCM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameters `name`, `certificate`, `certificate_chain` and `private_key` to `required` of the interface `ImportCertificate`.
  - Set the request parameters `target_project` and `target_service` to `required` of the interface `PushCertificate`.

### HuaweiCloud SDK SMN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `enterprise_project_id`, `name` and `fuzzy_name` to the interface `ListTopics`.
  - Add the request parameters `protocol`, `status` and `endpoint` to the interface `ListSubscriptions`.

# 0.0.57 2021-08-25

### HuaweiCloud SDK CBS

- _Features_
  - Support the `Conversational Bot Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCraft

- _Features_
  - Support the `Conversational Bot Service`.
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `score` of the interfaces `CreateCompetitionScore` and `UpdateCompetitionScore`: `float` -> `string`

### HuaweiCloud SDK DDS

- _Features_
  - Support the following interfaces:
    - `ShowJobDetail`
    - `SwitchSlowlogDesensitization`
    - `ListFlavorInfos`
    - `UpdateInstanceRemark`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `job_id` to the interfaces `RestoreInstance`, `CreateManualBackup` and `RestoreInstanceFromCollection`.
  - Add the response parameter `remark` to the interface `ListInstances`.

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `ListServerInterfaces`:
    - `delete_on_termination`
    - `driver_mode`
    - `min_rate`
    - `multiqueue_num`
    - `pci_address`
  - Add the response parameters `cpu_options` and `hypervisor` to the interface `ListServersDetails`.

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request and response parameter `public_border_group` to the interface `BatchCreateSharedBandwidths`.
  - Add the response parameter `public_border_group` to the interface `AddPublicipsIntoSharedBandwidth`.

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the interface: `AuthorizeFaceRecognitionService` -> `ShowSubscribes`.

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `function_urn` and `type` to the interface `ExportFunction`.
  - -The optional value of the request parameter `filter` of the interface `ListStatistics` is modified to [`monitor_data`, `monthly_report`]

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - Support the following interfaces:
    - `ListDedicatedResources`
    - `ListFlavorInfos`
    - `ListConfigurationTemplates`
    - `ListInstancesByResourceTags`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `dedicated_resource_id` to the interface `CreateInstance`.
  - Add the response parameter `dedicated_resource_id` to the interface `ListInstances`.

### HuaweiCloud SDK ImageSearch

- _Features_
  - Support the service `ImageSearch`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `RunRecord`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `aspect_ratio` of the interface `CreateTransTemplate`.
  - Remove the request parameter `GOP_structure` and `sr_factor` of the interface `CreateTranscodingTask`.
  - Remove the request parameter `GOP_structure` and `sr_factor` of the interface `CreateTemplateGroup`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the response parameter of the interface `ListJobInfoDetail`: `taskDetail` -> `task_detail`
  - Add the response parameter `count` to the interface `ListJobInfoDetail`.

# 0.0.56 2021-08-17

### HuaweiCloud SDK AntiDDoS

- _Features_
  - Support the service `Anti-DDoS`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `ListNodePools`:
    - `annotations`
    - `updateTimestamp`
    - `creationTimestamp`
    - `creatingNode`
    - `deletingNode`
    - `conditions`
    - `enterprise_project_id`
  - Add the response parameters `orderID` and `upgradefrom` to the interface `ListClusters`.

### HuaweiCloud SDK CloudTest

- _Features_
  - Support the service `CloudTest`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ecs:instance_architecture` to the interface `ListServerResizeFlavors`.
  - Add the request parameters `tenancy` and `dedicated_host_id` to the interface `NovaCreateServers`.

### HuaweiCloud SDK ELB

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of abnormal creation of LB.
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `key` from the interface `CreateRecordCallbackConfig`.
  - Add the response parameter `sign_type` to the interface `ListRecordCallbackConfigs`.
  - Add the response parameters `status_describe` and `service_area` to the interface `ShowDomain`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `readonly` to the interfaces `AllowSqlserverDbUserPrivilege` and `RevokeSqlserverDbUserPrivilege`.

### HuaweiCloud SDK SMS

- _Features_
  - Support `Server Migration Service`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.55 2021-08-10

### HuaweiCloud SDK AS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the request parameters `key` and `value` of the interface `ListResourceInstances` as required parameters.

### HuaweiCloud SDK CloudBuild

- _Features_
    - Support the service `CloudBuild`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of the response parameter of the interfaces `ListBandwidths` and `ShowPublicip`: `publicip_border_group` -> `public_border_group`

### HuaweiCloud SDK EVS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `server_id` to the interface `ListVolumes`.

### HuaweiCloud SDK FRS

- _Features_
    - Support the `Face Recognition Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `order_id` from the interface `CreateDeployment`.

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the request parameter `value` of the interface `UpdateImage` as a required parameter.

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `new_partition_numbers` to the interface `UpdateInstanceTopic`.

### HuaweiCloud SDK LTS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `highlight` to the interface `ListLogs`.
    - Add the optional value `backwards` to the request parameter `search_type` of the interface `ListLogs`.

### HuaweiCloud SDK MPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `dynamic_range` from the interface `ListTranscodingTask`.
    - Remove the request parameter `tenant_id` from the interface `CreateTransTemplate`.
    - Remove the request parameters `percent` and `frame_type` from the interface `CreateThumbnailsTask`.
    - Remove the request parameter `black_enhance` from the interface `CreateTranscodingTask`.

### HuaweiCloud SDK MRS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `start_time` and `state` to the interface `DescribeCluster`.
    
### HuaweiCloud SDK SWR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the following response parameters to the interface `ShowAccessDomain`:
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

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `enable_dhcp` to the interface `NeutronListSubnets`.
    - Add the response parameter `security_groups_links` to the interface `NeutronListSecurityGroups`.

# 0.0.54 2021-07-27

### HuaweiCloud SDK Classroom

- _Features_
    - Support the interfaces `ApplyJudgement`,`ShowJudgementDetail`,`ShowJudgementFile`.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.53 2021-07-26

### HuaweiCloud SDK CDN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameters `urls`, `task_id` of the interface `ShowHistoryTasks`.
    - Remove the response parameters `task_id`, `process_reason`, modify the type of the request parameter `process_reason`：`integer`->`string`
    - Remove the request parameters `user_domain_id`, `task_id` of the interface `ShowTopUrl`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - Support the interface `ShowPlans`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `dcs_cluster_proxy2_node` to the interface `UpdateConfigurations`.

### HuaweiCloud SDK DDM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the request parameter `extend_authority` of the interface `UpdateUser`.

### HuaweiCloud SDK DDS

- _Features_
    - Support the interface `UpdateClientNetwork`.
- _Bug Fix_
    - None
- _Change_
    - Change the request parameters `start_time`,`stop_time` from `optional` to `required` of the interface `SetBalancerWindow`.
    - Add the request parameter `port` and response parameter `port` to the interface `CreateInstance`.

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support the interface `EnableLtsLogs`.
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `concurrent_num`,`id`,`encrypted_user_data`.
    - Add the response parameters `func_vpc_id`,`encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` to the interface `ListFunctions`, and remove the response parameters `version_description`,`last_modified_utc`,`dependencies` of this interface.
    - Remove the request parameter `name`,`last_modified`,`alias_urn` of the interface `UpdateVersionAlias`.
    - Add the response parameters `encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` of the interface `ShowFunctionConfig`, and remove the response parameters `version_description`,`concurrency` of this interface.
    - Add the response parameters `encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` to the interface `ListFunctionVersions`, and remove the response parameters `version_description`,`concurrency`,`depend_list`.
    - Add the response parameters `encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` to the interface `ListFunctionVersions`, remove the response parameters `last_modified_utc`,`concurrency`.
    - Modify the type of the request parameter `size` of the interface `UpdateTrigger`: `string`->`integer`
    - Modify the type of the response parameter `size` of the interface `ShowDependency`: `string`->`integer`
    - Modify the type of the response parameter `size` of the interface `UpdateDependency`: `string`->`integer`

### HuaweiCloud SDK HSS

- _Features_
    - Support the interface `ListEvents`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `domain_source` of the interface `ShowDomain`.

### HuaweiCloud SDK MPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `language` of the interface `ListTranscodeDetail`.
    - Remove the request parameter `project_id`,`tenant_project_id`,`domain_name`,`canonical_grant_id` of the interface `CreateThumbnailsTask`.
    - Add the response parameter `audit_report` to the interface `ListTranscodeDetail`.
    - Remove the response parameter `output_url` of the interface `QueryTranscodingsTask`.
    - Add the request parameter `audit` to the interface, and remove the request parameter `special_effect`,`quality_enhance`,`template_extend` of this interface.
    - Remove the response parameter `template_id`,`error` of the interface `ListWatermarkTemplate`.
    - Remove the request parameter `multidrm`,`preview_duration` of the interface `CreateVodTranscodingTask`.

### HuaweiCloud SDK VOD

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the request parameter `auto_publish` of the interface `CreateAssetByFileUpload`, and configure the optional values `0`,`1`.

### HuaweiCloud SDK WAF

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameters `response_time`,`response_size` of the interface `ListEvent`: `string`->`integer`.

# 0.0.52 2021-07-16

### HuaweiCloud SDK AS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `description` to the interface `CreateScalingV2Policy`.
    - Add the response parameter `description` to the interfaces `ShowScalingV2Policy`, `ShowScalingGroup`.

### HuaweiCloud SDK DCS

- _Features_
    - Support more interfaces:
        - `CreateDiagnosisTask`
        - `CreateRedislog`
        - `CreateRedislogDownloadLink`
        - `ListDiagnosisTasks`
        - `ListRedislog`
        - `ListSlowlog`
        - `ShowDiagnosisTaskDetails`
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `include_delete` to the interface `ListInstances`.

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - [Issue 40](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/40): Fix the issue that the type of the response parameter `__lazyloading` is incorrectly defined.
- _Change_
    - None

# 0.0.51 2021-07-09

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the issue of authentication failure when the query parameter name contains uppercase letters.
- _Change_
    - None

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - Fix the issue that the data structure of the response parameter `addresses` of the interface `ListBareMetalServers` is incorrectly defined.
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `smn_notify`,`threshold` to the interface `ListProtectable`.
    - Add the request parameter `add_policy_ids` and the response parameters `without_any_tag`,`smn_notify`,`threshold` to the interface `AssociateVaultPolicy`.

### HuaweiCloud SDK CCE

- _Features_
    - Support the interfaces `RemoveNode`,`MigrateNode`.
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `tobedeleted` to the interface `DeleteCluster`.

### HuaweiCloud SDK CDN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change the request parameters `start_time`,`end_time` from optional to required of the interface `ShowTopUrl`, and add the optional value `outside_mainland_china` to the request parameter `domain_name`.
    - Add the request parameter `service_area` to the interface `ShowDomainItemDetails`.

### HuaweiCloud SDK DDM

- _Features_
    - Support the `Distributed Database Middleware` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DNS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameters `masters`, `zones` of the interface `CreatePublicZone`: `string`->`array`

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameter `publicip_border_group` to the interfaces `CreateSharedBandwidth`,`ListBandwidths`.

### HuaweiCloud SDK GaussDB

- _Features_
    - Support `GaussDB` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `__root_origin`,`checksum`,`size` to the interfaces `GlanceCreateImageMetadata`.
    - Remove the request parameters `deleted`, `deleted_at` of the interface `GlanceAddImageMember`, and add the following request parameters:
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

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `edge_node_ids`, `last_update_time` to the interface `ListRules`.

### HuaweiCloud SDK LTS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameter `context` of the interface `ListStructuredLogsWithTimeRange`: `string`->`array`
    - Modify the name of the interfaces::
        - `UpdateLogContents`->`ListLogs`
        - `UpdateLogContents2`->`ListQueryStructuredLogs`
        - `UpdateLogContents3`->`ListStructuredLogsWithTimeRange`

### HuaweiCloud SDK MRS

- _Features_
    - Support the interface `ListClusters`.
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameter `clusterType` of the interface `DescribeCluster`: `string`->`integer`

### HuaweiCloud SDK SWR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `domain_id`,`priority` to the interface `ShowRepository`.
    - Add the response parameter `template` to the interface `CreateRetention`.

# 0.0.50 2021-06-29

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add a request parameter `storage` to interfaces `CreateNodePool`,`ShowNodePool`,`UpdateNodePool`,`DeleteNodePool`.

### HuaweiCloud SDK DRS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the parameter `selected` of the interface `BatchUpdateUser`: `string`->`boolean`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameter `port` of the interface `ListInstances`: `string`->`integer`.
    - Modify the name of response parameter of the interface `ListInstances`: `storage_engine`->`mode`
    - Remove a response parameter `node_name` and add a response parameter `time` to the interface `ListSlowLogs`.

### HuaweiCloud SDK MRS

- _Features_
    - None
- _Bug Fix_
    - Fix the issue of incorrect definition of some parameters.
- _Change_
    - Remove the parameters `start_time`,`state` of the interface `CreateCluster`.

### HuaweiCloud SDK NAT

- _Features_
    - None
- _Bug Fix_
    - Fix the issue that the request parameter `project_id` of the interface `ListNatGateways` is duplicated.
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of response parameters `port`,`node_num` of the interface `ShowInformationAboutDatabaseProxy`: `string`->`integer`

# 0.0.49 2021-06-25

### HuaweiCloud SDK APIG

- _Features_
    - Support more ineterfaces:
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
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BMS

- _Features_
    - Support interface `ChangeBaremetalServerOs`.
- _Bug Fix_
    - None
- _Change_
    - Modify the name of reponse parameter of interface `ChangeBaremetalServerName`: `server_tags`->`sys_tags`.

### HuaweiCloud SDK CDN

- _Features_
    - Support interface `ShowQuota`.
- _Bug Fix_
    - None
- _Change_
    - Modify the type of request parameter `url` of interface `ShowHistoryTaskDetails`: `integer`->`string`.

### HuaweiCloud SDK CloudRTC

- _Features_
    - Support `CloudRTC` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DRS

- _Features_
    - Support interface `ShowQuotas`.
- _Bug Fix_
    - None
- _Change_
    - Modify the type of request parameters `is_transfer`,`selected` of interface `BatchUpdateUser`: `string`->`boolean`.

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request parameters`permission_type`,`display_name`,`catalog`,`type` of interface `KeystoneListPermissions`.

### HuaweiCloud SDK LTS

- _Features_
    - Support `Log Tank Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Meeting

- _Features_
    - Support interface `InviteShare`.
- _Bug Fix_
    - None
- _Change_
    - Add request parameter `multiPicSaveOnly` to interface `SetMultiPicture`.
    - Add reponse parameter `leftReason` to interface `SearchHisMeetings`.

### HuaweiCloud SDK VOD

- _Features_
    - Support `Video on Demand` service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.48 2021-06-21

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add response parameters `server_tags`,`enterprise_project_id`,`group` to interface `ChangeBaremetalServerName`.

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - [Issue 22](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/22): Modify the optional value of response parameter `status` of interface `ListAddonInstances`.
- _Change_
    - None

### HuaweiCloud SDK CDN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the request parameter `user_domain_id` of interface `ListDomains`.
    - Modify the name of interface: `ShowRefer` -> `ShowReferer`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request parameters to interface `ShowTemplateDetail`:
        - `template_url`
        - `create_time`
        - `last_modify_time`
        - `can_update`
        - `can_delete`
        - `need_hub`

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces:
        - `CreateRecordCallbackConfig`
        - `ShowRecordCallbackConfig`
        - `UpdateRecordCallbackConfig`
        - `DeleteRecordCallbackConfig`
        - `ListRecordCallbackConfigs`
        - `UpdateRecordRule`
        - `ShowRecordRule`
- _Bug Fix_
    - None
- _Change_
    - Modify the name of some interfaces:
        - `CreateRecordConfig` -> `CreateRecordRule`
        - `DeleteRecordConfig` -> `DeleteRecordRule`
        - `ListRecordConfigs` -> `ListRecordRules`
    - Remove some interfaces:
        - `ShowTraffic`
        - `ShowBandwidth`
        - `ShowOnlineUsers`

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of response parameter `partitions` of interface `ShowGroups`: `array[string]` -> `array[integer]`

### HuaweiCloud SDK RabbitMQ

- _Features_
    - None
- _Bug Fix_
    - Fix the issue of compilation failure.
- _Change_
    - None

# 0.0.47 2021-06-10

### HuaweiCloud SDK BSS

- _Features_
    - Support interfaces `ListFreeResources`,`ListFreeResourceUsages`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CDN

- _Features_
    - Support `Content Delivery Network` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DRS

- _Features_
    - Support `Data Replication Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support interfaces
        - `ImportFunction`
        - `ExportFunction`
        - `AsyncInvokeReservedFunction`
        - `DeleteReservedInstanceById`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK OSM

- _Features_
    - Support `Online Service Management`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support interfaces `SetBinlogClearPolicy`,`ShowBinlogClearPolicy`.
- _Bug Fix_
    - None
- _Change_
    - Add request parameters `offset`,`limit` to interface `ListOffSiteInstances`.

# 0.0.46 2021-06-04

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - [Issue 20](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/20): Fix the issue that the type of `extendParam`
      is defined incorrectly.
- _Change_
    - Add the request parameter `tobedeleted` to the interface `DeleteCluster`.

### HuaweiCloud SDK DDS

- _Features_
    - Support the interface `ShowQuotas`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IoTDA

- _Features_
    - Support interface `ListComplexQueryDevice`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
    - Support `GaussDBforNoSQL` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support the interface `ShowQuotas`
- _Bug Fix_
    - None
- _Change_
    - Modify the type of request parameter `restart` of the interface `StartInstanceRestartAction`: string -> object

# 0.0.45 2021-05-25

### HuaweiCloud SDK AS

- _Features_
    - Support more interfaces:
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
- _Bug Fix_
    - None
- _Change_
    - Modify the operation name of the following interfaces:
        - from `ExecuteScalingPolicies` to `BatchDeleteScalingPolicies`
        - from `EnableOrDisableScalingGroup` to `ResumeScalingGroup`
        - from `UpdateScalingGroupInstance` to `BatchAddScalingInstances`
        - from `CompleteLifecycleAction` to `AttachCallbackInstanceLifeCycleHook`
    - Remove the interface: `DeleteScalingTags`
    - Add the parameter `enterprise_project_id` to the interface `ListScalingGroups`.
    - Add the parameter `log_id` to the interface `ListScalingActivityV2Logs`.

### HuaweiCloud SDK BSS

- _Features_
    - Support interface `ListCustomerBillsMonthlyBreakDown` and `ListOrderDiscounts`.
- _Bug Fix_
    - None
- _Change_
    - Add query parameter _bill_date_begin_ and _bill_date_end_ to interface `ListSubCustomerResFeeRecords`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - Support interface: `StopPipelineNew`.
- _Bug Fix_
    - None
- _Change_
    - Remove interfaces `StartPipeline`, `StopPipeline`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove interfaces `StartPipeline`,`StopPipeline`.

### HuaweiCloud SDK DMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of interface from `ShowProjectTags` to `ShowQueueProjectTags`.

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change request parameter `offset` of interface `ListEnterpriseProject` from required to optional.

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support more interfaces:
        - `ListFunctionAsyncInvokeConfig`
        - `ShowFunctionAsyncInvokeConfig`
        - `DeleteFunctionAsyncInvokeConfig`
        - `UpdateFunctionAsyncInvokeConfig`
- _Bug Fix_
    - None
- _Change_
    - Modify the name of request parameter of interfaces `DeleteVersionAlias`,`UpdateVersionAlias`
      ,`ShowVersionAlias`: `name` -> `alias_name`
    - Modify the name of request parameter of interfaces `DeleteFunctionTrigger`,`UpdateTrigger`
      ,`ShowFunctionTrigger`: `triggerId` -> `trigger_id`

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the parameter `access_mode` to interface `CreateUsers`.
    - Change the parameter `authentication_code` of interface `DeleteBindingDevice` from required to optional.

### HuaweiCloud SDK Kafka

- _Features_
    - Support more interfaces:
        - `CreateInstanceUser`
        - `BatchDeleteInstanceUsers`
        - `ShowInstanceUsers`
        - `ShowTopicAccessPolicy`
        - `UpdateTopicAccessPolicy`
        - `ShowKafkaTopicPartitionDiskusage`
        - `ShowInstanceMessages`
        - `ResetUserPassword`
- _Bug Fix_
    - None
- _Change_
    - Modify the name of the interface:
        - from `ShowInstanceTags` to `ShowKafkaTags`
        - from `ShowProjectTags` to `ShowKafkaProjectTags`
        - from `BatchCreateOrDeleteInstanceTag` to `BatchCreateOrDeleteKafkaTag`
    - Modify the request body name of the interface:
        - from `BatchCreateOrDeleteInstanceTagRequestBody` to `BatchCreateOrDeleteKafkaTagRequestBody`
    - Modify the data type of parameter `sink_max_tasks` in the request body of interface `UpdateSinkTaskQuota` from
      String to Integer.

### HuaweiCloud SDK OMS

- _Features_
    - Support `Object Storage Migration Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RabbitMQ

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of the following interfaces:
        - from `BatchCreateOrDeleteInstanceTag` to `BatchCreateOrDeleteRabbitMqTag`;
        - from `ShowProjectTags` to `ShowRabbitMqProjectTags`;
        - from `ShowInstanceTags` to `ShowRabbitMqTags`.
    - Modify the request body name of interface `BatchCreateOrDeleteInstanceTag`
      from `BatchCreateOrDeleteInstanceTagRequestBody` to
      `BatchCreateOrDeleteRabbitMqTagRequestBody`.

# 0.0.43-rc 2021-05-14

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Solve the issue of abnormal parsing result when using interface `NovaShowKeypair` to obtain the secret key.
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add new result values `CLOUDSSD` and `LOCALSSD` to response field `type` of interface `ListInstances`.
    - Add an optional request parameter `complete_version` to interface `ListBackups`.
    - Change request parameter `type` of interface `ListSlowlogStatistics` from optional to required.

# 0.0.42-rc 2021-05-10

### HuaweiCloud SDK BMS

- _Features_
    - Support interface `BatchCreateBaremetalServerTags`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support interfaces `MigrateAz`, `ListAz2Migrate`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/17): Fix the issue that `EpDetailType` enum
      is defined incorrectly.
- _Change_
    - None

### HuaweiCloud SDK KPS

- _Features_
    - None
- _Bug Fix_
    - [Issue 19](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/19): Fix the issue of type of response body
      in `ListKeypairs`.
- _Change_
    - None

### HuaweiCloud SDK MRS

- _Features_
    - Support more interfaces:
        - `BatchDeleteClusterTags`
        - `CreateClusterTag`
        - `DeleteClusterTag`
        - `ListClusterTags`
        - `ListAllTags`
        - `BatchCreateClusterTags`
        - `ListClustersByTags`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of response body of interface `ListOffSiteInstances`: `OffSiteInstanceListResponse`
      -> `OffSiteInstanceListResponseBody`
    - Modify the name of response field of interface `ListOffSiteInstances`: `offsite_backup_instances`
      -> `offsite_backup_instance`

# 0.0.41-rc 2021-04-30

### HuaweiCloud SDK BCS

- _Features_
    - Support interface `ListOpRecord`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support interfaces:
        - `ShowShardingBalancer`
        - `SetBalancerSwitch`
        - `SetBalancerWindow`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK HSS

- _Features_
    - Support interface `ListHosts`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add optional values to request parameter `type` of the interface `ShowDomainQuota`:
        - `assigment_group_mp`
        - `assigment_agency_mp`
        - `assigment_group_ep`
        - `assigment_user_ep`

### HuaweiCloud SDK IoTDA

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove interfaces:
        - `ListSubscriptions`
        - `CreateSubscription`
        - `UpdateSubscription`
        - `ShowSubscription`
        - `DeleteSubscription`

### HuaweiCloud SDK MPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request parameters `language`、`sky_switch` to the interface `CreateMpeCallBack`.
    - Update optional values of request parameter `subtitle_type` of interface `CreateTranscodingTask`.

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add a field `project_code` to response body of the interface `ShowProjectInfoV4`.

# 0.0.40-rc 2021-04-15

### HuaweiCloud SDK Moderation

- _Features_
    - Support `Moderation` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Image

- _Features_
    - Support `Image` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support more interfaces about database management operations.
        - `CreateSqlserverDatabase`
        - `DeleteSqlserverDatabase`
        - `ListSqlserverDatabases`
    - Support more interfaces about user management operations.
        - `CreateSqlserverDbUser`
        - `ListSqlserverDbUsers`
        - `ListAuthorizedSqlserverDbUsers`
        - `DeleteSqlserverDbUser`
        - `AllowSqlserverDbUserPrivilege`
        - `RevokeSqlserverDbUserPrivilege`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces `DeleteDatabaseUser`,`DeleteDatabaseRole`,`ShowConnectionStatistics`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add fields `closed_time` ,`id` ,`created_time` to reponse body of interfaces `ListIssuesV4`, `ListChildIssuesV4`.

### HuaweiCloud SDK VPC

- _Features_
    - None
- _Bug Fix_
    - Fix the bug, open the tags of the VPC and subnet.
- _Change_
    - None

# 0.0.39-rc 2021-03-30

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - Fix the issue that the interface for querying messages does not contain the timestamp field.
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the attribute `name` to the response body `IssueResponseV4` of the interfaces `ShowIssueV4`
      and `UpdateIssueV4`.
    - Change the attribute `work_time` to `work_date` in `ShowProjectWorkHoursResponseBody` in the response body of the
      interfaces `ShowProjectWorkHours` and `ListProjectWorkHours`.

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change the request parameter `protocol`  of the interface `PublishMessage`  from mandatory to optional.
    - Change the attribute `subject`  of the class `PublishMessageRequestBody` in the request body of the
      interface `PublishMessage`  from mandatory to optional.

# 0.0.38-rc 2021-03-26

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that authorization failed in auto acquisition of domain id.
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of deserialization error of the response of interface `ListLiveStreamsOnline`.
- _Change_
    - None

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change the property `protocol` in `ListMessageTemplates` from required to optional.

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that some fields in the response body of interface `ListSlowlogStatistics` are empty.
- _Change_
    - None

# 0.0.37-rc 2021-03-19

### HuaweiCloud SDK Core

- _Features_
    - Support the file upload feature.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of deserialization failure of response body of interface `ListFlavors`.
- _Change_
    - None

# 0.0.36-rc 2021-03-16

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request field `enterprise_project_id` in interface `CreatePrePaidPublicip`.

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that the SDK can not be used.
- _Change_
    - None

# 0.0.35-rc 2021-03-15

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of deserialization failure when the response body of the API contains special characters such
      as `\n`.
- _Change_
    - If the `endpoint` input by the user does not contain a protocol prefix, the `https` prefix will be automatically
      added.

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Class adjustment in interface `CreateAlarmRequestBody`: change class definition of property `metric`
      from `MetricInfoForAlarm` to `MetricForAlarm`.
    - Make the property `name` in class `MetricsDimension` required, which affects interfaces of `BatchListMetricData`
      , `CreateMetricData`, `CreateResourceGroup` , and `UpdateResourceGroup`.

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces:
        - `RestoreNewInstance`
        - `ListSessions`
        - `DeleteSession`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - Support more interface: `ShowServerGroup`.
- _Bug Fix_
    - None
- _Change_
    - Change the interface name from `ShowWindowsServerPassword` to `ShowServerPassword`.
    - Change the interface name from `DeleteWindowsServerPassword` to `DeleteServerPassword`.

### HuaweiCloud SDK ELB

- _Features_
    - Support more interface: `ListAllMembers`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK FunctionGraph

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `ListDependencies` adjustment: change the data type of property `size` of the response definition from
      string to int64.

### HuaweiCloud SDK IAM

- _Features_
    - Support more interfaces:
        - `KeystoneShowIdentityProvider`
        - `KeystoneCreateIdentityProvider`
        - `KeystoneUpdateIdentityProvider`
        - `KeystoneDeleteIdentityProvider`
        - `CreateTokenWithIdToken`
- _Bug Fix_
    - None
- _Change_
    - Do not support interface `CreateUnscopeTokenByIdpInitiated` anymore.

### HuaweiCloud SDK IMS

- _Features_
    - Support more interfaces:
        - `ListImageByTags` which mead list images queried by tags.
        - `ListImagesTags` which means list all tags of all images in current account.
        - `ListImageTags` which means list all tags of specified image.
        - `AddImageTag`
        - `DeleteImageTag`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IoTDA

- _Features_
    - Support `IoT Device Access` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - Support more interfaces:
        - `CreateCustomfields`
        - `ShowBugsPerDeveloper`
        - `ShowCompletionRate`
        - `ShowBugDensityV2`
        - `ShowProjectInfoV4`
- _Bug Fix_
    - Change the incorrect name of interface from `ShowtIssueCompletionRate` to `ShowIssueCompletionRate`.
- _Change_
    - Change the data type of property `created_time` and `updated_time` in class `ListProjectV4ResponseBody` from
      string to int64.

### HuaweiCloud SDK RDS

- _Features_
    - Support `Postgresql` related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK SCM

- _Features_
    - Support `SSL Certificate Management` service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.34-rc 2021-02-27

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Optimize the description of `README` and the format of `CHANGELOG`.

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `WindowsBaremetalServerCleanPwd` to `DeleteWindowsBareMetalServerPassword`.
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - All properties with type `float32` have been changed to `float64`.

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of property `Content-Type` is required when sending requests and returns
      error `Unsupported Content Type`.
- _Change_
    - None

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of circular dependency in the `CreateAlarmResponse` class.
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces: `DownloadSlowlog` and `DownloadErrorlog`.
- _Bug Fix_
    - Correct operation name from `ModifyConfigurationParameter` to `UpdateConfigurationParameter`, change the class
      name of this operation from `ModifyConfigurationParameterRequestBody` to `UpdateConfigurationParameterRequestBody`
      .
- _Change_
    - None

### HuaweiCloud SDK DGC

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `ModifyJob` to `UpdateJob`.
    - Correct operation name from `ModifyScript` to `UpdateScript`.
    - Correct operation name from `ModifyResource` to `UpdateResource`.
- _Change_
    - None

### HuaweiCloud SDK DIS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of circular dependency in the classes of `CommitCheckpointRequest`, `PutRecordsRequest`
      , `CreateAppRequest`, `UpdatePartitionCountRequest`.
- _Change_
    - None

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `ListMenbers` to `ListMembers`.
- _Change_
    - None

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `ModifyEnterpriseProject` to `UpdateEnterpriseProject`.
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - Correct property from `pwd_stength` to `pwd_strength` in class `KeystoneUserResult`.
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `DoManualBackup` to `CreateManualBackup`.
    - Correct operation name from `ModifyConfiguration` to `UpdateConfiguration`.
    - Correct operation name from `ModifyInstanceConfiguration` to `UpdateInstanceConfiguration`.
    - Fix the problem of circular dependency in the classes of `CreateInstanceResponse`
      and `CreateConfigurationResponse`.
    - Fix the unavailable problem of operation `CreateInstance`.
- _Change_
    - Add property `is_auto_pay` to the operation `StartInstanceAction` in the scenario of changing a single-node system
      to a primary/standby mode.

# 0.0.33-rc 2021-02-07

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that request sends fail when the data type of request body is `string`.
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `ListOsVersions` adjustment: change the data type of `os_bit` which is the property of response of the
      interface from string to integer.

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces: ListLiveSampleLogs, CreateDomain, DeleteDomain, UpdateDomain, ShowDomain,
      CreateDomainMapping, DeleteDomainMapping
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK MRS

- _Features_
    - Support `MapReduce Service`.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.32-rc 2021-01-30

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces: `ListApiVersion`, `ShowApiVersion`,`SetAuditlogPolicy`, `ShowAuditlogPolicy`
      , `ListAuditlogs`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DNS

- _Features_
    - Support `Domain Name Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change interface name from `UpdateAutoTerminateTimeServer` to `UpdateServerAutoTerminateTime`.

### HuaweiCloud SDK EVS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `CinderCreateVolume` is supported to specify the id of dedicated storage pool using
      property `OS-SCH-HNT:scheduler_hints`.
    - Modify property type of `allocated` of class `QuotaDetails` from `String` to `Integer`.

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Increases the property `access_mode` of response class of interface `ShowUser`.

# 0.0.31-rc 2021-01-25

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - The default value of `DefaultTimeout` is set to 120 seconds.

### HuaweiCloud SDK BSS

- _Features_
    - Support more interface: ListOrderDiscounts.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DNS

- _Features_
    - Support `Domain Name Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - Support more interface: `UpdateAutoTerminateTimeServer`.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.30-rc 2021-01-15

### HuaweiCloud SDK Core

- _Features_
    - Support function `ValueOf` to get region information.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CloudBuild

- _Features_
    - Support more interface: `ShowListHistory`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DGC

- _Features_
    - Support more interfaces: `Job` related interfaces, `Script` related interfaces, `Resource` related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the data type of response field `is_domain_owner` from string to boolean of interface `ShowUser`
      and `CreateUser`.

### HuaweiCloud SDK Live

- _Features_
    - Support more interface: `ListLiveStreamsOnline`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support more interfaces: ShowOffSiteBackupPolicy, SetOffSiteBackupPolicy, ListOffSiteBackups,
      ListOffSiteRestoreTimes, ListOffSiteRestoreTimes
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK SWR

- _Features_
    - Support `Software Repository for Container` service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.29-beta 2020-12-31

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of interface: ListBareMetalServers.
    - Fix the problem of interface: ListBareMetalServerDetails.
    - Modify interface fields: ShowBaremetalServerInterfaceAttachments.
- _Change_
    - None

### HuaweiCloud SDK CloudIDE

- _Features_
    - Support more interface: ShowAccountStatus.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - Modify the interface return data type to prevent deserialization failure:
        - ListSlowlog: change data type of `Tags` from Tag to ResourceTag.
        - ListInstances: change data type of `duration` from int32 to string.
        - ShowBigkeyScanTaskDetails: change data type of `db` from int32 to string.
        - ShowHotkeyTaskDetails: change data type of `db` from int32 to string.
- _Change_
    - None

### HuaweiCloud SDK DGC

- _Features_
    - Support `Data Lake Governance Center` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DIS

- _Features_
    - Support `Data Ingestion Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface modification: response type of interface `CreateInstance` adjustment.

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - Modify the request parameters of interface `PublishMessage` from Object to Map<String, String>
- _Change_
    - None

# 0.0.28-beta 2020-12-28

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix response exception when using temporary AK/SK.
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - Change property type of `port` from string to integer.
- _Change_
    - None

# 0.0.27-beta 2020-12-25

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of compilation failure in Mac OS.
- _Change_
    - Query parameter in interface `ListInstances` modification: id → instance_id.

### HuaweiCloud SDK DDS

- _Features_
    - Support Document Database Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of compilation failure in Mac OS.
- _Change_
    - None

### HuaweiCloud SDK RabbitMQ

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of compilation failure in Mac OS.
- _Change_
    - None

### HuaweiCloud SDK RMS

- _Features_
    - Support more interfaces: `Resources` related interfaces and `Region` related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.26-beta 2020-12-23

### HuaweiCloud SDK Core

- _Features_
    - Support Endpoint Resolver: it's supported to use {Service}Region when initializing {ServiceClient} which can
      automatically backfill endpoint. After choosing a region, the projectId/domainId will be backfilled automatically.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - Support more interfaces: ListMeasureUnits.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Update interface: ShowMetricData

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces: ShowStreamPortrait.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK MPC

- _Features_
    - Support more interfaces: QualityEnhanceTemplate related interfaces and MergeChannelsTask related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support Relational Database Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Update field type in message_template_name.

# 0.0.25-beta 2020-12-15

### HuaweiCloud SDK CCE

- _Features_
    - Support Cloud Container Engine service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that sending request to interface `CreateListener` returns empty response.
    - Fix the problem that sending request to interface `CreateListener` returns response with wrong type.
- _Change_
    - None

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support more interfaces: UpdateFunctionReservedInstances.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK NAT

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that using interface `BatchCreateNatGatewayDnatRules` failed.
- _Change_
    - None

# 0.0.24-beta 2020-12-04

### HuaweiCloud SDK SMN

- _Features_
    - Support Simple Message Notification service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.23-beta 2020-11-30

### HuaweiCloud SDK BCS

- _Features_
    - Support BlockChain Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BMS

- _Features_
    - Support Bare Metal Server service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - Support more interfaces: ListUsageTypes, ModPeriodToOnDemand.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - Support more interfaces: MigrateVaultResource
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CES

- _Features_
    - Support more interfaces:
        - ListEvents
        - ListEventDetail
        - CreateResourceGroup
        - UpdateResourceGroup
        - DeleteResourceGroup
        - ListResourceGroup
        - UpdateAlarm
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - Support Distributed Cache Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - [Issue 21](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/21) Open related interface.

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support FunctionGraph service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - Support more interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Name of service client modification: LiveAPIClient → LiveClient.

# 0.0.22-beta 2020-11-17

### HuaweiCloud SDK AS

- _Features_
    - None
- _Bug Fix_
    - [Issue 8](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/8) Fix the problem that creating scaling
      policy failed.
- _Change_
    - None

### HuaweiCloud SDK DMS

- _Features_
    - Support Distributed Message Services, provide Kafka premium instances and RabbitMQ premium instances with
      dedicated resources.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Property adjustment:  increase property `dry_run` in interfaces `CreatePostPaidServers` and `CreateServers` which
      means whether parameters will be checked before sending real requests.

### HuaweiCloud SDK NAT

- _Features_
    - Support NAT Gateway service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK VPC

- _Features_
    - Support more interfaces: interfaces related to Network ACLs.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.21-beta 2020-11-11

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Update core code from [Pull requests #11](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/pull/11).
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface adjustment: property `object_type` in interface `CreateVault` support key `turbo`.
    - Interface adjustment: property `description` in interface `UpdatePolicy` is removed.

### HuaweiCloud SDK CES

- _Features_
    - Add examples of interface response and adjust some filed description.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CloudPipeline

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of generated Client class: devcloudpipeline_client → cloudPipeline_client
    - Modify the name of generated Meta class: devcloudpipeline_meta → cloudPipeline_meta

### HuaweiCloud SDK DevStar

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify interface parameters and adjust sample code.

# 0.0.20-beta 2020-11-02

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface adjustment: property `alarm_type` in class `CreateAlarmRequestBody` support key `RESOURCE_GROUP`.
    - Interface adjustment: property `meta_data` in class `ListAlarmHistoriesResponse` only returns total number of
      alarm histories.

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Modify wrong parameter in class `CreateL7ruleRequestBody`.
- _Change_
    - None

# 0.0.19-beta 2020-10-31

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix: fix the problem that when query parameter contains enumerated variables the request will fail.
    - [Issue 7](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/7) resolve the problem of using
      json.Marshal()
      returns object{}.
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - Support more interfaces: interfaces related to `TAG`.
- _Bug Fix_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/17) fix the problem of interface:
      ListBackups.
- _Change_
    - None

### HuaweiCloud SDK CTS

- _Features_
    - Support more interface: ListQuotas
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust interfaces' names, replace abbreviations of `EP` with `EnterpriseProject`, the involved interfaces are:
        1. ListEP → ListEnterpriseProject
        2. CreateEP → CreateEnterpriseProject
        3. ShowEP → ShowEnterpriseProject
        4. ModifyEP → ModifyEnterpriseProject
        5. EnableEP → EnableEnterpriseProject
        6. ShowEPQuota → ShowEnterpriseProjectQuota
        7. ShowResourceBindEP → ShowResourceBindEnterpriseProject
        8. DisableEP → DisableEnterpriseProject

### HuaweiCloud SDK Iam

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust interfaces' names, the involved interfaces are:
        1. KeystoneCreateUserTokenByPasswordAndMFA → KeystoneCreateUserTokenByPasswordAndMfa
        2. CreateUnscopeTokenByIDPInitiated → CreateUnscopeTokenByIdpInitiated

### HuaweiCloud SDK ProjectMan

- _Features_
    - Support more interfaces: iteration information, user information, project members, project information, project
      indicators, project statistics, etc.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.18-beta 2020-10-20

### HuaweiCloud SDK ELB

- _Features_
    - Support more interfaces of version v2.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.17-beta 2020-10-14

### HuaweiCloud SDK BSS

- _Features_
    - Partner center supports exporting product catalog prices.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces of version v2 of Live.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.16-beta 2020-10-12

### HuaweiCloud SDK CTS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Delete deprecated interfaces of version v1.

### HuaweiCloud SDK DevStar

- _Features_
    - None
- _Bug Fix_
    - Modify the credential type of DevStarClient: the correct credential type is GlobalCredentials.
- _Change_
    - None

# 0.0.15-beta 2020-09-30

### HuaweiCloud SDK AS

- _Features_
    - Support Auto Scaling service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.14-beta 2020-09-24

### HuaweiCloud SDK BSS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that the class `BssClient` cannot be loaded.
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `ListPublicips` adjustment: enterprise_project_id does not support multi-value query.

# 0.0.13-beta 2020-09-16

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Fix parameter type of interfaces.
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Update interfaces.

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that not support transfer multiple values.
- _Change_
    - None

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that some parameters are wrong.
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust descriptions of interfaces.

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust descriptions of banned interface.

# 0.0.12.1-beta 2020-09-16

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Fix parameter type of interfaces.
- _Change_
    - None

### HuaweiCloud SDK All

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - When the optional parameter type is list, the parameter type will be changed to a pointer, for example, []string
      to *[]string

# 0.0.12-beta 2020-09-11

### HuaweiCloud SDK KPS

- _Features_
    - Support KPS service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EVS

- _Features_
    - Support EVS service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - Fix response value type definition errors.
- _Change_
    - None

# 0.0.11-beta 2020-09-09

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that resources related interfaces have wrong data type.
- _Change_
    - None

### HuaweiCloud SDK VPC

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that security group related interfaces have wrong data type.
- _Change_
    - None

# 0.0.10-beta 2020-09-04

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that the enumeration code cannot be generated for integer enumeration parameters without format
      defined in yaml.
- _Change_
    - Modify User Agent of Http Request header.

# 0.0.9-beta 2020-08-28

### HuaweiCloud SDK CloudPipeline

- _Features_
    - Support CloudPipeline service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - Support more APIs: tags related APIs and shared bandwidth related APIs.
- _Bug Fix_
    - Interface BatchCreateBandwidth: modify field billing_info.
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - Support more APIs: consistency of console related APIs.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - Support Project Management service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK VPC

- _Features_
    - Support more APIs: security group, security group rules, available IP count related APIs.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.8-beta 2020-08-25

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change optional enum variable type to pointer.

### HuaweiCloud SDK ELB

- _Features_
    - Support Elastic Load Balance service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.7-beta 2020-08-20

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that some enum variables are unreadable.
- _Change_
    - Add 'E' as prefix to enum Variables which start with '_'.

# 0.0.6-beta 2020-08-20

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of missing the imports when the struct contains enum variables.
- _Change_
    - None

# 0.0.5-beta 2020-08-19

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the deserialization failure of enum variables.
    - Fix the deserialization failure of error response in some scenarios.
- _Change_
    - None

# 0.0.4-beta 2020-08-18

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of missing default values of Go basic type when serializing.
- _Change_
    - None

# 0.0.3-beta 2020-08-14

### HuaweiCloud SDK APIG

- _Features_
    - Support API Gateway service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - Support Business Support System.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.2-beta 2020-8-11

### HuaweiCloud SDK Core

- _Features_
    - Support temporary AK/SK authentication mode.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - Support Cloud Backup and Recovery service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CloudIDE

- _Features_
    - Support Cloud IDE service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - Support Elastic Cloud Server service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - Support Elastic IP service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EVS

- _Features_
    - Support Elastic Volume Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - Support Image Management Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - Support Live service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK MPC

- _Features_
    - Support Media Processing Center.
- _Bug Fix_
    - None
- _Change_
    - None

## 3.0.1-beta 2020-07-30

### First Release

- _Supported Services_
    - Classroom
    - Cloud Trace Service(CTS)
    - DevStar
    - Enterprise Project Management Service(EPS)
    - Identity and Access Management(IAM)
    - Tag Management Service(TMS)
    - Virtual Private Cloud(VPC)
