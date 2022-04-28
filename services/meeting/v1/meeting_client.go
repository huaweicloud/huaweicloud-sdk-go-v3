package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/meeting/v1/model"
)

type MeetingClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMeetingClient(hcClient *http_client.HcHttpClient) *MeetingClient {
	return &MeetingClient{HcClient: hcClient}
}

func MeetingClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("v1.MeetingCredentials")
	return builder
}

// SP管理员创建企业
//
// 创建企业，默认管理员及分配资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddCorp(request *model.AddCorpRequest) (*model.AddCorpResponse, error) {
	requestDef := GenReqDefForAddCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCorpResponse), nil
	}
}

// 添加企业管理员
//
// 企业默认管理员添加企业普通管理员
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddCorpAdmin(request *model.AddCorpAdminRequest) (*model.AddCorpAdminResponse, error) {
	requestDef := GenReqDefForAddCorpAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddCorpAdminResponse), nil
	}
}

// 添加部门
//
// 企业管理员通过该接口添加部门，最多支持10级部门，每级子部门最多支持100个，默认企业最大部门数量为3000个。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddDepartment(request *model.AddDepartmentRequest) (*model.AddDepartmentResponse, error) {
	requestDef := GenReqDefForAddDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDepartmentResponse), nil
	}
}

// 增加终端
//
// 企业管理员通过该接口添加硬终端。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddDevice(request *model.AddDeviceRequest) (*model.AddDeviceResponse, error) {
	requestDef := GenReqDefForAddDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeviceResponse), nil
	}
}

// 新增信息窗素材
//
// 新增信息窗素材（上传素材文件）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddMaterial(request *model.AddMaterialRequest) (*model.AddMaterialResponse, error) {
	requestDef := GenReqDefForAddMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddMaterialResponse), nil
	}
}

// 新增信息窗节目
//
// 新增信息窗节目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddProgram(request *model.AddProgramRequest) (*model.AddProgramResponse, error) {
	requestDef := GenReqDefForAddProgram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddProgramResponse), nil
	}
}

// 新增信息窗发布
//
// 新增信息窗发布
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddPublication(request *model.AddPublicationRequest) (*model.AddPublicationResponse, error) {
	requestDef := GenReqDefForAddPublication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPublicationResponse), nil
	}
}

// SP管理员分配企业资源
//
// 企业新增资源发放。优化适配，该接口同时支持修改，带resourceId后会判断该资源是否存在，存在即修改（支持修改的参数见修改接口），否则按新增处理
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddResource(request *model.AddResourceRequest) (*model.AddResourceResponse, error) {
	requestDef := GenReqDefForAddResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddResourceResponse), nil
	}
}

// 保存会议纪要到个人云空间
//
// 用户使用手机扫码后，手机端请求服务端将当前会议纪要文件保存到个人云空间。二维码内容  cloudlink://cloudlink.huawei.com/h5page?action&#x3D;SAVE_MEETING_FILE&amp;key1&#x3D;value1&amp;key2&#x3D;value2    key/value的个数可能变化，终端解析后，在发起后续请求时，将所有key/value存为map，作为入参即可。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddToPersonalSpace(request *model.AddToPersonalSpaceRequest) (*model.AddToPersonalSpaceResponse, error) {
	requestDef := GenReqDefForAddToPersonalSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddToPersonalSpaceResponse), nil
	}
}

// 添加用户
//
// 企业管理员通过该接口添加企业用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AddUser(request *model.AddUserRequest) (*model.AddUserResponse, error) {
	requestDef := GenReqDefForAddUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddUserResponse), nil
	}
}

// 与会者自己解除静音
//
// 决定与会者是否可以自己解除静音。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AllowGuestUnmute(request *model.AllowGuestUnmuteRequest) (*model.AllowGuestUnmuteResponse, error) {
	requestDef := GenReqDefForAllowGuestUnmute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowGuestUnmuteResponse), nil
	}
}

// 分配云会议室
//
// 企业管理员通过该接口将云会议室分配给用户、硬终端（当前仅支持分配TE10、TE20、HUAWEI Board、HUAWEI Bar 500及HUAWEI Box系列硬件终端）。云会议室分配给硬件终端后，需要重启或重新激活硬件终端。若需要管理云会议室、预约会议、录制会议或进行完整的会控操作，请同时将该云会议室分配给会议用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) AssociateVmr(request *model.AssociateVmrRequest) (*model.AssociateVmrResponse, error) {
	requestDef := GenReqDefForAssociateVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateVmrResponse), nil
	}
}

// 批量删除企业管理员
//
// 批量删除企业管理员
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteCorpAdmins(request *model.BatchDeleteCorpAdminsRequest) (*model.BatchDeleteCorpAdminsResponse, error) {
	requestDef := GenReqDefForBatchDeleteCorpAdmins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteCorpAdminsResponse), nil
	}
}

// 批量删除终端
//
// 企业管理员通过该接口批量删除终端，返回删除失败的列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteDevices(request *model.BatchDeleteDevicesRequest) (*model.BatchDeleteDevicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDevicesResponse), nil
	}
}

// 删除信息窗素材
//
// 删除信息窗素材
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteMaterials(request *model.BatchDeleteMaterialsRequest) (*model.BatchDeleteMaterialsResponse, error) {
	requestDef := GenReqDefForBatchDeleteMaterials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMaterialsResponse), nil
	}
}

// 删除信息窗节目
//
// 删除信息窗节目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchDeletePrograms(request *model.BatchDeleteProgramsRequest) (*model.BatchDeleteProgramsResponse, error) {
	requestDef := GenReqDefForBatchDeletePrograms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteProgramsResponse), nil
	}
}

// 删除信息窗发布
//
// 删除信息窗发布
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchDeletePublications(request *model.BatchDeletePublicationsRequest) (*model.BatchDeletePublicationsResponse, error) {
	requestDef := GenReqDefForBatchDeletePublications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePublicationsResponse), nil
	}
}

// 批量删除用户
//
// 企业管理员通过该接口批量删除企业用户，全量成功或全量失败。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchDeleteUsers(request *model.BatchDeleteUsersRequest) (*model.BatchDeleteUsersResponse, error) {
	requestDef := GenReqDefForBatchDeleteUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUsersResponse), nil
	}
}

// 批量修改终端状态
//
// 批量修改终端状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchUpdateDevicesStatus(request *model.BatchUpdateDevicesStatusRequest) (*model.BatchUpdateDevicesStatusResponse, error) {
	requestDef := GenReqDefForBatchUpdateDevicesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateDevicesStatusResponse), nil
	}
}

// 批量修改用户状态
//
// 企业管理员通过该接口批量修改用户状态，当用户账号数资源或者第三方电子白板资源到期后，若企业内对应资源的用户账号超过数量后会被系统随机自动停用，此时可通过该接口修改用户的状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BatchUpdateUserStatus(request *model.BatchUpdateUserStatusRequest) (*model.BatchUpdateUserStatusResponse, error) {
	requestDef := GenReqDefForBatchUpdateUserStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUserStatusResponse), nil
	}
}

