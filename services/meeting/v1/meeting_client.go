package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

// AddCorp SP管理员创建企业
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

// AddCorpInvoker SP管理员创建企业
func (c *MeetingClient) AddCorpInvoker(request *model.AddCorpRequest) *AddCorpInvoker {
	requestDef := GenReqDefForAddCorp()
	return &AddCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddCorpAdmin 添加企业管理员
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

// AddCorpAdminInvoker 添加企业管理员
func (c *MeetingClient) AddCorpAdminInvoker(request *model.AddCorpAdminRequest) *AddCorpAdminInvoker {
	requestDef := GenReqDefForAddCorpAdmin()
	return &AddCorpAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDepartment 添加部门
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

// AddDepartmentInvoker 添加部门
func (c *MeetingClient) AddDepartmentInvoker(request *model.AddDepartmentRequest) *AddDepartmentInvoker {
	requestDef := GenReqDefForAddDepartment()
	return &AddDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDevice 增加终端
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

// AddDeviceInvoker 增加终端
func (c *MeetingClient) AddDeviceInvoker(request *model.AddDeviceRequest) *AddDeviceInvoker {
	requestDef := GenReqDefForAddDevice()
	return &AddDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddMaterial 新增信息窗素材
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

// AddMaterialInvoker 新增信息窗素材
func (c *MeetingClient) AddMaterialInvoker(request *model.AddMaterialRequest) *AddMaterialInvoker {
	requestDef := GenReqDefForAddMaterial()
	return &AddMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddProgram 新增信息窗节目
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

// AddProgramInvoker 新增信息窗节目
func (c *MeetingClient) AddProgramInvoker(request *model.AddProgramRequest) *AddProgramInvoker {
	requestDef := GenReqDefForAddProgram()
	return &AddProgramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddPublication 新增信息窗发布
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

// AddPublicationInvoker 新增信息窗发布
func (c *MeetingClient) AddPublicationInvoker(request *model.AddPublicationRequest) *AddPublicationInvoker {
	requestDef := GenReqDefForAddPublication()
	return &AddPublicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddResource SP管理员分配企业资源
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

// AddResourceInvoker SP管理员分配企业资源
func (c *MeetingClient) AddResourceInvoker(request *model.AddResourceRequest) *AddResourceInvoker {
	requestDef := GenReqDefForAddResource()
	return &AddResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddToPersonalSpace 保存会议纪要到个人云空间
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

// AddToPersonalSpaceInvoker 保存会议纪要到个人云空间
func (c *MeetingClient) AddToPersonalSpaceInvoker(request *model.AddToPersonalSpaceRequest) *AddToPersonalSpaceInvoker {
	requestDef := GenReqDefForAddToPersonalSpace()
	return &AddToPersonalSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddUser 添加用户
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

// AddUserInvoker 添加用户
func (c *MeetingClient) AddUserInvoker(request *model.AddUserRequest) *AddUserInvoker {
	requestDef := GenReqDefForAddUser()
	return &AddUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AllowGuestUnmute 与会者自己解除静音
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

// AllowGuestUnmuteInvoker 与会者自己解除静音
func (c *MeetingClient) AllowGuestUnmuteInvoker(request *model.AllowGuestUnmuteRequest) *AllowGuestUnmuteInvoker {
	requestDef := GenReqDefForAllowGuestUnmute()
	return &AllowGuestUnmuteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateVmr 分配云会议室
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

// AssociateVmrInvoker 分配云会议室
func (c *MeetingClient) AssociateVmrInvoker(request *model.AssociateVmrRequest) *AssociateVmrInvoker {
	requestDef := GenReqDefForAssociateVmr()
	return &AssociateVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteCorpAdmins 批量删除企业管理员
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

// BatchDeleteCorpAdminsInvoker 批量删除企业管理员
func (c *MeetingClient) BatchDeleteCorpAdminsInvoker(request *model.BatchDeleteCorpAdminsRequest) *BatchDeleteCorpAdminsInvoker {
	requestDef := GenReqDefForBatchDeleteCorpAdmins()
	return &BatchDeleteCorpAdminsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDevices 批量删除终端
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

// BatchDeleteDevicesInvoker 批量删除终端
func (c *MeetingClient) BatchDeleteDevicesInvoker(request *model.BatchDeleteDevicesRequest) *BatchDeleteDevicesInvoker {
	requestDef := GenReqDefForBatchDeleteDevices()
	return &BatchDeleteDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMaterials 删除信息窗素材
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

// BatchDeleteMaterialsInvoker 删除信息窗素材
func (c *MeetingClient) BatchDeleteMaterialsInvoker(request *model.BatchDeleteMaterialsRequest) *BatchDeleteMaterialsInvoker {
	requestDef := GenReqDefForBatchDeleteMaterials()
	return &BatchDeleteMaterialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePrograms 删除信息窗节目
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

// BatchDeleteProgramsInvoker 删除信息窗节目
func (c *MeetingClient) BatchDeleteProgramsInvoker(request *model.BatchDeleteProgramsRequest) *BatchDeleteProgramsInvoker {
	requestDef := GenReqDefForBatchDeletePrograms()
	return &BatchDeleteProgramsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePublications 删除信息窗发布
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

// BatchDeletePublicationsInvoker 删除信息窗发布
func (c *MeetingClient) BatchDeletePublicationsInvoker(request *model.BatchDeletePublicationsRequest) *BatchDeletePublicationsInvoker {
	requestDef := GenReqDefForBatchDeletePublications()
	return &BatchDeletePublicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUsers 批量删除用户
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

// BatchDeleteUsersInvoker 批量删除用户
func (c *MeetingClient) BatchDeleteUsersInvoker(request *model.BatchDeleteUsersRequest) *BatchDeleteUsersInvoker {
	requestDef := GenReqDefForBatchDeleteUsers()
	return &BatchDeleteUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateDevicesStatus 批量修改终端状态
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

// BatchUpdateDevicesStatusInvoker 批量修改终端状态
func (c *MeetingClient) BatchUpdateDevicesStatusInvoker(request *model.BatchUpdateDevicesStatusRequest) *BatchUpdateDevicesStatusInvoker {
	requestDef := GenReqDefForBatchUpdateDevicesStatus()
	return &BatchUpdateDevicesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateUserStatus 批量修改用户状态
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

// BatchUpdateUserStatusInvoker 批量修改用户状态
func (c *MeetingClient) BatchUpdateUserStatusInvoker(request *model.BatchUpdateUserStatusRequest) *BatchUpdateUserStatusInvoker {
	requestDef := GenReqDefForBatchUpdateUserStatus()
	return &BatchUpdateUserStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BroadcastParticipant 广播会场
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

// BroadcastParticipantInvoker 广播会场
func (c *MeetingClient) BroadcastParticipantInvoker(request *model.BroadcastParticipantRequest) *BroadcastParticipantInvoker {
	requestDef := GenReqDefForBroadcastParticipant()
	return &BroadcastParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelMeeting 取消预约会议
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

// CancelMeetingInvoker 取消预约会议
func (c *MeetingClient) CancelMeetingInvoker(request *model.CancelMeetingRequest) *CancelMeetingInvoker {
	requestDef := GenReqDefForCancelMeeting()
	return &CancelMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRecurringMeeting 取消周期会议
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

// CancelRecurringMeetingInvoker 取消周期会议
func (c *MeetingClient) CancelRecurringMeetingInvoker(request *model.CancelRecurringMeetingRequest) *CancelRecurringMeetingInvoker {
	requestDef := GenReqDefForCancelRecurringMeeting()
	return &CancelRecurringMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRecurringSubMeeting 取消周期子会议
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

// CancelRecurringSubMeetingInvoker 取消周期子会议
func (c *MeetingClient) CancelRecurringSubMeetingInvoker(request *model.CancelRecurringSubMeetingRequest) *CancelRecurringSubMeetingInvoker {
	requestDef := GenReqDefForCancelRecurringSubMeeting()
	return &CancelRecurringSubMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckSlideVerifyCode 校验滑块验证码
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

// CheckSlideVerifyCodeInvoker 校验滑块验证码
func (c *MeetingClient) CheckSlideVerifyCodeInvoker(request *model.CheckSlideVerifyCodeRequest) *CheckSlideVerifyCodeInvoker {
	requestDef := GenReqDefForCheckSlideVerifyCode()
	return &CheckSlideVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckToken 校验Token
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

// CheckTokenInvoker 校验Token
func (c *MeetingClient) CheckTokenInvoker(request *model.CheckTokenRequest) *CheckTokenInvoker {
	requestDef := GenReqDefForCheckToken()
	return &CheckTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckVeriCodeForUpdateUserInfo 校验手机和邮箱对应的验证码
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

// CheckVeriCodeForUpdateUserInfoInvoker 校验手机和邮箱对应的验证码
func (c *MeetingClient) CheckVeriCodeForUpdateUserInfoInvoker(request *model.CheckVeriCodeForUpdateUserInfoRequest) *CheckVeriCodeForUpdateUserInfoInvoker {
	requestDef := GenReqDefForCheckVeriCodeForUpdateUserInfo()
	return &CheckVeriCodeForUpdateUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckVerifyCode 校验验证码
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

// CheckVerifyCodeInvoker 校验验证码
func (c *MeetingClient) CheckVerifyCodeInvoker(request *model.CheckVerifyCodeRequest) *CheckVerifyCodeInvoker {
	requestDef := GenReqDefForCheckVerifyCode()
	return &CheckVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAnonymousAuthRandom 匿名用户会议鉴权
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

// CreateAnonymousAuthRandomInvoker 匿名用户会议鉴权
func (c *MeetingClient) CreateAnonymousAuthRandomInvoker(request *model.CreateAnonymousAuthRandomRequest) *CreateAnonymousAuthRandomInvoker {
	requestDef := GenReqDefForCreateAnonymousAuthRandom()
	return &CreateAnonymousAuthRandomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfToken 获取会控Token
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

// CreateConfTokenInvoker 获取会控Token
func (c *MeetingClient) CreateConfTokenInvoker(request *model.CreateConfTokenRequest) *CreateConfTokenInvoker {
	requestDef := GenReqDefForCreateConfToken()
	return &CreateConfTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMeeting 创建会议
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

// CreateMeetingInvoker 创建会议
func (c *MeetingClient) CreateMeetingInvoker(request *model.CreateMeetingRequest) *CreateMeetingInvoker {
	requestDef := GenReqDefForCreateMeeting()
	return &CreateMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePortalRefNonce 获取页面免登陆跳转的nonce信息
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

// CreatePortalRefNonceInvoker 获取页面免登陆跳转的nonce信息
func (c *MeetingClient) CreatePortalRefNonceInvoker(request *model.CreatePortalRefNonceRequest) *CreatePortalRefNonceInvoker {
	requestDef := GenReqDefForCreatePortalRefNonce()
	return &CreatePortalRefNonceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecurringMeeting 创建周期会议
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

// CreateRecurringMeetingInvoker 创建周期会议
func (c *MeetingClient) CreateRecurringMeetingInvoker(request *model.CreateRecurringMeetingRequest) *CreateRecurringMeetingInvoker {
	requestDef := GenReqDefForCreateRecurringMeeting()
	return &CreateRecurringMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVisionActiveCode 企业管理员生成激活码
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

// CreateVisionActiveCodeInvoker 企业管理员生成激活码
func (c *MeetingClient) CreateVisionActiveCodeInvoker(request *model.CreateVisionActiveCodeRequest) *CreateVisionActiveCodeInvoker {
	requestDef := GenReqDefForCreateVisionActiveCode()
	return &CreateVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWebSocketToken 获取websocket鉴权token
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

// CreateWebSocketTokenInvoker 获取websocket鉴权token
func (c *MeetingClient) CreateWebSocketTokenInvoker(request *model.CreateWebSocketTokenRequest) *CreateWebSocketTokenInvoker {
	requestDef := GenReqDefForCreateWebSocketToken()
	return &CreateWebSocketTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWebinar 预约网络研讨会
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

// CreateWebinarInvoker 预约网络研讨会
func (c *MeetingClient) CreateWebinarInvoker(request *model.CreateWebinarRequest) *CreateWebinarInvoker {
	requestDef := GenReqDefForCreateWebinar()
	return &CreateWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAttendees 删除与会者
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

// DeleteAttendeesInvoker 删除与会者
func (c *MeetingClient) DeleteAttendeesInvoker(request *model.DeleteAttendeesRequest) *DeleteAttendeesInvoker {
	requestDef := GenReqDefForDeleteAttendees()
	return &DeleteAttendeesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCorp SP管理员删除企业
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

// DeleteCorpInvoker SP管理员删除企业
func (c *MeetingClient) DeleteCorpInvoker(request *model.DeleteCorpRequest) *DeleteCorpInvoker {
	requestDef := GenReqDefForDeleteCorp()
	return &DeleteCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCorpVmr 删除云会议室
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

// DeleteCorpVmrInvoker 删除云会议室
func (c *MeetingClient) DeleteCorpVmrInvoker(request *model.DeleteCorpVmrRequest) *DeleteCorpVmrInvoker {
	requestDef := GenReqDefForDeleteCorpVmr()
	return &DeleteCorpVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDepartment 删除部门
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

// DeleteDepartmentInvoker 删除部门
func (c *MeetingClient) DeleteDepartmentInvoker(request *model.DeleteDepartmentRequest) *DeleteDepartmentInvoker {
	requestDef := GenReqDefForDeleteDepartment()
	return &DeleteDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordings 批量删除录制
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

// DeleteRecordingsInvoker 批量删除录制
func (c *MeetingClient) DeleteRecordingsInvoker(request *model.DeleteRecordingsRequest) *DeleteRecordingsInvoker {
	requestDef := GenReqDefForDeleteRecordings()
	return &DeleteRecordingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResource SP管理员根据删除企业资源
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

// DeleteResourceInvoker SP管理员根据删除企业资源
func (c *MeetingClient) DeleteResourceInvoker(request *model.DeleteResourceRequest) *DeleteResourceInvoker {
	requestDef := GenReqDefForDeleteResource()
	return &DeleteResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVisionActiveCode 企业管理员删除激活码
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

// DeleteVisionActiveCodeInvoker 企业管理员删除激活码
func (c *MeetingClient) DeleteVisionActiveCodeInvoker(request *model.DeleteVisionActiveCodeRequest) *DeleteVisionActiveCodeInvoker {
	requestDef := GenReqDefForDeleteVisionActiveCode()
	return &DeleteVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWebHookConfig 删除事件订阅配置信息
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

// DeleteWebHookConfigInvoker 删除事件订阅配置信息
func (c *MeetingClient) DeleteWebHookConfigInvoker(request *model.DeleteWebHookConfigRequest) *DeleteWebHookConfigInvoker {
	requestDef := GenReqDefForDeleteWebHookConfig()
	return &DeleteWebHookConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWebinar 取消网络研讨会
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

// DeleteWebinarInvoker 取消网络研讨会
func (c *MeetingClient) DeleteWebinarInvoker(request *model.DeleteWebinarRequest) *DeleteWebinarInvoker {
	requestDef := GenReqDefForDeleteWebinar()
	return &DeleteWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateVmr 回收云会议室
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

// DisassociateVmrInvoker 回收云会议室
func (c *MeetingClient) DisassociateVmrInvoker(request *model.DisassociateVmrRequest) *DisassociateVmrInvoker {
	requestDef := GenReqDefForDisassociateVmr()
	return &DisassociateVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Hand 举手
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

// HandInvoker 举手
func (c *MeetingClient) HandInvoker(request *model.HandRequest) *HandInvoker {
	requestDef := GenReqDefForHand()
	return &HandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HangUp 挂断与会者
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

// HangUpInvoker 挂断与会者
func (c *MeetingClient) HangUpInvoker(request *model.HangUpRequest) *HangUpInvoker {
	requestDef := GenReqDefForHangUp()
	return &HangUpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteOperateVideo 主持人邀请与会者开启、关闭摄像头
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

// InviteOperateVideoInvoker 主持人邀请与会者开启、关闭摄像头
func (c *MeetingClient) InviteOperateVideoInvoker(request *model.InviteOperateVideoRequest) *InviteOperateVideoInvoker {
	requestDef := GenReqDefForInviteOperateVideo()
	return &InviteOperateVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteParticipant 邀请与会者
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

// InviteParticipantInvoker 邀请与会者
func (c *MeetingClient) InviteParticipantInvoker(request *model.InviteParticipantRequest) *InviteParticipantInvoker {
	requestDef := GenReqDefForInviteParticipant()
	return &InviteParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteShare 邀请共享
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

// InviteShareInvoker 邀请共享
func (c *MeetingClient) InviteShareInvoker(request *model.InviteShareRequest) *InviteShareInvoker {
	requestDef := GenReqDefForInviteShare()
	return &InviteShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteUser 邀请用户
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

// InviteUserInvoker 邀请用户
func (c *MeetingClient) InviteUserInvoker(request *model.InviteUserRequest) *InviteUserInvoker {
	requestDef := GenReqDefForInviteUser()
	return &InviteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InviteWithPwd 通过会议ID和密码邀请与会者
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

// InviteWithPwdInvoker 通过会议ID和密码邀请与会者
func (c *MeetingClient) InviteWithPwdInvoker(request *model.InviteWithPwdRequest) *InviteWithPwdInvoker {
	requestDef := GenReqDefForInviteWithPwd()
	return &InviteWithPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryWebinars 查询历史召开的网络研讨会列表
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

// ListHistoryWebinarsInvoker 查询历史召开的网络研讨会列表
func (c *MeetingClient) ListHistoryWebinarsInvoker(request *model.ListHistoryWebinarsRequest) *ListHistoryWebinarsInvoker {
	requestDef := GenReqDefForListHistoryWebinars()
	return &ListHistoryWebinarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOngoingWebinars 查询正在召开的网络研讨会列表
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

// ListOngoingWebinarsInvoker 查询正在召开的网络研讨会列表
func (c *MeetingClient) ListOngoingWebinarsInvoker(request *model.ListOngoingWebinarsRequest) *ListOngoingWebinarsInvoker {
	requestDef := GenReqDefForListOngoingWebinars()
	return &ListOngoingWebinarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUpComingWebinars 查询即将召开的网络研讨会列表
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

// ListUpComingWebinarsInvoker 查询即将召开的网络研讨会列表
func (c *MeetingClient) ListUpComingWebinarsInvoker(request *model.ListUpComingWebinarsRequest) *ListUpComingWebinarsInvoker {
	requestDef := GenReqDefForListUpComingWebinars()
	return &ListUpComingWebinarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Live 启停会议直播
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

// LiveInvoker 启停会议直播
func (c *MeetingClient) LiveInvoker(request *model.LiveRequest) *LiveInvoker {
	requestDef := GenReqDefForLive()
	return &LiveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockMeeting 锁定会议
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

// LockMeetingInvoker 锁定会议
func (c *MeetingClient) LockMeetingInvoker(request *model.LockMeetingRequest) *LockMeetingInvoker {
	requestDef := GenReqDefForLockMeeting()
	return &LockMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockView 锁定会场视频源
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

// LockViewInvoker 锁定会场视频源
func (c *MeetingClient) LockViewInvoker(request *model.LockViewRequest) *LockViewInvoker {
	requestDef := GenReqDefForLockView()
	return &LockViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MuteMeeting 全场静音
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

// MuteMeetingInvoker 全场静音
func (c *MeetingClient) MuteMeetingInvoker(request *model.MuteMeetingRequest) *MuteMeetingInvoker {
	requestDef := GenReqDefForMuteMeeting()
	return &MuteMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MuteParticipant 静音与会者
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

// MuteParticipantInvoker 静音与会者
func (c *MeetingClient) MuteParticipantInvoker(request *model.MuteParticipantRequest) *MuteParticipantInvoker {
	requestDef := GenReqDefForMuteParticipant()
	return &MuteParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ProlongMeeting 延长会议
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

// ProlongMeetingInvoker 延长会议
func (c *MeetingClient) ProlongMeetingInvoker(request *model.ProlongMeetingRequest) *ProlongMeetingInvoker {
	requestDef := GenReqDefForProlongMeeting()
	return &ProlongMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Record 启停会议录制
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

// RecordInvoker 启停会议录制
func (c *MeetingClient) RecordInvoker(request *model.RecordRequest) *RecordInvoker {
	requestDef := GenReqDefForRecord()
	return &RecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenameParticipant 重命名与会者
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

// RenameParticipantInvoker 重命名与会者
func (c *MeetingClient) RenameParticipantInvoker(request *model.RenameParticipantRequest) *RenameParticipantInvoker {
	requestDef := GenReqDefForRenameParticipant()
	return &RenameParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetActivecode 企业管理员重置硬终端激活码
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

// ResetActivecodeInvoker 企业管理员重置硬终端激活码
func (c *MeetingClient) ResetActivecodeInvoker(request *model.ResetActivecodeRequest) *ResetActivecodeInvoker {
	requestDef := GenReqDefForResetActivecode()
	return &ResetActivecodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwd 用户重置密码
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

// ResetPwdInvoker 用户重置密码
func (c *MeetingClient) ResetPwdInvoker(request *model.ResetPwdRequest) *ResetPwdInvoker {
	requestDef := GenReqDefForResetPwd()
	return &ResetPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwdByAdmin 企业管理员重置企业成员密码
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

// ResetPwdByAdminInvoker 企业管理员重置企业成员密码
func (c *MeetingClient) ResetPwdByAdminInvoker(request *model.ResetPwdByAdminRequest) *ResetPwdByAdminInvoker {
	requestDef := GenReqDefForResetPwdByAdmin()
	return &ResetPwdByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetVisionActiveCode 企业管理员重置账号的激活码
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

// ResetVisionActiveCodeInvoker 企业管理员重置账号的激活码
func (c *MeetingClient) ResetVisionActiveCodeInvoker(request *model.ResetVisionActiveCodeRequest) *ResetVisionActiveCodeInvoker {
	requestDef := GenReqDefForResetVisionActiveCode()
	return &ResetVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollcallParticipant 点名会场
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

// RollcallParticipantInvoker 点名会场
func (c *MeetingClient) RollcallParticipantInvoker(request *model.RollcallParticipantRequest) *RollcallParticipantInvoker {
	requestDef := GenReqDefForRollcallParticipant()
	return &RollcallParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchAttendanceRecordsOfHisMeeting 查询历史会议的与会者记录
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

// SearchAttendanceRecordsOfHisMeetingInvoker 查询历史会议的与会者记录
func (c *MeetingClient) SearchAttendanceRecordsOfHisMeetingInvoker(request *model.SearchAttendanceRecordsOfHisMeetingRequest) *SearchAttendanceRecordsOfHisMeetingInvoker {
	requestDef := GenReqDefForSearchAttendanceRecordsOfHisMeeting()
	return &SearchAttendanceRecordsOfHisMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorp SP管理员分页搜索企业
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

// SearchCorpInvoker SP管理员分页搜索企业
func (c *MeetingClient) SearchCorpInvoker(request *model.SearchCorpRequest) *SearchCorpInvoker {
	requestDef := GenReqDefForSearchCorp()
	return &SearchCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpAdmins 分页查询企业管理员
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

// SearchCorpAdminsInvoker 分页查询企业管理员
func (c *MeetingClient) SearchCorpAdminsInvoker(request *model.SearchCorpAdminsRequest) *SearchCorpAdminsInvoker {
	requestDef := GenReqDefForSearchCorpAdmins()
	return &SearchCorpAdminsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpDir 查询企业通讯录
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

// SearchCorpDirInvoker 查询企业通讯录
func (c *MeetingClient) SearchCorpDirInvoker(request *model.SearchCorpDirRequest) *SearchCorpDirInvoker {
	requestDef := GenReqDefForSearchCorpDir()
	return &SearchCorpDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpResources 企业管理员分页查询企业资源订单列表
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

// SearchCorpResourcesInvoker 企业管理员分页查询企业资源订单列表
func (c *MeetingClient) SearchCorpResourcesInvoker(request *model.SearchCorpResourcesRequest) *SearchCorpResourcesInvoker {
	requestDef := GenReqDefForSearchCorpResources()
	return &SearchCorpResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCorpVmr 企业管理员分页查询企业云会议室
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

// SearchCorpVmrInvoker 企业管理员分页查询企业云会议室
func (c *MeetingClient) SearchCorpVmrInvoker(request *model.SearchCorpVmrRequest) *SearchCorpVmrInvoker {
	requestDef := GenReqDefForSearchCorpVmr()
	return &SearchCorpVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCtlRecordsOfHisMeeting 查询历史会议的会控记录
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

// SearchCtlRecordsOfHisMeetingInvoker 查询历史会议的会控记录
func (c *MeetingClient) SearchCtlRecordsOfHisMeetingInvoker(request *model.SearchCtlRecordsOfHisMeetingRequest) *SearchCtlRecordsOfHisMeetingInvoker {
	requestDef := GenReqDefForSearchCtlRecordsOfHisMeeting()
	return &SearchCtlRecordsOfHisMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDepartmentByName 按名称查询所有的部门
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

// SearchDepartmentByNameInvoker 按名称查询所有的部门
func (c *MeetingClient) SearchDepartmentByNameInvoker(request *model.SearchDepartmentByNameRequest) *SearchDepartmentByNameInvoker {
	requestDef := GenReqDefForSearchDepartmentByName()
	return &SearchDepartmentByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDevices 分页查询终端
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

// SearchDevicesInvoker 分页查询终端
func (c *MeetingClient) SearchDevicesInvoker(request *model.SearchDevicesRequest) *SearchDevicesInvoker {
	requestDef := GenReqDefForSearchDevices()
	return &SearchDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchHisMeetings 查询历史会议列表
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

// SearchHisMeetingsInvoker 查询历史会议列表
func (c *MeetingClient) SearchHisMeetingsInvoker(request *model.SearchHisMeetingsRequest) *SearchHisMeetingsInvoker {
	requestDef := GenReqDefForSearchHisMeetings()
	return &SearchHisMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMaterials 分页查询信息窗素材
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

// SearchMaterialsInvoker 分页查询信息窗素材
func (c *MeetingClient) SearchMaterialsInvoker(request *model.SearchMaterialsRequest) *SearchMaterialsInvoker {
	requestDef := GenReqDefForSearchMaterials()
	return &SearchMaterialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMeetingFileList 查询会议纪要列表
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

// SearchMeetingFileListInvoker 查询会议纪要列表
func (c *MeetingClient) SearchMeetingFileListInvoker(request *model.SearchMeetingFileListRequest) *SearchMeetingFileListInvoker {
	requestDef := GenReqDefForSearchMeetingFileList()
	return &SearchMeetingFileListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMeetings 查询会议列表
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

// SearchMeetingsInvoker 查询会议列表
func (c *MeetingClient) SearchMeetingsInvoker(request *model.SearchMeetingsRequest) *SearchMeetingsInvoker {
	requestDef := GenReqDefForSearchMeetings()
	return &SearchMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchMemberVmr 普通用户分页查询云会议室及个人会议ID
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

// SearchMemberVmrInvoker 普通用户分页查询云会议室及个人会议ID
func (c *MeetingClient) SearchMemberVmrInvoker(request *model.SearchMemberVmrRequest) *SearchMemberVmrInvoker {
	requestDef := GenReqDefForSearchMemberVmr()
	return &SearchMemberVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchOnlineMeetings 查询在线会议列表
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

// SearchOnlineMeetingsInvoker 查询在线会议列表
func (c *MeetingClient) SearchOnlineMeetingsInvoker(request *model.SearchOnlineMeetingsRequest) *SearchOnlineMeetingsInvoker {
	requestDef := GenReqDefForSearchOnlineMeetings()
	return &SearchOnlineMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPrograms 查询信息窗节目
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

// SearchProgramsInvoker 查询信息窗节目
func (c *MeetingClient) SearchProgramsInvoker(request *model.SearchProgramsRequest) *SearchProgramsInvoker {
	requestDef := GenReqDefForSearchPrograms()
	return &SearchProgramsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPublications 查询信息窗发布
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

// SearchPublicationsInvoker 查询信息窗发布
func (c *MeetingClient) SearchPublicationsInvoker(request *model.SearchPublicationsRequest) *SearchPublicationsInvoker {
	requestDef := GenReqDefForSearchPublications()
	return &SearchPublicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchRecordings 查询录制列表
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

// SearchRecordingsInvoker 查询录制列表
func (c *MeetingClient) SearchRecordingsInvoker(request *model.SearchRecordingsRequest) *SearchRecordingsInvoker {
	requestDef := GenReqDefForSearchRecordings()
	return &SearchRecordingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResource SP管理员根据分页查询企业资源
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

// SearchResourceInvoker SP管理员根据分页查询企业资源
func (c *MeetingClient) SearchResourceInvoker(request *model.SearchResourceRequest) *SearchResourceInvoker {
	requestDef := GenReqDefForSearchResource()
	return &SearchResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResourceOpRecord SP管理员根据分页查询企业资源操作记录
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

// SearchResourceOpRecordInvoker SP管理员根据分页查询企业资源操作记录
func (c *MeetingClient) SearchResourceOpRecordInvoker(request *model.SearchResourceOpRecordRequest) *SearchResourceOpRecordInvoker {
	requestDef := GenReqDefForSearchResourceOpRecord()
	return &SearchResourceOpRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchUsers 分页查询用户
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

// SearchUsersInvoker 分页查询用户
func (c *MeetingClient) SearchUsersInvoker(request *model.SearchUsersRequest) *SearchUsersInvoker {
	requestDef := GenReqDefForSearchUsers()
	return &SearchUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchVisionActiveCode 企业管理员分页查询激活码
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

// SearchVisionActiveCodeInvoker 企业管理员分页查询激活码
func (c *MeetingClient) SearchVisionActiveCodeInvoker(request *model.SearchVisionActiveCodeRequest) *SearchVisionActiveCodeInvoker {
	requestDef := GenReqDefForSearchVisionActiveCode()
	return &SearchVisionActiveCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendSlideVerifyCode 发送滑块验证码
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

// SendSlideVerifyCodeInvoker 发送滑块验证码
func (c *MeetingClient) SendSlideVerifyCodeInvoker(request *model.SendSlideVerifyCodeRequest) *SendSlideVerifyCodeInvoker {
	requestDef := GenReqDefForSendSlideVerifyCode()
	return &SendSlideVerifyCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVeriCodeForChangePwd 发送验证码
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

// SendVeriCodeForChangePwdInvoker 发送验证码
func (c *MeetingClient) SendVeriCodeForChangePwdInvoker(request *model.SendVeriCodeForChangePwdRequest) *SendVeriCodeForChangePwdInvoker {
	requestDef := GenReqDefForSendVeriCodeForChangePwd()
	return &SendVeriCodeForChangePwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendVeriCodeForUpdateUserInfo 获取验证码
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

// SendVeriCodeForUpdateUserInfoInvoker 获取验证码
func (c *MeetingClient) SendVeriCodeForUpdateUserInfoInvoker(request *model.SendVeriCodeForUpdateUserInfoRequest) *SendVeriCodeForUpdateUserInfoInvoker {
	requestDef := GenReqDefForSendVeriCodeForUpdateUserInfo()
	return &SendVeriCodeForUpdateUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetCustomMultiPicture 设置自定义多画面
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

// SetCustomMultiPictureInvoker 设置自定义多画面
func (c *MeetingClient) SetCustomMultiPictureInvoker(request *model.SetCustomMultiPictureRequest) *SetCustomMultiPictureInvoker {
	requestDef := GenReqDefForSetCustomMultiPicture()
	return &SetCustomMultiPictureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetHostView 主持人选看
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

// SetHostViewInvoker 主持人选看
func (c *MeetingClient) SetHostViewInvoker(request *model.SetHostViewRequest) *SetHostViewInvoker {
	requestDef := GenReqDefForSetHostView()
	return &SetHostViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetMultiPicture 设置多画面
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

// SetMultiPictureInvoker 设置多画面
func (c *MeetingClient) SetMultiPictureInvoker(request *model.SetMultiPictureRequest) *SetMultiPictureInvoker {
	requestDef := GenReqDefForSetMultiPicture()
	return &SetMultiPictureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetParticipantView 会场选看
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

// SetParticipantViewInvoker 会场选看
func (c *MeetingClient) SetParticipantViewInvoker(request *model.SetParticipantViewRequest) *SetParticipantViewInvoker {
	requestDef := GenReqDefForSetParticipantView()
	return &SetParticipantViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRole 申请主持人
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

// SetRoleInvoker 申请主持人
func (c *MeetingClient) SetRoleInvoker(request *model.SetRoleRequest) *SetRoleInvoker {
	requestDef := GenReqDefForSetRole()
	return &SetRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSsoConfig 设置SSO鉴权配置
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

// SetSsoConfigInvoker 设置SSO鉴权配置
func (c *MeetingClient) SetSsoConfigInvoker(request *model.SetSsoConfigRequest) *SetSsoConfigInvoker {
	requestDef := GenReqDefForSetSsoConfig()
	return &SetSsoConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetWebHookConfig 设置事件订阅配置信息
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

// SetWebHookConfigInvoker 设置事件订阅配置信息
func (c *MeetingClient) SetWebHookConfigInvoker(request *model.SetWebHookConfigRequest) *SetWebHookConfigInvoker {
	requestDef := GenReqDefForSetWebHookConfig()
	return &SetWebHookConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfOrg 通过会议ID查询企业ID
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

// ShowConfOrgInvoker 通过会议ID查询企业ID
func (c *MeetingClient) ShowConfOrgInvoker(request *model.ShowConfOrgRequest) *ShowConfOrgInvoker {
	requestDef := GenReqDefForShowConfOrg()
	return &ShowConfOrgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorp SP管理员查询企业
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

// ShowCorpInvoker SP管理员查询企业
func (c *MeetingClient) ShowCorpInvoker(request *model.ShowCorpRequest) *ShowCorpInvoker {
	requestDef := GenReqDefForShowCorp()
	return &ShowCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorpAdmin 查询企业管理员
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

// ShowCorpAdminInvoker 查询企业管理员
func (c *MeetingClient) ShowCorpAdminInvoker(request *model.ShowCorpAdminRequest) *ShowCorpAdminInvoker {
	requestDef := GenReqDefForShowCorpAdmin()
	return &ShowCorpAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorpBasicInfo 企业管理员查询企业注册信息
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

// ShowCorpBasicInfoInvoker 企业管理员查询企业注册信息
func (c *MeetingClient) ShowCorpBasicInfoInvoker(request *model.ShowCorpBasicInfoRequest) *ShowCorpBasicInfoInvoker {
	requestDef := GenReqDefForShowCorpBasicInfo()
	return &ShowCorpBasicInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCorpResource 企业管理员查询企业内资源及业务权限
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

// ShowCorpResourceInvoker 企业管理员查询企业内资源及业务权限
func (c *MeetingClient) ShowCorpResourceInvoker(request *model.ShowCorpResourceRequest) *ShowCorpResourceInvoker {
	requestDef := GenReqDefForShowCorpResource()
	return &ShowCorpResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDepartment 通过部门编码查询部门信息
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

// ShowDepartmentInvoker 通过部门编码查询部门信息
func (c *MeetingClient) ShowDepartmentInvoker(request *model.ShowDepartmentRequest) *ShowDepartmentInvoker {
	requestDef := GenReqDefForShowDepartment()
	return &ShowDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeptAndChildDept 查询部门及其一级子部门列表
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

// ShowDeptAndChildDeptInvoker 查询部门及其一级子部门列表
func (c *MeetingClient) ShowDeptAndChildDeptInvoker(request *model.ShowDeptAndChildDeptRequest) *ShowDeptAndChildDeptInvoker {
	requestDef := GenReqDefForShowDeptAndChildDept()
	return &ShowDeptAndChildDeptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceDetail 查询终端详情
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

// ShowDeviceDetailInvoker 查询终端详情
func (c *MeetingClient) ShowDeviceDetailInvoker(request *model.ShowDeviceDetailRequest) *ShowDeviceDetailInvoker {
	requestDef := GenReqDefForShowDeviceDetail()
	return &ShowDeviceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceStatus 查询设备状态
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

// ShowDeviceStatusInvoker 查询设备状态
func (c *MeetingClient) ShowDeviceStatusInvoker(request *model.ShowDeviceStatusRequest) *ShowDeviceStatusInvoker {
	requestDef := GenReqDefForShowDeviceStatus()
	return &ShowDeviceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeviceTypes 获取所有终端类型
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

// ShowDeviceTypesInvoker 获取所有终端类型
func (c *MeetingClient) ShowDeviceTypesInvoker(request *model.ShowDeviceTypesRequest) *ShowDeviceTypesInvoker {
	requestDef := GenReqDefForShowDeviceTypes()
	return &ShowDeviceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHisMeetingDetail 查询历史会议详情
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

// ShowHisMeetingDetailInvoker 查询历史会议详情
func (c *MeetingClient) ShowHisMeetingDetailInvoker(request *model.ShowHisMeetingDetailRequest) *ShowHisMeetingDetailInvoker {
	requestDef := GenReqDefForShowHisMeetingDetail()
	return &ShowHisMeetingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMeetingDetail 查询会议详情
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

// ShowMeetingDetailInvoker 查询会议详情
func (c *MeetingClient) ShowMeetingDetailInvoker(request *model.ShowMeetingDetailRequest) *ShowMeetingDetailInvoker {
	requestDef := GenReqDefForShowMeetingDetail()
	return &ShowMeetingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMeetingFile 查询会议纪要详情
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

// ShowMeetingFileInvoker 查询会议纪要详情
func (c *MeetingClient) ShowMeetingFileInvoker(request *model.ShowMeetingFileRequest) *ShowMeetingFileInvoker {
	requestDef := GenReqDefForShowMeetingFile()
	return &ShowMeetingFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMeetingFileList 打开会议纪要文件列表
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

// ShowMeetingFileListInvoker 打开会议纪要文件列表
func (c *MeetingClient) ShowMeetingFileListInvoker(request *model.ShowMeetingFileListRequest) *ShowMeetingFileListInvoker {
	requestDef := GenReqDefForShowMeetingFileList()
	return &ShowMeetingFileListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMyInfo 用户查询自己的信息
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

// ShowMyInfoInvoker 用户查询自己的信息
func (c *MeetingClient) ShowMyInfoInvoker(request *model.ShowMyInfoRequest) *ShowMyInfoInvoker {
	requestDef := GenReqDefForShowMyInfo()
	return &ShowMyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOnlineMeetingDetail 查询在线会议详情
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

// ShowOnlineMeetingDetailInvoker 查询在线会议详情
func (c *MeetingClient) ShowOnlineMeetingDetailInvoker(request *model.ShowOnlineMeetingDetailRequest) *ShowOnlineMeetingDetailInvoker {
	requestDef := GenReqDefForShowOnlineMeetingDetail()
	return &ShowOnlineMeetingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrgRes 查询企业的资源使用信息
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

// ShowOrgResInvoker 查询企业的资源使用信息
func (c *MeetingClient) ShowOrgResInvoker(request *model.ShowOrgResRequest) *ShowOrgResInvoker {
	requestDef := GenReqDefForShowOrgRes()
	return &ShowOrgResInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProgram 根据ID查询节目详情
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

// ShowProgramInvoker 根据ID查询节目详情
func (c *MeetingClient) ShowProgramInvoker(request *model.ShowProgramRequest) *ShowProgramInvoker {
	requestDef := GenReqDefForShowProgram()
	return &ShowProgramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublication 根据ID查询信息窗发布详情
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

// ShowPublicationInvoker 根据ID查询信息窗发布详情
func (c *MeetingClient) ShowPublicationInvoker(request *model.ShowPublicationRequest) *ShowPublicationInvoker {
	requestDef := GenReqDefForShowPublication()
	return &ShowPublicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRealTimeInfoOfMeeting 查询会议实时信息
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

// ShowRealTimeInfoOfMeetingInvoker 查询会议实时信息
func (c *MeetingClient) ShowRealTimeInfoOfMeetingInvoker(request *model.ShowRealTimeInfoOfMeetingRequest) *ShowRealTimeInfoOfMeetingInvoker {
	requestDef := GenReqDefForShowRealTimeInfoOfMeeting()
	return &ShowRealTimeInfoOfMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordingDetail 查询录制详情
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

// ShowRecordingDetailInvoker 查询录制详情
func (c *MeetingClient) ShowRecordingDetailInvoker(request *model.ShowRecordingDetailRequest) *ShowRecordingDetailInvoker {
	requestDef := GenReqDefForShowRecordingDetail()
	return &ShowRecordingDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordingFileDownloadUrls 查询录制文件下载链接
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

// ShowRecordingFileDownloadUrlsInvoker 查询录制文件下载链接
func (c *MeetingClient) ShowRecordingFileDownloadUrlsInvoker(request *model.ShowRecordingFileDownloadUrlsRequest) *ShowRecordingFileDownloadUrlsInvoker {
	requestDef := GenReqDefForShowRecordingFileDownloadUrls()
	return &ShowRecordingFileDownloadUrlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRegionInfoOfMeeting 查询会议所在区域信息
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

// ShowRegionInfoOfMeetingInvoker 查询会议所在区域信息
func (c *MeetingClient) ShowRegionInfoOfMeetingInvoker(request *model.ShowRegionInfoOfMeetingRequest) *ShowRegionInfoOfMeetingInvoker {
	requestDef := GenReqDefForShowRegionInfoOfMeeting()
	return &ShowRegionInfoOfMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRoomSetting 查询直播间高级设置
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

// ShowRoomSettingInvoker 查询直播间高级设置
func (c *MeetingClient) ShowRoomSettingInvoker(request *model.ShowRoomSettingRequest) *ShowRoomSettingInvoker {
	requestDef := GenReqDefForShowRoomSetting()
	return &ShowRoomSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSpRes 查询SP的共享资源使用信息
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

// ShowSpResInvoker 查询SP的共享资源使用信息
func (c *MeetingClient) ShowSpResInvoker(request *model.ShowSpResRequest) *ShowSpResInvoker {
	requestDef := GenReqDefForShowSpRes()
	return &ShowSpResInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSpResource SP管理员查询资源信息
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

// ShowSpResourceInvoker SP管理员查询资源信息
func (c *MeetingClient) ShowSpResourceInvoker(request *model.ShowSpResourceRequest) *ShowSpResourceInvoker {
	requestDef := GenReqDefForShowSpResource()
	return &ShowSpResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSsoConfig 查询SSO鉴权配置
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

// ShowSsoConfigInvoker 查询SSO鉴权配置
func (c *MeetingClient) ShowSsoConfigInvoker(request *model.ShowSsoConfigRequest) *ShowSsoConfigInvoker {
	requestDef := GenReqDefForShowSsoConfig()
	return &ShowSsoConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserDetail 查询用户详情
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

// ShowUserDetailInvoker 查询用户详情
func (c *MeetingClient) ShowUserDetailInvoker(request *model.ShowUserDetailRequest) *ShowUserDetailInvoker {
	requestDef := GenReqDefForShowUserDetail()
	return &ShowUserDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWebHookConfig 查询事件订阅配置信息
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

// ShowWebHookConfigInvoker 查询事件订阅配置信息
func (c *MeetingClient) ShowWebHookConfigInvoker(request *model.ShowWebHookConfigRequest) *ShowWebHookConfigInvoker {
	requestDef := GenReqDefForShowWebHookConfig()
	return &ShowWebHookConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWebinar 查询网络研讨会详情
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

// ShowWebinarInvoker 查询网络研讨会详情
func (c *MeetingClient) ShowWebinarInvoker(request *model.ShowWebinarRequest) *ShowWebinarInvoker {
	requestDef := GenReqDefForShowWebinar()
	return &ShowWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartMeeting 通过会议ID和密码激活会议
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

// StartMeetingInvoker 通过会议ID和密码激活会议
func (c *MeetingClient) StartMeetingInvoker(request *model.StartMeetingRequest) *StartMeetingInvoker {
	requestDef := GenReqDefForStartMeeting()
	return &StartMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopMeeting 结束会议
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

// StopMeetingInvoker 结束会议
func (c *MeetingClient) StopMeetingInvoker(request *model.StopMeetingRequest) *StopMeetingInvoker {
	requestDef := GenReqDefForStopMeeting()
	return &StopMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchMode 切换视频显示策略
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

// SwitchModeInvoker 切换视频显示策略
func (c *MeetingClient) SwitchModeInvoker(request *model.SwitchModeRequest) *SwitchModeInvoker {
	requestDef := GenReqDefForSwitchMode()
	return &SwitchModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateContact 修改手机或邮箱
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

// UpdateContactInvoker 修改手机或邮箱
func (c *MeetingClient) UpdateContactInvoker(request *model.UpdateContactRequest) *UpdateContactInvoker {
	requestDef := GenReqDefForUpdateContact()
	return &UpdateContactInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCorp SP管理员修改企业
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

// UpdateCorpInvoker SP管理员修改企业
func (c *MeetingClient) UpdateCorpInvoker(request *model.UpdateCorpRequest) *UpdateCorpInvoker {
	requestDef := GenReqDefForUpdateCorp()
	return &UpdateCorpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCorpBasicInfo 企业管理员修改企业注册信息
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

// UpdateCorpBasicInfoInvoker 企业管理员修改企业注册信息
func (c *MeetingClient) UpdateCorpBasicInfoInvoker(request *model.UpdateCorpBasicInfoRequest) *UpdateCorpBasicInfoInvoker {
	requestDef := GenReqDefForUpdateCorpBasicInfo()
	return &UpdateCorpBasicInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDepartment 修改部门
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

// UpdateDepartmentInvoker 修改部门
func (c *MeetingClient) UpdateDepartmentInvoker(request *model.UpdateDepartmentRequest) *UpdateDepartmentInvoker {
	requestDef := GenReqDefForUpdateDepartment()
	return &UpdateDepartmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDevice 修改终端
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

// UpdateDeviceInvoker 修改终端
func (c *MeetingClient) UpdateDeviceInvoker(request *model.UpdateDeviceRequest) *UpdateDeviceInvoker {
	requestDef := GenReqDefForUpdateDevice()
	return &UpdateDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMaterial 更新信息窗素材
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

// UpdateMaterialInvoker 更新信息窗素材
func (c *MeetingClient) UpdateMaterialInvoker(request *model.UpdateMaterialRequest) *UpdateMaterialInvoker {
	requestDef := GenReqDefForUpdateMaterial()
	return &UpdateMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMeeting 编辑预约会议
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

// UpdateMeetingInvoker 编辑预约会议
func (c *MeetingClient) UpdateMeetingInvoker(request *model.UpdateMeetingRequest) *UpdateMeetingInvoker {
	requestDef := GenReqDefForUpdateMeeting()
	return &UpdateMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMemberVmr 修改用会议室及个人会议ID信息
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

// UpdateMemberVmrInvoker 修改用会议室及个人会议ID信息
func (c *MeetingClient) UpdateMemberVmrInvoker(request *model.UpdateMemberVmrRequest) *UpdateMemberVmrInvoker {
	requestDef := GenReqDefForUpdateMemberVmr()
	return &UpdateMemberVmrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMyInfo 用户修改自己的信息
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

// UpdateMyInfoInvoker 用户修改自己的信息
func (c *MeetingClient) UpdateMyInfoInvoker(request *model.UpdateMyInfoRequest) *UpdateMyInfoInvoker {
	requestDef := GenReqDefForUpdateMyInfo()
	return &UpdateMyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProgram 更新信息窗节目
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

// UpdateProgramInvoker 更新信息窗节目
func (c *MeetingClient) UpdateProgramInvoker(request *model.UpdateProgramRequest) *UpdateProgramInvoker {
	requestDef := GenReqDefForUpdateProgram()
	return &UpdateProgramInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublication 修改信息窗发布
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

// UpdatePublicationInvoker 修改信息窗发布
func (c *MeetingClient) UpdatePublicationInvoker(request *model.UpdatePublicationRequest) *UpdatePublicationInvoker {
	requestDef := GenReqDefForUpdatePublication()
	return &UpdatePublicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePwd 修改密码
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

// UpdatePwdInvoker 修改密码
func (c *MeetingClient) UpdatePwdInvoker(request *model.UpdatePwdRequest) *UpdatePwdInvoker {
	requestDef := GenReqDefForUpdatePwd()
	return &UpdatePwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecurringMeeting 修改预定周期会议
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

// UpdateRecurringMeetingInvoker 修改预定周期会议
func (c *MeetingClient) UpdateRecurringMeetingInvoker(request *model.UpdateRecurringMeetingRequest) *UpdateRecurringMeetingInvoker {
	requestDef := GenReqDefForUpdateRecurringMeeting()
	return &UpdateRecurringMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecurringSubMeeting 修改预定周期子会议
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

// UpdateRecurringSubMeetingInvoker 修改预定周期子会议
func (c *MeetingClient) UpdateRecurringSubMeetingInvoker(request *model.UpdateRecurringSubMeetingRequest) *UpdateRecurringSubMeetingInvoker {
	requestDef := GenReqDefForUpdateRecurringSubMeeting()
	return &UpdateRecurringSubMeetingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResource SP管理员根据修改企业资源
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

// UpdateResourceInvoker SP管理员根据修改企业资源
func (c *MeetingClient) UpdateResourceInvoker(request *model.UpdateResourceRequest) *UpdateResourceInvoker {
	requestDef := GenReqDefForUpdateResource()
	return &UpdateResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRoomSetting 高级设置 - 直播间设置
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

// UpdateRoomSettingInvoker 高级设置 - 直播间设置
func (c *MeetingClient) UpdateRoomSettingInvoker(request *model.UpdateRoomSettingRequest) *UpdateRoomSettingInvoker {
	requestDef := GenReqDefForUpdateRoomSetting()
	return &UpdateRoomSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStartedConfConfig 会中修改配置
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

// UpdateStartedConfConfigInvoker 会中修改配置
func (c *MeetingClient) UpdateStartedConfConfigInvoker(request *model.UpdateStartedConfConfigRequest) *UpdateStartedConfConfigInvoker {
	requestDef := GenReqDefForUpdateStartedConfConfig()
	return &UpdateStartedConfConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateToken 刷新Token
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

// UpdateTokenInvoker 刷新Token
func (c *MeetingClient) UpdateTokenInvoker(request *model.UpdateTokenRequest) *UpdateTokenInvoker {
	requestDef := GenReqDefForUpdateToken()
	return &UpdateTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改用户
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

// UpdateUserInvoker 修改用户
func (c *MeetingClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWebHookConfigStatus 变更订阅配置使用状态
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

// UpdateWebHookConfigStatusInvoker 变更订阅配置使用状态
func (c *MeetingClient) UpdateWebHookConfigStatusInvoker(request *model.UpdateWebHookConfigStatusRequest) *UpdateWebHookConfigStatusInvoker {
	requestDef := GenReqDefForUpdateWebHookConfigStatus()
	return &UpdateWebHookConfigStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWebinar 编辑网络研讨会
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

// UpdateWebinarInvoker 编辑网络研讨会
func (c *MeetingClient) UpdateWebinarInvoker(request *model.UpdateWebinarRequest) *UpdateWebinarInvoker {
	requestDef := GenReqDefForUpdateWebinar()
	return &UpdateWebinarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFile 开放接口 - 文件上传
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

// UploadFileInvoker 开放接口 - 文件上传
func (c *MeetingClient) UploadFileInvoker(request *model.UploadFileRequest) *UploadFileInvoker {
	requestDef := GenReqDefForUploadFile()
	return &UploadFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosHistoryMeetings 查询QoS历史会议列表
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

// SearchQosHistoryMeetingsInvoker 查询QoS历史会议列表
func (c *MeetingClient) SearchQosHistoryMeetingsInvoker(request *model.SearchQosHistoryMeetingsRequest) *SearchQosHistoryMeetingsInvoker {
	requestDef := GenReqDefForSearchQosHistoryMeetings()
	return &SearchQosHistoryMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosOnlineMeetings 查询QoS在线会议列表
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

// SearchQosOnlineMeetingsInvoker 查询QoS在线会议列表
func (c *MeetingClient) SearchQosOnlineMeetingsInvoker(request *model.SearchQosOnlineMeetingsRequest) *SearchQosOnlineMeetingsInvoker {
	requestDef := GenReqDefForSearchQosOnlineMeetings()
	return &SearchQosOnlineMeetingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosParticipantDetail 查询与会者的QoS数据
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

// SearchQosParticipantDetailInvoker 查询与会者的QoS数据
func (c *MeetingClient) SearchQosParticipantDetailInvoker(request *model.SearchQosParticipantDetailRequest) *SearchQosParticipantDetailInvoker {
	requestDef := GenReqDefForSearchQosParticipantDetail()
	return &SearchQosParticipantDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchQosParticipants 查询QoS会议与会者列表
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

// SearchQosParticipantsInvoker 查询QoS会议与会者列表
func (c *MeetingClient) SearchQosParticipantsInvoker(request *model.SearchQosParticipantsRequest) *SearchQosParticipantsInvoker {
	requestDef := GenReqDefForSearchQosParticipants()
	return &SearchQosParticipantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetQosThreshold 设置企业租户指定类型的会议质量阈值
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

// SetQosThresholdInvoker 设置企业租户指定类型的会议质量阈值
func (c *MeetingClient) SetQosThresholdInvoker(request *model.SetQosThresholdRequest) *SetQosThresholdInvoker {
	requestDef := GenReqDefForSetQosThreshold()
	return &SetQosThresholdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQosThreshold 查询企业租户指定类型的会议质量阈值
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

// ShowQosThresholdInvoker 查询企业租户指定类型的会议质量阈值
func (c *MeetingClient) ShowQosThresholdInvoker(request *model.ShowQosThresholdRequest) *ShowQosThresholdInvoker {
	requestDef := GenReqDefForShowQosThreshold()
	return &ShowQosThresholdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticConferenceInfo 查询企业级会议总体统计数据
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

// SearchStatisticConferenceInfoInvoker 查询企业级会议总体统计数据
func (c *MeetingClient) SearchStatisticConferenceInfoInvoker(request *model.SearchStatisticConferenceInfoRequest) *SearchStatisticConferenceInfoInvoker {
	requestDef := GenReqDefForSearchStatisticConferenceInfo()
	return &SearchStatisticConferenceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticConferenceParticipant 查询企业级会议与会统计数据
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

// SearchStatisticConferenceParticipantInvoker 查询企业级会议与会统计数据
func (c *MeetingClient) SearchStatisticConferenceParticipantInvoker(request *model.SearchStatisticConferenceParticipantRequest) *SearchStatisticConferenceParticipantInvoker {
	requestDef := GenReqDefForSearchStatisticConferenceParticipant()
	return &SearchStatisticConferenceParticipantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticResourceInfo 查询企业级会议已购资源使用统计数据
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

// SearchStatisticResourceInfoInvoker 查询企业级会议已购资源使用统计数据
func (c *MeetingClient) SearchStatisticResourceInfoInvoker(request *model.SearchStatisticResourceInfoRequest) *SearchStatisticResourceInfoInvoker {
	requestDef := GenReqDefForSearchStatisticResourceInfo()
	return &SearchStatisticResourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchStatisticUserInfo 查询企业级会议的用户统计数据
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

// SearchStatisticUserInfoInvoker 查询企业级会议的用户统计数据
func (c *MeetingClient) SearchStatisticUserInfoInvoker(request *model.SearchStatisticUserInfoRequest) *SearchStatisticUserInfoInvoker {
	requestDef := GenReqDefForSearchStatisticUserInfo()
	return &SearchStatisticUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
