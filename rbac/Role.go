package rbac

// Role 角色
type Role RoleMap

// RoleMap 角色授权
type RoleMap map[RoleUUID]struct {
	Right map[RightUUID]struct{} // 角色所拥有的权限
	Group map[GroupUUID]struct{} // 角色所在组
	json  string                 // 角色信息
}