// 广播会场
//
// 同一时间，只允许一个与会者被广播。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) BroadcastParticipant(request *model.BroadcastParticipantRequest) (*model.BroadcastParticipantResponse, error) {
	requestDef := GenReqDefForBroadcastParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BroadcastParticipantResponse), nil
	}
}

// 取消预约会议
//
// 取消预约会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CancelMeeting(request *model.CancelMeetingRequest) (*model.CancelMeetingResponse, error) {
	requestDef := GenReqDefForCancelMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelMeetingResponse), nil
	}
}

// 取消周期会议
//
// 管理员或UC账号可以通过该接口取消周期会议
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CancelRecurringMeeting(request *model.CancelRecurringMeetingRequest) (*model.CancelRecurringMeetingResponse, error) {
	requestDef := GenReqDefForCancelRecurringMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRecurringMeetingResponse), nil
	}
}

// 取消周期子会议
//
// 管理员或UC账号可以通过该接口取消周期会议
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CancelRecurringSubMeeting(request *model.CancelRecurringSubMeetingRequest) (*model.CancelRecurringSubMeetingResponse, error) {
	requestDef := GenReqDefForCancelRecurringSubMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRecurringSubMeetingResponse), nil
	}
}

// 校验滑块验证码
//
// 该接口提供校验滑块验证码。服务器收到请求，返回校验结果。用户在前台界面通过滑块操作匹配图形，使得抠图和原图吻合。然后服务器进行校验滑块验证码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CheckSlideVerifyCode(request *model.CheckSlideVerifyCodeRequest) (*model.CheckSlideVerifyCodeResponse, error) {
	requestDef := GenReqDefForCheckSlideVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckSlideVerifyCodeResponse), nil
	}
}

// 校验Token
//
// 该接口提供校验token合法性功能。服务器收到请求后，验证token合法性并返回结果。如果参数needGenNewToken为true时，生成新的token并返回。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CheckToken(request *model.CheckTokenRequest) (*model.CheckTokenResponse, error) {
	requestDef := GenReqDefForCheckToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTokenResponse), nil
	}
}

// 校验手机和邮箱对应的验证码
//
// 企业用户通过该接口校验手机和邮箱对应的验证码，一分钟内记录尝试次数不得超过5次。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CheckVeriCodeForUpdateUserInfo(request *model.CheckVeriCodeForUpdateUserInfoRequest) (*model.CheckVeriCodeForUpdateUserInfoResponse, error) {
	requestDef := GenReqDefForCheckVeriCodeForUpdateUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVeriCodeForUpdateUserInfoResponse), nil
	}
}

// 校验验证码
//
// 该接口提供校验验证码，服务器收到请求，返回结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CheckVerifyCode(request *model.CheckVerifyCodeRequest) (*model.CheckVerifyCodeResponse, error) {
	requestDef := GenReqDefForCheckVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckVerifyCodeResponse), nil
	}
}

// 匿名用户会议鉴权
//
// 未登陆终端，通过输入会议ID进行会议鉴权，返回鉴权随机数。如果需要密码则返回需要会议密码错误码，然后终端弹出输入会议ID输入框，用户输入密码后，终端再次调用该接口进行鉴权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateAnonymousAuthRandom(request *model.CreateAnonymousAuthRandomRequest) (*model.CreateAnonymousAuthRandomResponse, error) {
	requestDef := GenReqDefForCreateAnonymousAuthRandom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAnonymousAuthRandomResponse), nil
	}
}

// 获取会控Token
//
// 获取会控授权令牌，然后会议会被拉起。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateConfToken(request *model.CreateConfTokenRequest) (*model.CreateConfTokenResponse, error) {
	requestDef := GenReqDefForCreateConfToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfTokenResponse), nil
	}
}

// 创建会议
//
// 您可根据需要创建立即会议和预约会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateMeeting(request *model.CreateMeetingRequest) (*model.CreateMeetingResponse, error) {
	requestDef := GenReqDefForCreateMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMeetingResponse), nil
	}
}

// 获取页面免登陆跳转的nonce信息
//
// 通过token生成页面免登陆跳转到华为云会议的Portal的nonce信息。获取到nonce信息后，通过链接https://bmeeting.huaweicloud.com/?lang&#x3D;zh-CN&amp;nonce&#x3D;xxxxxxxxxxxxx#/login进行免登陆跳转。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreatePortalRefNonce(request *model.CreatePortalRefNonceRequest) (*model.CreatePortalRefNonceResponse, error) {
	requestDef := GenReqDefForCreatePortalRefNonce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePortalRefNonceResponse), nil
	}
}

// 创建周期会议
//
// 管理员或UC账号可以通过该接口创建周期会议
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateRecurringMeeting(request *model.CreateRecurringMeetingRequest) (*model.CreateRecurringMeetingResponse, error) {
	requestDef := GenReqDefForCreateRecurringMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecurringMeetingResponse), nil
	}
}

// 企业管理员生成激活码
//
// 企业管理员生成智慧屏、电子白板、Ideahub的激活码
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateVisionActiveCode(request *model.CreateVisionActiveCodeRequest) (*model.CreateVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForCreateVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVisionActiveCodeResponse), nil
	}
}

// 获取websocket鉴权token
//
// 获取websocket鉴权token。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateWebSocketToken(request *model.CreateWebSocketTokenRequest) (*model.CreateWebSocketTokenResponse, error) {
	requestDef := GenReqDefForCreateWebSocketToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWebSocketTokenResponse), nil
	}
}

// 预约网络研讨会
//
// 您可根据需要预约网络研讨会。注意：暂不支持添加外部联系人作为与会嘉宾
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) CreateWebinar(request *model.CreateWebinarRequest) (*model.CreateWebinarResponse, error) {
	requestDef := GenReqDefForCreateWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWebinarResponse), nil
	}
}

// 删除与会者
//
// 删除与会者。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteAttendees(request *model.DeleteAttendeesRequest) (*model.DeleteAttendeesResponse, error) {
	requestDef := GenReqDefForDeleteAttendees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAttendeesResponse), nil
	}
}

// SP管理员删除企业
//
// 删除企业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteCorp(request *model.DeleteCorpRequest) (*model.DeleteCorpResponse, error) {
	requestDef := GenReqDefForDeleteCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCorpResponse), nil
	}
}

// 删除云会议室
//
// 企业管理员通过该接口删除企业的云会议室
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteCorpVmr(request *model.DeleteCorpVmrRequest) (*model.DeleteCorpVmrResponse, error) {
	requestDef := GenReqDefForDeleteCorpVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCorpVmrResponse), nil
	}
}

// 删除部门
//
// 企业管理员通过该接口删除部门。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteDepartment(request *model.DeleteDepartmentRequest) (*model.DeleteDepartmentResponse, error) {
	requestDef := GenReqDefForDeleteDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDepartmentResponse), nil
	}
}

