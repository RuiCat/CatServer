package rbac

// Right 权限
type Right interface{}

// RBAC 权限管理
type RBAC struct {
	// 权限定义
	//+ 权限管理服务提供
	User  map[UserUUID]User   // 用户列表
	Role  map[RoleUUID]Role   // 角色列表
	Group map[GroupUUID]Group // 组列表
	Right map[RightUUID]Right // 权限列表
	// 授权信息
	//+ 授权服务提供
	Permission PermissionMap
}
