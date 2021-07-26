package model

import (
	"encoding/json"

	"strings"
)

type HlsRecordConfig struct {
	//  周期录制时长。  取值范围：[60，43200]或者0，如果为0则整个流录制为一个文件。  单位：秒。

	RecordCycle int32 `json:"record_cycle"`
	// 录制m3u8文件含路径和文件名的前缀  默认{app_id}/{record_format}/{stream}\\_{file_start_time}/{stream}\\_{file_start_time}  上述特殊变量含义： - {app_id}：应用id - {record_format}：录制格式 - {stream}：流名 - {file_start_time}：文件生成时间

	RecordPrefix *string `json:"record_prefix,omitempty"`
	//  录制HLS时ts的切片时长，非必填。  取值范围：[2，60]，缺省为10。  单位：秒。

	RecordSliceDuration *int32 `json:"record_slice_duration,omitempty"`
	// 录制HLS文件拼接时长，如果流中断超过该时间，则生成新文件。  取值范围：[-1，300]，缺省为0。  单位：秒。  - 如果为0表示流中断就生成新文件。 - 如果为-1则表示流中断恢复后追加到原来的文件中（相隔不超过30天）。

	RecordMaxDurationToMergeFile *int32 `json:"record_max_duration_to_merge_file,omitempty"`
}

func (o HlsRecordConfig) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HlsRecordConfig struct{}"
	}

	return strings.Join([]string{"HlsRecordConfig", string(data)}, " ")
}