// 批量删除录制
//
// 批量删除录制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteRecordings(request *model.DeleteRecordingsRequest) (*model.DeleteRecordingsResponse, error) {
	requestDef := GenReqDefForDeleteRecordings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordingsResponse), nil
	}
}

// SP管理员根据删除企业资源
//
// 企业删除资源项，删除资源项后，企业资源总数会自动减少
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteResource(request *model.DeleteResourceRequest) (*model.DeleteResourceResponse, error) {
	requestDef := GenReqDefForDeleteResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceResponse), nil
	}
}

// 企业管理员删除激活码
//
// 企业管理员批量删除激活码
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteVisionActiveCode(request *model.DeleteVisionActiveCodeRequest) (*model.DeleteVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForDeleteVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVisionActiveCodeResponse), nil
	}
}

// 删除事件订阅配置信息
//
// 管理员可以通过该接口删除事件订阅(webhook)配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteWebHookConfig(request *model.DeleteWebHookConfigRequest) (*model.DeleteWebHookConfigResponse, error) {
	requestDef := GenReqDefForDeleteWebHookConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWebHookConfigResponse), nil
	}
}

// 取消网络研讨会
//
// 您可根据需要取消网络研讨会。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DeleteWebinar(request *model.DeleteWebinarRequest) (*model.DeleteWebinarResponse, error) {
	requestDef := GenReqDefForDeleteWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWebinarResponse), nil
	}
}

// 回收云会议室
//
// 企业管理员通过该接口回收云会议室
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) DisassociateVmr(request *model.DisassociateVmrRequest) (*model.DisassociateVmrResponse, error) {
	requestDef := GenReqDefForDisassociateVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateVmrResponse), nil
	}
}

// 举手
//
// 所有来宾可以举手。来宾举手后，可以取消自己的举手。主持人可以取消所有来宾的举手。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) Hand(request *model.HandRequest) (*model.HandResponse, error) {
	requestDef := GenReqDefForHand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandResponse), nil
	}
}

// 挂断与会者
//
// 挂断正在通话中的与会者。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) HangUp(request *model.HangUpRequest) (*model.HangUpResponse, error) {
	requestDef := GenReqDefForHangUp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HangUpResponse), nil
	}
}

// 主持人邀请与会者开启、关闭摄像头
//
// 主持人邀请与会者开启、关闭摄像头
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) InviteOperateVideo(request *model.InviteOperateVideoRequest) (*model.InviteOperateVideoResponse, error) {
	requestDef := GenReqDefForInviteOperateVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteOperateVideoResponse), nil
	}
}

// 邀请与会者
//
// 邀请与会者加入会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) InviteParticipant(request *model.InviteParticipantRequest) (*model.InviteParticipantResponse, error) {
	requestDef := GenReqDefForInviteParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteParticipantResponse), nil
	}
}

// 邀请共享
//
// 场景描述：主席邀请、取消邀请会场共享 功能描述：主席邀请、取消邀请会场共享
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) InviteShare(request *model.InviteShareRequest) (*model.InviteShareResponse, error) {
	requestDef := GenReqDefForInviteShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteShareResponse), nil
	}
}

// 邀请用户
//
// 通过手机号码或者邮箱地址邀请用户加入企业。
// * 若被邀请用户在华为云会议系统中不存在，则：
//   - 华为云会议免费版和华为云会议标准版发送短信/邮件邀请用户完成注册后加入企业。用户注册成功后，加入该企业。
//   - 华为云会议旗舰版在企业内直接添加该用户。用户会收到华为云会议的初始密码，用户第一次以手机号或者邮箱登录时，需要修改密码。
// * 若被邀请用户在华为云会议系统中存在，则该用户会收到短信或者邮件确认。确认完成后改用户加入企业内。该用户的密码保持原来的密码不变。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) InviteUser(request *model.InviteUserRequest) (*model.InviteUserResponse, error) {
	requestDef := GenReqDefForInviteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteUserResponse), nil
	}
}

// 通过会议ID和密码邀请与会者
//
// 通过会议ID和密码邀请与会者
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) InviteWithPwd(request *model.InviteWithPwdRequest) (*model.InviteWithPwdResponse, error) {
	requestDef := GenReqDefForInviteWithPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InviteWithPwdResponse), nil
	}
}

// 查询历史召开的网络研讨会列表
//
// 查询历史召开的网络研讨会列表，企业管理员可查询企业内所有历史召开的网络研讨会，普通账号查询自己历史召开的网络研讨会
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ListHistoryWebinars(request *model.ListHistoryWebinarsRequest) (*model.ListHistoryWebinarsResponse, error) {
	requestDef := GenReqDefForListHistoryWebinars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryWebinarsResponse), nil
	}
}

// 查询正在召开的网络研讨会列表
//
// 查询正在召开的网络研讨会列表：企业管理员可查询企业内所有正在召开的网络研讨会，普通账号查询自己正在召开的网络研讨会
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ListOngoingWebinars(request *model.ListOngoingWebinarsRequest) (*model.ListOngoingWebinarsResponse, error) {
	requestDef := GenReqDefForListOngoingWebinars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOngoingWebinarsResponse), nil
	}
}

// 查询即将召开的网络研讨会列表
//
// 查询即将召开的网络研讨会列表：企业管理员可查询企业内所有即将召开的网络研讨会，普通账号查询自己即将召开的网络研讨会
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ListUpComingWebinars(request *model.ListUpComingWebinarsRequest) (*model.ListUpComingWebinarsResponse, error) {
	requestDef := GenReqDefForListUpComingWebinars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUpComingWebinarsResponse), nil
	}
}

// 启停会议直播
//
// 启动或停止会议直播。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) Live(request *model.LiveRequest) (*model.LiveResponse, error) {
	requestDef := GenReqDefForLive()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LiveResponse), nil
	}
}

// 锁定会议
//
// 锁定或解锁会议。锁定会议后，不允许与会者加入会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) LockMeeting(request *model.LockMeetingRequest) (*model.LockMeetingResponse, error) {
	requestDef := GenReqDefForLockMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockMeetingResponse), nil
	}
}

// 锁定会场视频源
//
// 锁定或者解锁某在线会场的视频源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) LockView(request *model.LockViewRequest) (*model.LockViewResponse, error) {
	requestDef := GenReqDefForLockView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockViewResponse), nil
	}
}

// 全场静音
//
// 主持人可以通过该接口静音/取消静音整个会议所有与会者（主持人除外）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) MuteMeeting(request *model.MuteMeetingRequest) (*model.MuteMeetingResponse, error) {
	requestDef := GenReqDefForMuteMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MuteMeetingResponse), nil
	}
}

// 静音与会者
//
// 主持人可以静音/取消静音任意与会者，来宾也可静音/取消静音自己。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) MuteParticipant(request *model.MuteParticipantRequest) (*model.MuteParticipantResponse, error) {
	requestDef := GenReqDefForMuteParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MuteParticipantResponse), nil
	}
}

