package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Mp4RecordConfig struct {

	//  周期录制时长。  取值范围：[60，10800]。  单位：秒。
	RecordCycle int32 `json:"record_cycle"`

	// 录制文件含路径和文件名的前缀。  默认{app_id}/{record_format}/{stream}\\_{file_start_time}/{stream}\\_{file_start_time}  可自定义以下特殊变量： - {app_id}：应用id - {record_format}：录制格式 - {stream}：流名 - {file_start_time}：文件生成时间
	RecordPrefix *string `json:"record_prefix,omitempty"`

	// 录制MP4拼接时长，如果流中断超过该时间，则生成新文件。  取值范围：[0，300]，缺省为0。  单位：秒。  如果为0表示流中断就生成新文件。
	RecordMaxDurationToMergeFile *int32 `json:"record_max_duration_to_merge_file,omitempty"`
}

func (o Mp4RecordConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Mp4RecordConfig struct{}"
	}

	return strings.Join([]string{"Mp4RecordConfig", string(data)}, " ")
}
