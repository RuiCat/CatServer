package rbac

// Permission 授权
//+ 储存用户授权信息
type Permission []*UUID

// PermissionMap 授权信息列表
//+ 储存用户会话许可状态
//+   UUID: 当前会话密钥
type PermissionMap map[UUID]struct {
	User       UserUUID   // 绑定的用户对象 ID
	Permission Permission // 当前 Key 授权权限信息
}