// 延长会议
//
// 延长会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ProlongMeeting(request *model.ProlongMeetingRequest) (*model.ProlongMeetingResponse, error) {
	requestDef := GenReqDefForProlongMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ProlongMeetingResponse), nil
	}
}

// 启停会议录制
//
// 启动或停止会议录制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) Record(request *model.RecordRequest) (*model.RecordResponse, error) {
	requestDef := GenReqDefForRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecordResponse), nil
	}
}

// 重命名与会者
//
// 重命名某个与会者。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) RenameParticipant(request *model.RenameParticipantRequest) (*model.RenameParticipantResponse, error) {
	requestDef := GenReqDefForRenameParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenameParticipantResponse), nil
	}
}

// 企业管理员重置硬终端激活码
//
// 当硬终端激活码失效时，企业管理员可以通过该接口重置激活码，使用重新获取的激活码激活终端，每24小时可重新激活5次。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ResetActivecode(request *model.ResetActivecodeRequest) (*model.ResetActivecodeResponse, error) {
	requestDef := GenReqDefForResetActivecode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetActivecodeResponse), nil
	}
}

// 用户重置密码
//
// 该接口提供给用户重置密码功能，服务器收到请求，重新设置用户密码并返回结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ResetPwd(request *model.ResetPwdRequest) (*model.ResetPwdResponse, error) {
	requestDef := GenReqDefForResetPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdResponse), nil
	}
}

// 企业管理员重置企业成员密码
//
// 企业管理员通过该接口提供企业管理员重置企业成员密码的功能。当服务器收到重置密码的请求时，发送新的密码到企业成员的邮箱或者短信，并返回结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ResetPwdByAdmin(request *model.ResetPwdByAdminRequest) (*model.ResetPwdByAdminResponse, error) {
	requestDef := GenReqDefForResetPwdByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdByAdminResponse), nil
	}
}

// 企业管理员重置账号的激活码
//
// 企业管理员重置账号的激活码，重置后，原设备直接解绑，必须重新激活使用,若手机邮箱不填，则不会发送新的激活码
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ResetVisionActiveCode(request *model.ResetVisionActiveCodeRequest) (*model.ResetVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForResetVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetVisionActiveCodeResponse), nil
	}
}

// 点名会场
//
// 同一时间，只允许一个与会者被点名。点名会场的效果是除了主持人外，点名与会者为非静音状态，未点名的与会者统一为静音状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) RollcallParticipant(request *model.RollcallParticipantRequest) (*model.RollcallParticipantResponse, error) {
	requestDef := GenReqDefForRollcallParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollcallParticipantResponse), nil
	}
}

// 查询历史会议的与会者记录
//
// 查询指定历史会议的与会者记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchAttendanceRecordsOfHisMeeting(request *model.SearchAttendanceRecordsOfHisMeetingRequest) (*model.SearchAttendanceRecordsOfHisMeetingResponse, error) {
	requestDef := GenReqDefForSearchAttendanceRecordsOfHisMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchAttendanceRecordsOfHisMeetingResponse), nil
	}
}

// SP管理员分页搜索企业
//
// 分页搜索企业,支持名称、企业ID搜索
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchCorp(request *model.SearchCorpRequest) (*model.SearchCorpResponse, error) {
	requestDef := GenReqDefForSearchCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpResponse), nil
	}
}

// 分页查询企业管理员
//
// 通过该接口分页查询企业管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchCorpAdmins(request *model.SearchCorpAdminsRequest) (*model.SearchCorpAdminsResponse, error) {
	requestDef := GenReqDefForSearchCorpAdmins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpAdminsResponse), nil
	}
}

// 查询企业通讯录
//
// 企业用户（含管理员）通过该接口查询该企业的通讯录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchCorpDir(request *model.SearchCorpDirRequest) (*model.SearchCorpDirResponse, error) {
	requestDef := GenReqDefForSearchCorpDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpDirResponse), nil
	}
}

// 企业管理员分页查询企业资源订单列表
//
// 企业管理员根据条件查询企业资源订单列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchCorpResources(request *model.SearchCorpResourcesRequest) (*model.SearchCorpResourcesResponse, error) {
	requestDef := GenReqDefForSearchCorpResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpResourcesResponse), nil
	}
}

// 企业管理员分页查询企业云会议室
//
// 企业管理员通过该接口分页查询企业的云会议室。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchCorpVmr(request *model.SearchCorpVmrRequest) (*model.SearchCorpVmrResponse, error) {
	requestDef := GenReqDefForSearchCorpVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCorpVmrResponse), nil
	}
}

// 查询历史会议的会控记录
//
// 查询指定历史会议的会控记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchCtlRecordsOfHisMeeting(request *model.SearchCtlRecordsOfHisMeetingRequest) (*model.SearchCtlRecordsOfHisMeetingResponse, error) {
	requestDef := GenReqDefForSearchCtlRecordsOfHisMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCtlRecordsOfHisMeetingResponse), nil
	}
}

// 按名称查询所有的部门
//
// 企业管理员通过该接口按名称查询所有的部门。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchDepartmentByName(request *model.SearchDepartmentByNameRequest) (*model.SearchDepartmentByNameResponse, error) {
	requestDef := GenReqDefForSearchDepartmentByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDepartmentByNameResponse), nil
	}
}

// 分页查询终端
//
// 企业管理员通过该接口分页查询终端信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchDevices(request *model.SearchDevicesRequest) (*model.SearchDevicesResponse, error) {
	requestDef := GenReqDefForSearchDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDevicesResponse), nil
	}
}

// 查询历史会议列表
//
// 管理员可以查询管理权限域内所有的历史会议，普通用户仅能查询当前帐号管理的历史会议。不带查询参数时，默认查询权限范围内的历史会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchHisMeetings(request *model.SearchHisMeetingsRequest) (*model.SearchHisMeetingsResponse, error) {
	requestDef := GenReqDefForSearchHisMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchHisMeetingsResponse), nil
	}
}

// 分页查询信息窗素材
//
// 分页查询信息窗素材
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchMaterials(request *model.SearchMaterialsRequest) (*model.SearchMaterialsResponse, error) {
	requestDef := GenReqDefForSearchMaterials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMaterialsResponse), nil
	}
}

// 查询会议纪要列表
//
// 用户查询自己的会议纪要列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchMeetingFileList(request *model.SearchMeetingFileListRequest) (*model.SearchMeetingFileListResponse, error) {
	requestDef := GenReqDefForSearchMeetingFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMeetingFileListResponse), nil
	}
}

// 查询会议列表
//
// 管理员可以查询管理权限域内所有的会议，普通用户仅能查询当前帐号管理的会议。不带查询参数时，默认查询权限范围内正在召开或还未召开的会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchMeetings(request *model.SearchMeetingsRequest) (*model.SearchMeetingsResponse, error) {
	requestDef := GenReqDefForSearchMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMeetingsResponse), nil
	}
}

