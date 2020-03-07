package rbac

// Group 组
type Group GroupMap

// GroupMap 组
type GroupMap map[GroupUUID]struct {
	Right map[RightUUID]struct{} // 组所拥有的权限
	json  string                 // 组信息
}
