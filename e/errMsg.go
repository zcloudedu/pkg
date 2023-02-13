package e

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	// server
	message[OK] = "SUCCESS"
	message[UnsupportedMethod] = "不支持的方法"
	message[ServerCommonError] = "请求失败"
	message[RequestParamError] = "参数错误"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[AuthorityServerError] = "认证服务器异常"
	message[RdbError] = "缓存数据库异常"
	message[VMClusterError] = "虚机集群异常"
	// auth
	message[TokenExpireError] = "Token失效，请重新登陆"
	message[TokenGenerateError] = "生成Token失败"
	message[TokenVerifyError] = "Token验证失败"
	message[PermissionExpireError] = "权限验证失败"
	//
	message[PermissionDenied] = "用户无权限"
	message[ErrorRoleNotFound] = "角色不存在"
	message[ErrorRoleAlreadyExist] = "角色已存在"
	//
	message[ErrorRuleAlreadyExist] = "权限已存在"
	message[ErrorRuleNotFound] = "权限不存在"

	// user
	message[ErrInviteCodeNotFound] = "邀请码不存在"
	message[ErrUserAlreadyExist] = "用户已存在"
	message[ErrUserAccountExist] = "账号已存在"
	message[ErrUserNotFound] = "用户不存在"
	message[ErrUserState] = "用户被禁用"
	message[ErrUserLogin] = "用户登录失败"
	message[ErrUserEmailExist] = "注册邮箱已存在"
	message[ErrUserLoginCount] = "用户登录计次不存在"
	message[ErrUserLoginRecordNotFound] = "用户登录记录不存在"
	message[ErrOrgNotFound] = "组织不存在"
	message[ErrOrgTypeNotFound] = "组织类型不存在"
	message[ErrUserCourseAlreadyExist] = "用户课程已存在"
	//	course
	message[ErrCourseNotFound] = "课程不存在"
	message[ErrCourseTypeExist] = "课程类型已存在"
	message[ErrCourseTypeNotFound] = "课程类型不存在"
	message[ErrCourseSectionNotFound] = "课程章节不存在"
	message[ErrCourseSectionContentNotFound] = "课程章节内容不存在"
	message[ErrBoxNotFound] = "课程实验不存在"
	message[ErrBoxNoteBookNotFound] = "课程实验对应手册不存在"
	message[ErrResourceNotFound] = "课程实验资源不存在"
	message[ErrNoteBookNotFound] = "实训手册不存在"
	message[ErrSkillNotFound] = "技能点不存在"
	message[ErrCourseSkillExist] = "课程对应技能点不存在"
	//	resource
	message[ErrorFlavorNotFound] = "实例类型不存在"
	message[ErrorFlavorExist] = "实例类型已存在"
	message[ErrorImageNotFound] = "镜像不存在"
	message[ErrorImageExist] = "镜像已存在"
	message[ErrImageNotInCluster] = "集群中找不到镜像"
	message[ErrImageCreateFailed] = "镜像创建失败"
	message[ErrVmNotFound] = "找不到虚拟机"
	message[ErrVmExist] = "虚拟机已存在"
	message[ErrRebuildColdTime] = "请勿在两分钟内重复此操作"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return message[ServerCommonError]
	}
}