// 普通用户分页查询云会议室及个人会议ID
//
// 企业用户通过该接口查询个人已分配的云会议室及个人会议ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchMemberVmr(request *model.SearchMemberVmrRequest) (*model.SearchMemberVmrResponse, error) {
	requestDef := GenReqDefForSearchMemberVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchMemberVmrResponse), nil
	}
}

// 查询在线会议列表
//
// 管理员可以查询管理权限域内所有在线会议，普通用户仅能查询当前自己帐号管理的在线会议。不带查询参数时，默认查询权限范围内的在线会议，按开始时间升序排列。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchOnlineMeetings(request *model.SearchOnlineMeetingsRequest) (*model.SearchOnlineMeetingsResponse, error) {
	requestDef := GenReqDefForSearchOnlineMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchOnlineMeetingsResponse), nil
	}
}

// 查询信息窗节目
//
// 获取信息窗节目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchPrograms(request *model.SearchProgramsRequest) (*model.SearchProgramsResponse, error) {
	requestDef := GenReqDefForSearchPrograms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchProgramsResponse), nil
	}
}

// 查询信息窗发布
//
// 获取信息窗发布
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchPublications(request *model.SearchPublicationsRequest) (*model.SearchPublicationsResponse, error) {
	requestDef := GenReqDefForSearchPublications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchPublicationsResponse), nil
	}
}

// 查询录制列表
//
// 管理员可以查询管理权限域内所有的录制，普通用户仅能查询当前帐号管理的录制。不带查询参数时，默认查询权限范围内的录制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchRecordings(request *model.SearchRecordingsRequest) (*model.SearchRecordingsResponse, error) {
	requestDef := GenReqDefForSearchRecordings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchRecordingsResponse), nil
	}
}

// SP管理员根据分页查询企业资源
//
// SP根据条件查询企业的资源项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchResource(request *model.SearchResourceRequest) (*model.SearchResourceResponse, error) {
	requestDef := GenReqDefForSearchResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceResponse), nil
	}
}

// SP管理员根据分页查询企业资源操作记录
//
// SP根据根据条件查询企业的资源操作记录，支持根据resourceId模糊搜索
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchResourceOpRecord(request *model.SearchResourceOpRecordRequest) (*model.SearchResourceOpRecordResponse, error) {
	requestDef := GenReqDefForSearchResourceOpRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceOpRecordResponse), nil
	}
}

// 分页查询用户
//
// 企业管理员通过该接口分页查询企业用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchUsers(request *model.SearchUsersRequest) (*model.SearchUsersResponse, error) {
	requestDef := GenReqDefForSearchUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchUsersResponse), nil
	}
}

// 企业管理员分页查询激活码
//
// 企业管理员分页查询激活码，支持激活码、终端名称模糊查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchVisionActiveCode(request *model.SearchVisionActiveCodeRequest) (*model.SearchVisionActiveCodeResponse, error) {
	requestDef := GenReqDefForSearchVisionActiveCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchVisionActiveCodeResponse), nil
	}
}

// 发送滑块验证码
//
// 该接口提供发送滑块验证码。服务器收到请求，返回抠图以及抠图后的原图等结果。需要在前台界面显示出抠图以及抠图后的原图，用户通过滑块操作来匹配图形。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SendSlideVerifyCode(request *model.SendSlideVerifyCodeRequest) (*model.SendSlideVerifyCodeResponse, error) {
	requestDef := GenReqDefForSendSlideVerifyCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendSlideVerifyCodeResponse), nil
	}
}

// 发送验证码
//
// 该接口提供发送验证码，服务器收到请求，发送验证码到邮箱或者短信并返回结果。用户在前台界面通过滑块验证后，再进行发送验证码操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SendVeriCodeForChangePwd(request *model.SendVeriCodeForChangePwdRequest) (*model.SendVeriCodeForChangePwdResponse, error) {
	requestDef := GenReqDefForSendVeriCodeForChangePwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVeriCodeForChangePwdResponse), nil
	}
}

// 获取验证码
//
// 获取验证码，向手机或邮箱发送，一分钟内只会发送一次。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SendVeriCodeForUpdateUserInfo(request *model.SendVeriCodeForUpdateUserInfoRequest) (*model.SendVeriCodeForUpdateUserInfoResponse, error) {
	requestDef := GenReqDefForSendVeriCodeForUpdateUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendVeriCodeForUpdateUserInfoResponse), nil
	}
}

// 设置自定义多画面
//
// 场景描述：会议管理员在confportal手动设置多画面 功能描述：提供给会议管理员手动设置多画面的功能
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetCustomMultiPicture(request *model.SetCustomMultiPictureRequest) (*model.SetCustomMultiPictureResponse, error) {
	requestDef := GenReqDefForSetCustomMultiPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetCustomMultiPictureResponse), nil
	}
}

// 主持人选看
//
// 用于主持人轮询、主持人选看多画面、主持人选看会场操作。目前只适用于硬终端为主持人的场景。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetHostView(request *model.SetHostViewRequest) (*model.SetHostViewResponse, error) {
	requestDef := GenReqDefForSetHostView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetHostViewResponse), nil
	}
}

// 设置多画面
//
// 设置会议多画面。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetMultiPicture(request *model.SetMultiPictureRequest) (*model.SetMultiPictureResponse, error) {
	requestDef := GenReqDefForSetMultiPicture()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetMultiPictureResponse), nil
	}
}

// 会场选看
//
// 目前只适用于硬终端选看其他会场人的场景。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetParticipantView(request *model.SetParticipantViewRequest) (*model.SetParticipantViewResponse, error) {
	requestDef := GenReqDefForSetParticipantView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetParticipantViewResponse), nil
	}
}

// 申请主持人
//
// 申请或释放主持人。普通用户可申请主持人，主持人可释放主持人权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetRole(request *model.SetRoleRequest) (*model.SetRoleResponse, error) {
	requestDef := GenReqDefForSetRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRoleResponse), nil
	}
}

// 设置SSO鉴权配置
//
// 设置SSO鉴权配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetSsoConfig(request *model.SetSsoConfigRequest) (*model.SetSsoConfigResponse, error) {
	requestDef := GenReqDefForSetSsoConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSsoConfigResponse), nil
	}
}

// 设置事件订阅配置信息
//
// 设置企业事件订阅配置设置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetWebHookConfig(request *model.SetWebHookConfigRequest) (*model.SetWebHookConfigResponse, error) {
	requestDef := GenReqDefForSetWebHookConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetWebHookConfigResponse), nil
	}
}

// 通过会议ID查询企业ID
//
// 与某个会议在同一个SP下的用户，可以通过会议ID查询到该会议对应的企业ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowConfOrg(request *model.ShowConfOrgRequest) (*model.ShowConfOrgResponse, error) {
	requestDef := GenReqDefForShowConfOrg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfOrgResponse), nil
	}
}

// SP管理员查询企业
//
// 获取企业
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowCorp(request *model.ShowCorpRequest) (*model.ShowCorpResponse, error) {
	requestDef := GenReqDefForShowCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpResponse), nil
	}
}

