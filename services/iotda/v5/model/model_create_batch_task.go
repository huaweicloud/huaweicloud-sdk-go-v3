package model

import (
	"encoding/json"

	"strings"
)

// 添加任务结构体
type CreateBatchTask struct {
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的批量任务归属到哪个资源空间下，否则创建的批量任务将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。
	AppId *string `json:"app_id,omitempty"`
	// 批量任务名称。
	TaskName string `json:"task_name"`
	// 批量任务类型，取值范围：firmwareUpgrade，softwareUpgrade，createDevices，deleteDevices，freezeDevices，unfreezeDevices，createCommands，createAsyncCommands。 - softwareUpgrade: 软件升级任务 - firmwareUpgrade: 固件升级任务 - createDevices: 批量创建设备任务 - deleteDevices: 批量删除设备任务 - freezeDevices: 批量冻结设备任务 - unfreezeDevices: 批量解冻设备任务 - createCommands: 批量创建同步命令任务 - createAsyncCommands: 批量创建异步命令任务
	TaskType string `json:"task_type"`
	// 执行批量任务的目标，此处填写device_id列表，且最多支持3万个device_id。当task_type为firmwareUpgrade，softwareUpgrade，deleteDevices，createCommands，createAsyncCommands，支持该参数。同时使用targets、targets_filter、document_source参数时，只有一个参数会生效，且平台优先使用targets，其次是targets_filter，最后是document_source。
	Targets *[]string `json:"targets,omitempty"`
	// 任务目标筛选参数。Json格式，里面是一个个键值对，（K,V）格式标识筛选targets需要的参数，目前支持的K有group_ids（V填写group_id数组，eg:[\"e495cf17-ff79-4294-8f64-4d367919d665\"]，任务则会筛选出来符合该群组条件的设备作为目标）。当task_type为firmwareUpgrade，softwareUpgrade，createDevices，deleteDevices，freezeDevices，unfreezeDevices，createCommands，createAsyncCommands，支持该参数。同时使用targets、targets_filter、document_source参数时，只有一个参数会生效，且平台优先使用targets，其次是targets_filter，最后是document_source。
	TargetsFilter map[string]interface{} `json:"targets_filter,omitempty"`
	// 执行任务数据文档，Json格式，Json里面是(K,V)键值对。当task_type为firmwareUpgrade，softwareUpgrade，createCommands，createAsyncCommands，支持该参数。 - softwareUpgrade|firmwareUpgrade，需要填写key为package_id，value为在平台上传的软固件附件id，id由portal软件库包管理上传并查询获得。eg：“{\"package_id\": \"32822e5744a45ede319d2c50\"}”。 - createCommands，需要填写同步命令相关参数，eg：“{\"service_id\":\"water\",\"command_name\":\"ON_OFF\",\"paras\":{\"value\":\"ON\"}}”,参考[设备同步命令](https://support.huaweicloud.com/api-iothub/iot_06_v5_0038.html))。 - createAsyncCommands，需要填写异步命令相关参数，eg：“{\"service_id\":\"water\",\"command_name\":\"ON_OFF\",\"paras\":{\"value\":\"ON\"},\"expire_time\":0,\"send_strategy\":\"immediately\"}”,参考[设备异步命令](https://support.huaweicloud.com/api-iothub/iot_06_v5_0040.html))。
	Document   *interface{} `json:"document,omitempty"`
	TaskPolicy *TaskPolicy  `json:"task_policy,omitempty"`
	// 上传的批量任务文件ID。当task_type为createDevices，deleteDevices，freezeDevices，unfreezeDevices时，支持该参数。使用该参数时，需要先调用批量任务的文件管理接口上传文件来获取文件ID，文件样例请参见 [批量注册设备模板](https://developer.obs.cn-north-4.myhuaweicloud.com/template/BatchCreateDevices_Template.xlsx)，[批量删除设备模板](https://developer.obs.cn-north-4.myhuaweicloud.com/template/BatchDeleteDevices_Template.xlsx)，[批量冻结设备模板](https://developer.obs.cn-north-4.myhuaweicloud.com/template/BatchFreezeDevices_Template.xlsx)，[批量解冻设备模板](https://developer.obs.cn-north-4.myhuaweicloud.com/template/BatchUnfreezeDevices_Template.xlsx)。同时使用targets、targets_filter、document_source参数时，只有一个参数会生效，且平台优先使用targets，其次是targets_filter，最后是document_source。
	DocumentSource *string `json:"document_source,omitempty"`
}

func (o CreateBatchTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateBatchTask struct{}"
	}

	return strings.Join([]string{"CreateBatchTask", string(data)}, " ")
}
