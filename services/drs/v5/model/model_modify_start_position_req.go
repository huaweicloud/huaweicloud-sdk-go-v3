package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyStartPositionReq 更新增量位点请求体。
type ModifyStartPositionReq struct {

	// - MySQL为源通过show master status命令获取源库位点，根据提示分别填写File:Position。例如：mysql-bin.000277:805  文件名只能为1-60个字符且不能包含< > & : \" \\' / \\\\\\\\ 特殊字符，文件编号只能为3-20个数字，binlog事件位置只能为1-20个数字，且总长度不能超过100个字符。格式为：文件名.文件编号:事件位点  - mogo为源的任务，任务的源库日志从位点开始获取（含当前启动位点），位点需设置在oplog范围以内。非集群通过db.getReplicationInfo()直接获得oplog范围，集群通过db.watch([], {startAtOperationTime: Timestamp(xx, xx)})命令，将启动位点填在xx处，校验位点是否在oplog范围以内。格式为：timestamp:incre。timestamp和incre均为范围在1~2,147,483,647之间的整数。
	FileAndPosition string `json:"file_and_position"`

	// MySQL为源的任务需要，通过show master status命令获取源库位点，根据提示填写Executed_Gtid_Set（如果源库为MySQL 5.5版本，则不支持使用同步任务）。 - 不能包含< > & \" \\' / \\\\\\\\ 特殊字符和中文。且不能超过2048个字符
	GtidSet *string `json:"gtid_set,omitempty"`
}

func (o ModifyStartPositionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyStartPositionReq struct{}"
	}

	return strings.Join([]string{"ModifyStartPositionReq", string(data)}, " ")
}