// 查询企业管理员
//
// 通过该接口查询企业管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowCorpAdmin(request *model.ShowCorpAdminRequest) (*model.ShowCorpAdminResponse, error) {
	requestDef := GenReqDefForShowCorpAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpAdminResponse), nil
	}
}

// 企业管理员查询企业注册信息
//
// 企业管理员通过该接口查询企业注册信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowCorpBasicInfo(request *model.ShowCorpBasicInfoRequest) (*model.ShowCorpBasicInfoResponse, error) {
	requestDef := GenReqDefForShowCorpBasicInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpBasicInfoResponse), nil
	}
}

// 企业管理员查询企业内资源及业务权限
//
// 企业管理员通过该接口查询企业内资源及业务权限，包括查询已使用的资源情况。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowCorpResource(request *model.ShowCorpResourceRequest) (*model.ShowCorpResourceResponse, error) {
	requestDef := GenReqDefForShowCorpResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCorpResourceResponse), nil
	}
}

// 通过部门编码查询部门信息
//
// 通过部门编码查询部门信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowDepartment(request *model.ShowDepartmentRequest) (*model.ShowDepartmentResponse, error) {
	requestDef := GenReqDefForShowDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDepartmentResponse), nil
	}
}

// 查询部门及其一级子部门列表
//
// 企业管理员通过该接口查询部门及其一级子部门列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowDeptAndChildDept(request *model.ShowDeptAndChildDeptRequest) (*model.ShowDeptAndChildDeptResponse, error) {
	requestDef := GenReqDefForShowDeptAndChildDept()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeptAndChildDeptResponse), nil
	}
}

// 查询终端详情
//
// 企业管理员通过该接口查询终端详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowDeviceDetail(request *model.ShowDeviceDetailRequest) (*model.ShowDeviceDetailResponse, error) {
	requestDef := GenReqDefForShowDeviceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceDetailResponse), nil
	}
}

// 查询设备状态
//
// 调用本接口可以查询硬终端的状态。
// 硬终端与发起查询请求的帐号需在同一企业下，否则会鉴权失败。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowDeviceStatus(request *model.ShowDeviceStatusRequest) (*model.ShowDeviceStatusResponse, error) {
	requestDef := GenReqDefForShowDeviceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceStatusResponse), nil
	}
}

// 获取所有终端类型
//
// 企业管理员通过该接口获取所有的终端类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowDeviceTypes(request *model.ShowDeviceTypesRequest) (*model.ShowDeviceTypesResponse, error) {
	requestDef := GenReqDefForShowDeviceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeviceTypesResponse), nil
	}
}

// 查询历史会议详情
//
// 管理员可以查询管理权限域内所有的历史会议详情，普通用户仅能查询当前帐号管理的历史会议详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowHisMeetingDetail(request *model.ShowHisMeetingDetailRequest) (*model.ShowHisMeetingDetailResponse, error) {
	requestDef := GenReqDefForShowHisMeetingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHisMeetingDetailResponse), nil
	}
}

// 查询会议详情
//
// 管理员可以查询管理权限域内所有会议的详情，普通用户仅能查询当前帐号管理的会议详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowMeetingDetail(request *model.ShowMeetingDetailRequest) (*model.ShowMeetingDetailResponse, error) {
	requestDef := GenReqDefForShowMeetingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeetingDetailResponse), nil
	}
}

// 查询会议纪要详情
//
// 用户查询单个会议纪要详情（主要目的是为了得到外链）。 IdeaHub是使用fileCode来查，所以终端保持一致。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowMeetingFile(request *model.ShowMeetingFileRequest) (*model.ShowMeetingFileResponse, error) {
	requestDef := GenReqDefForShowMeetingFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeetingFileResponse), nil
	}
}

// 打开会议纪要文件列表
//
// 用户使用手机扫码后，手机端请求服务端，让服务端通知指定IdeaHub打开指定用户的会议纪要文件列表。二维码内容  cloudlink://cloudlink.huawei.com/h5page?action&#x3D;OPEN_MEETING_FILE_LIST&amp;key1&#x3D;value1&amp;key2&#x3D;value2    key/value的个数可能变化，终端解析后，在发起后续请求时，将所有key/value存为map，作为入参即可。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowMeetingFileList(request *model.ShowMeetingFileListRequest) (*model.ShowMeetingFileListResponse, error) {
	requestDef := GenReqDefForShowMeetingFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeetingFileListResponse), nil
	}
}

// 用户查询自己的信息
//
// 企业用户通过该接口查询自己的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowMyInfo(request *model.ShowMyInfoRequest) (*model.ShowMyInfoResponse, error) {
	requestDef := GenReqDefForShowMyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMyInfoResponse), nil
	}
}

// 查询在线会议详情
//
// 管理员可以查询管理权限域内所有的在线会议详情，普通用户仅能查询当前自己的帐号管理的在线会议详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowOnlineMeetingDetail(request *model.ShowOnlineMeetingDetailRequest) (*model.ShowOnlineMeetingDetailResponse, error) {
	requestDef := GenReqDefForShowOnlineMeetingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOnlineMeetingDetailResponse), nil
	}
}

// 查询企业的资源使用信息
//
// 企业管理员查询资源使用信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowOrgRes(request *model.ShowOrgResRequest) (*model.ShowOrgResResponse, error) {
	requestDef := GenReqDefForShowOrgRes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrgResResponse), nil
	}
}

// 根据ID查询节目详情
//
// 根据ID获取节目详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowProgram(request *model.ShowProgramRequest) (*model.ShowProgramResponse, error) {
	requestDef := GenReqDefForShowProgram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProgramResponse), nil
	}
}

// 根据ID查询信息窗发布详情
//
// 根据ID获取发布详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowPublication(request *model.ShowPublicationRequest) (*model.ShowPublicationResponse, error) {
	requestDef := GenReqDefForShowPublication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicationResponse), nil
	}
}

// 查询会议实时信息
//
// 查询会议实时信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowRealTimeInfoOfMeeting(request *model.ShowRealTimeInfoOfMeetingRequest) (*model.ShowRealTimeInfoOfMeetingResponse, error) {
	requestDef := GenReqDefForShowRealTimeInfoOfMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealTimeInfoOfMeetingResponse), nil
	}
}

// 查询录制详情
//
// 查询某个录制详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowRecordingDetail(request *model.ShowRecordingDetailRequest) (*model.ShowRecordingDetailResponse, error) {
	requestDef := GenReqDefForShowRecordingDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordingDetailResponse), nil
	}
}

// 查询录制文件下载链接
//
// 查询某个录制文件下载链接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowRecordingFileDownloadUrls(request *model.ShowRecordingFileDownloadUrlsRequest) (*model.ShowRecordingFileDownloadUrlsResponse, error) {
	requestDef := GenReqDefForShowRecordingFileDownloadUrls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordingFileDownloadUrlsResponse), nil
	}
}

