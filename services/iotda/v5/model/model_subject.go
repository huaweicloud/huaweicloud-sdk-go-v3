package model

import (
	"encoding/json"

	"strings"
)

// 创建或修改订阅主题时，指定订阅资源及其范围
type Subject struct {
	// 订阅的资源名称。 - device：设备。 - device.data：设备数据。 - device.message.status：设备消息状态。 - device.message：设备消息。 - device.status：设备状态。 - batchtask.status：批量任务状态。

	Resource string `json:"resource"`
	// 订阅的资源事件，取值范围：activate、update、up。 event需要与resource关联使用，具体的“resource：event”映射关系如下： - device：activate（设备激活） - device.data：update（设备数据变化） - device.message.status：update（设备消息状态） - device.message：report（设备消息上报） - device.status：update （设备状态变化） - batchtask.status：update （批量任务状态变化）

	Event string `json:"event"`
}

func (o Subject) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Subject struct{}"
	}

	return strings.Join([]string{"Subject", string(data)}, " ")
}
