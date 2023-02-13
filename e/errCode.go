package e

// 全局错误码
/*
	0-1000 Server Code
	1000+ auth module Code
	2000+ user module Code
	3000+ info module Code
	4000+ course module Code
	5000+ resource module Code
*/
const (
	OK                   uint32 = 200
	PermissionDenied     uint32 = 401
	UnsupportedMethod    uint32 = 404
	ServerCommonError    uint32 = 500
	RequestParamError    uint32 = 503
	DbError              uint32 = 510
	AuthorityServerError uint32 = 511 // Cbs Error
	RdbError             uint32 = 512 // Redis Error
	VMClusterError       uint32 = 512 // Kubernetes Error
	// auth
	TokenExpireError      uint32 = 1001
	TokenGenerateError    uint32 = 1002
	TokenVerifyError      uint32 = 1003
	PermissionExpireError uint32 = 1004
	ErrorRoleNotFound     uint32 = 1005
	ErrorRoleAlreadyExist uint32 = 1007
	ErrorRuleAlreadyExist uint32 = 1008
	ErrorRuleNotFound     uint32 = 1009
	// user
	ErrInviteCodeNotFound      uint32 = 2001
	ErrUserAlreadyExist        uint32 = 2002
	ErrUserAccountExist        uint32 = 2003
	ErrUserNotFound            uint32 = 2004
	ErrUserState               uint32 = 2005
	ErrUserLogin               uint32 = 2006
	ErrUserEmailExist          uint32 = 2007
	ErrUserLoginCount          uint32 = 2008
	ErrUserLoginRecordNotFound uint32 = 2009
	ErrOrgNotFound             uint32 = 2010
	ErrOrgTypeNotFound         uint32 = 2011
	ErrUserCourseAlreadyExist  uint32 = 2012

	// course
	ErrCourseTypeExist              uint32 = 4001
	ErrCourseTypeNotFound           uint32 = 4002
	ErrCourseNotFound               uint32 = 4003
	ErrCourseSectionNotFound        uint32 = 4004
	ErrCourseSectionContentNotFound uint32 = 4005
	ErrBoxNotFound                  uint32 = 4006
	ErrBoxNoteBookNotFound          uint32 = 4007
	ErrResourceNotFound             uint32 = 4008
	ErrNoteBookNotFound             uint32 = 4009
	ErrSkillNotFound                uint32 = 4010
	ErrCourseSkillExist             uint32 = 4011
	//	resource
	ErrorFlavorNotFound  uint32 = 5001
	ErrorFlavorExist     uint32 = 5002
	ErrorImageNotFound   uint32 = 5003
	ErrorImageExist      uint32 = 5004
	ErrImageNotInCluster uint32 = 5005
	ErrImageCreateFailed uint32 = 5006
	ErrVmNotFound        uint32 = 5007
	ErrVmExist           uint32 = 5008
	ErrRebuildColdTime   uint32 = 5009
)