// 查询会议所在区域信息
//
// 查询会议所在区域信息，如果会议不存在或者会议未召开，返回对应的错误码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowRegionInfoOfMeeting(request *model.ShowRegionInfoOfMeetingRequest) (*model.ShowRegionInfoOfMeetingResponse, error) {
	requestDef := GenReqDefForShowRegionInfoOfMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRegionInfoOfMeetingResponse), nil
	}
}

// 查询直播间高级设置
//
// 查询直播间高级设置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowRoomSetting(request *model.ShowRoomSettingRequest) (*model.ShowRoomSettingResponse, error) {
	requestDef := GenReqDefForShowRoomSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRoomSettingResponse), nil
	}
}

// 查询SP的共享资源使用信息
//
// SP管理查询所属SP的共享资源使用信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowSpRes(request *model.ShowSpResRequest) (*model.ShowSpResResponse, error) {
	requestDef := GenReqDefForShowSpRes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpResResponse), nil
	}
}

// SP管理员查询资源信息
//
// SP管理员查询SP的所有资源，包括已使用的资源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowSpResource(request *model.ShowSpResourceRequest) (*model.ShowSpResourceResponse, error) {
	requestDef := GenReqDefForShowSpResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSpResourceResponse), nil
	}
}

// 查询SSO鉴权配置
//
// 查询SSO鉴权配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowSsoConfig(request *model.ShowSsoConfigRequest) (*model.ShowSsoConfigResponse, error) {
	requestDef := GenReqDefForShowSsoConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSsoConfigResponse), nil
	}
}

// 查询用户详情
//
// 企业管理员通过该接口查询企业用户详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowUserDetail(request *model.ShowUserDetailRequest) (*model.ShowUserDetailResponse, error) {
	requestDef := GenReqDefForShowUserDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserDetailResponse), nil
	}
}

// 查询事件订阅配置信息
//
// 查询企业事件订阅配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowWebHookConfig(request *model.ShowWebHookConfigRequest) (*model.ShowWebHookConfigResponse, error) {
	requestDef := GenReqDefForShowWebHookConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWebHookConfigResponse), nil
	}
}

// 查询网络研讨会详情
//
// 根据conference_id查询网络研讨会详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowWebinar(request *model.ShowWebinarRequest) (*model.ShowWebinarResponse, error) {
	requestDef := GenReqDefForShowWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWebinarResponse), nil
	}
}

// 通过会议ID和密码激活会议
//
// 终端到会管进行鉴权并激活会议，先通过该接口获取会议所在Region信息，该接口需要携带会议主席密码，在会议未召开的情况下，该接口会拉起会议。如果已存在会议，则直接返回在线会议所在Region信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) StartMeeting(request *model.StartMeetingRequest) (*model.StartMeetingResponse, error) {
	requestDef := GenReqDefForStartMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartMeetingResponse), nil
	}
}

// 结束会议
//
// 结束会议。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) StopMeeting(request *model.StopMeetingRequest) (*model.StopMeetingResponse, error) {
	requestDef := GenReqDefForStopMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMeetingResponse), nil
	}
}

// 切换视频显示策略
//
// 切换视频显示策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SwitchMode(request *model.SwitchModeRequest) (*model.SwitchModeResponse, error) {
	requestDef := GenReqDefForSwitchMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchModeResponse), nil
	}
}

// 修改手机或邮箱
//
// 企业用户通过该接口修改手机或邮箱，需要先获取验证码，验证多次失败会禁止修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateContact(request *model.UpdateContactRequest) (*model.UpdateContactResponse, error) {
	requestDef := GenReqDefForUpdateContact()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateContactResponse), nil
	}
}

// SP管理员修改企业
//
// 修改企业，若任一参数为null或者不携带则不修改
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateCorp(request *model.UpdateCorpRequest) (*model.UpdateCorpResponse, error) {
	requestDef := GenReqDefForUpdateCorp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCorpResponse), nil
	}
}

// 企业管理员修改企业注册信息
//
// 企业管理员通过该接口修改企业注册信息。当前只支持修改地址。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateCorpBasicInfo(request *model.UpdateCorpBasicInfoRequest) (*model.UpdateCorpBasicInfoResponse, error) {
	requestDef := GenReqDefForUpdateCorpBasicInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCorpBasicInfoResponse), nil
	}
}

// 修改部门
//
// 企业管理员通过该接口修改部门。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateDepartment(request *model.UpdateDepartmentRequest) (*model.UpdateDepartmentResponse, error) {
	requestDef := GenReqDefForUpdateDepartment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDepartmentResponse), nil
	}
}

// 修改终端
//
// 企业管理员通过该接口修改终端。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateDevice(request *model.UpdateDeviceRequest) (*model.UpdateDeviceResponse, error) {
	requestDef := GenReqDefForUpdateDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeviceResponse), nil
	}
}

// 更新信息窗素材
//
// 更新信息窗素材
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateMaterial(request *model.UpdateMaterialRequest) (*model.UpdateMaterialResponse, error) {
	requestDef := GenReqDefForUpdateMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMaterialResponse), nil
	}
}

// 编辑预约会议
//
// 编辑预约会议。会议开始后，不能被编辑。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateMeeting(request *model.UpdateMeetingRequest) (*model.UpdateMeetingResponse, error) {
	requestDef := GenReqDefForUpdateMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMeetingResponse), nil
	}
}

// 修改用会议室及个人会议ID信息
//
// 企业用户登录后可以修改分配给用户的云会议室及个人会议ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateMemberVmr(request *model.UpdateMemberVmrRequest) (*model.UpdateMemberVmrResponse, error) {
	requestDef := GenReqDefForUpdateMemberVmr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberVmrResponse), nil
	}
}

// 用户修改自己的信息
//
// 企业用户通过该接口修改自己的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateMyInfo(request *model.UpdateMyInfoRequest) (*model.UpdateMyInfoResponse, error) {
	requestDef := GenReqDefForUpdateMyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMyInfoResponse), nil
	}
}

// 更新信息窗节目
//
// 更新信息窗节目
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateProgram(request *model.UpdateProgramRequest) (*model.UpdateProgramResponse, error) {
	requestDef := GenReqDefForUpdateProgram()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProgramResponse), nil
	}
}

// 修改信息窗发布
//
// 修改信息窗发布
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdatePublication(request *model.UpdatePublicationRequest) (*model.UpdatePublicationResponse, error) {
	requestDef := GenReqDefForUpdatePublication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicationResponse), nil
	}
}

// 修改密码
//
// 企业成员通过该接口提供用户修改密码功能，服务器收到请求，修改用户密码并返回结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdatePwd(request *model.UpdatePwdRequest) (*model.UpdatePwdResponse, error) {
	requestDef := GenReqDefForUpdatePwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePwdResponse), nil
	}
}

// 修改预定周期会议
//
// 修改预定的周期会议；会议开始时，不能修改会议
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateRecurringMeeting(request *model.UpdateRecurringMeetingRequest) (*model.UpdateRecurringMeetingResponse, error) {
	requestDef := GenReqDefForUpdateRecurringMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecurringMeetingResponse), nil
	}
}

