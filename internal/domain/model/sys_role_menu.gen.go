// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysRoleMenu = "sys_role_menu"

// SysRoleMenu mapped from table <sys_role_menu>
type SysRoleMenu struct {
	RoleID int64 `gorm:"column:role_id;primaryKey;comment:角色ID" json:"roleId"` // 角色ID
	MenuID int64 `gorm:"column:menu_id;primaryKey;comment:菜单ID" json:"menuId"` // 菜单ID
}

// TableName SysRoleMenu's table name
func (*SysRoleMenu) TableName() string {
	return TableNameSysRoleMenu
}
