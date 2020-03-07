package rbac

// User 用户
type User UserMap

// UserMap 用户授权
type UserMap map[UserUUID]struct {
	Role  map[RoleUUID]struct{}  // 用户的角色
	Right map[RightUUID]struct{} // 用户所拥有的权限
	Group map[GroupUUID]struct{} // 用户所在组
	json  string                 // 用户信息
}