// 修改预定周期子会议
//
// 修改预定的周期子会议；会议开始时，不能修改会议
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateRecurringSubMeeting(request *model.UpdateRecurringSubMeetingRequest) (*model.UpdateRecurringSubMeetingResponse, error) {
	requestDef := GenReqDefForUpdateRecurringSubMeeting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecurringSubMeetingResponse), nil
	}
}

// SP管理员根据修改企业资源
//
// 企业修改资源的过期时间、停用状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateResource(request *model.UpdateResourceRequest) (*model.UpdateResourceResponse, error) {
	requestDef := GenReqDefForUpdateResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceResponse), nil
	}
}

// 高级设置 - 直播间设置
//
// 保存直播间高级设置。如有部分配置信息修改，则其他未修改的原始值也需要传入，否则部分字段会替换为默认值(即：只支持全量保存)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateRoomSetting(request *model.UpdateRoomSettingRequest) (*model.UpdateRoomSettingResponse, error) {
	requestDef := GenReqDefForUpdateRoomSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoomSettingResponse), nil
	}
}

// 会中修改配置
//
// 会中修改配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateStartedConfConfig(request *model.UpdateStartedConfConfigRequest) (*model.UpdateStartedConfConfigResponse, error) {
	requestDef := GenReqDefForUpdateStartedConfConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStartedConfConfigResponse), nil
	}
}

// 刷新Token
//
// 该接口提供刷新Token功能，根据传入的Token，刷新Token失效时间并返回结果。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateToken(request *model.UpdateTokenRequest) (*model.UpdateTokenResponse, error) {
	requestDef := GenReqDefForUpdateToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTokenResponse), nil
	}
}

// 修改用户
//
// 企业管理员通过该接口修改企业用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// 变更订阅配置使用状态
//
// 变更订阅配置使用状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateWebHookConfigStatus(request *model.UpdateWebHookConfigStatusRequest) (*model.UpdateWebHookConfigStatusResponse, error) {
	requestDef := GenReqDefForUpdateWebHookConfigStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWebHookConfigStatusResponse), nil
	}
}

// 编辑网络研讨会
//
// 您可根据需要修改普通网络研讨会和周期网络研讨会。注意：暂不支持添加外部联系人作为与会嘉宾
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UpdateWebinar(request *model.UpdateWebinarRequest) (*model.UpdateWebinarResponse, error) {
	requestDef := GenReqDefForUpdateWebinar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWebinarResponse), nil
	}
}

// 开放接口 - 文件上传
//
// 文件上传的开放接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) UploadFile(request *model.UploadFileRequest) (*model.UploadFileResponse, error) {
	requestDef := GenReqDefForUploadFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFileResponse), nil
	}
}

// 查询QoS历史会议列表
//
// * 查询企业内QoS历史会议列表。
// * 支持按照时间范围查询，可查询最近3个月内数据。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchQosHistoryMeetings(request *model.SearchQosHistoryMeetingsRequest) (*model.SearchQosHistoryMeetingsResponse, error) {
	requestDef := GenReqDefForSearchQosHistoryMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosHistoryMeetingsResponse), nil
	}
}

// 查询QoS在线会议列表
//
// * 查询企业内QoS在线会议列表。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchQosOnlineMeetings(request *model.SearchQosOnlineMeetingsRequest) (*model.SearchQosOnlineMeetingsResponse, error) {
	requestDef := GenReqDefForSearchQosOnlineMeetings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosOnlineMeetingsResponse), nil
	}
}

// 查询与会者的QoS数据
//
// * 查询企业内指定与会者的QoS数据，按照音频，视频，屏幕共享，CPU分类查询QoS数据。
// * QoS数据的打点周期为5秒。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchQosParticipantDetail(request *model.SearchQosParticipantDetailRequest) (*model.SearchQosParticipantDetailResponse, error) {
	requestDef := GenReqDefForSearchQosParticipantDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosParticipantDetailResponse), nil
	}
}

// 查询QoS会议与会者列表
//
// * 查询企业内QoS会议与会者列表。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchQosParticipants(request *model.SearchQosParticipantsRequest) (*model.SearchQosParticipantsResponse, error) {
	requestDef := GenReqDefForSearchQosParticipants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchQosParticipantsResponse), nil
	}
}

// 设置企业租户指定类型的会议质量阈值
//
// * 设置企业租户指定类型的会议质量阈值。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SetQosThreshold(request *model.SetQosThresholdRequest) (*model.SetQosThresholdResponse, error) {
	requestDef := GenReqDefForSetQosThreshold()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetQosThresholdResponse), nil
	}
}

// 查询企业租户指定类型的会议质量阈值
//
// * 查询企业租户指定类型的会议质量阈值。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) ShowQosThreshold(request *model.ShowQosThresholdRequest) (*model.ShowQosThresholdResponse, error) {
	requestDef := GenReqDefForShowQosThreshold()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQosThresholdResponse), nil
	}
}

// 查询企业级会议总体统计数据
//
// * 查询企业级会议指定时间范围内总体统计数据，按日/按月统计。
// * 查询企业级会议单日内总体统计数据，按小时统计。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticConferenceInfo(request *model.SearchStatisticConferenceInfoRequest) (*model.SearchStatisticConferenceInfoResponse, error) {
	requestDef := GenReqDefForSearchStatisticConferenceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticConferenceInfoResponse), nil
	}
}

// 查询企业级会议与会统计数据
//
// * 查询企业级会议与会用户统计数据，按日/按月统计。
// * 查询企业级会议与会硬件终端统计数据，按日/按月统计。
// * 查询企业级会议与会设备统计数据，按日/按月统计。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticConferenceParticipant(request *model.SearchStatisticConferenceParticipantRequest) (*model.SearchStatisticConferenceParticipantResponse, error) {
	requestDef := GenReqDefForSearchStatisticConferenceParticipant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticConferenceParticipantResponse), nil
	}
}

// 查询企业级会议已购资源使用统计数据
//
// * 查询企业级会议的已购资源使用状况，按日/按月统计。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticResourceInfo(request *model.SearchStatisticResourceInfoRequest) (*model.SearchStatisticResourceInfoResponse, error) {
	requestDef := GenReqDefForSearchStatisticResourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticResourceInfoResponse), nil
	}
}

// 查询企业级会议的用户统计数据
//
// * 查询企业级会议用户登录数据，按日/按月统计。
// * 查询企业级会议用户激活数据，按日/按月统计。
// * 查询企业级会议用户登录设备数据，按日/按月统计。
// * 权限角色 &#x3D; 旗舰版企业/标准版企业 + 管理员。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *MeetingClient) SearchStatisticUserInfo(request *model.SearchStatisticUserInfoRequest) (*model.SearchStatisticUserInfoResponse, error) {
	requestDef := GenReqDefForSearchStatisticUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchStatisticUserInfoResponse), nil
	}
}
