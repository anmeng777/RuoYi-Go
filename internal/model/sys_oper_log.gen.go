// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysOperLog = "sys_oper_log"

// SysOperLog mapped from table <sys_oper_log>
type SysOperLog struct {
	OperID        int64     `gorm:"column:oper_id;primaryKey;comment:日志主键" json:"oper_id"`                    // 日志主键
	Title         string    `gorm:"column:title;comment:模块标题" json:"title"`                                   // 模块标题
	BusinessType  int32     `gorm:"column:business_type;comment:业务类型（0其它 1新增 2修改 3删除）" json:"business_type"`  // 业务类型（0其它 1新增 2修改 3删除）
	Method        string    `gorm:"column:method;comment:方法名称" json:"method"`                                 // 方法名称
	RequestMethod string    `gorm:"column:request_method;comment:请求方式" json:"request_method"`                 // 请求方式
	OperatorType  int32     `gorm:"column:operator_type;comment:操作类别（0其它 1后台用户 2手机端用户）" json:"operator_type"` // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string    `gorm:"column:oper_name;comment:操作人员" json:"oper_name"`                           // 操作人员
	DeptName      string    `gorm:"column:dept_name;comment:部门名称" json:"dept_name"`                           // 部门名称
	OperURL       string    `gorm:"column:oper_url;comment:请求URL" json:"oper_url"`                            // 请求URL
	OperIP        string    `gorm:"column:oper_ip;comment:主机地址" json:"oper_ip"`                               // 主机地址
	OperLocation  string    `gorm:"column:oper_location;comment:操作地点" json:"oper_location"`                   // 操作地点
	OperParam     string    `gorm:"column:oper_param;comment:请求参数" json:"oper_param"`                         // 请求参数
	JSONResult    string    `gorm:"column:json_result;comment:返回参数" json:"json_result"`                       // 返回参数
	Status        int32     `gorm:"column:status;comment:操作状态（0正常 1异常）" json:"status"`                        // 操作状态（0正常 1异常）
	ErrorMsg      string    `gorm:"column:error_msg;comment:错误消息" json:"error_msg"`                           // 错误消息
	OperTime      time.Time `gorm:"column:oper_time;comment:操作时间" json:"oper_time"`                           // 操作时间
	CostTime      int64     `gorm:"column:cost_time;comment:消耗时间" json:"cost_time"`                           // 消耗时间
}

// TableName SysOperLog's table name
func (*SysOperLog) TableName() string {
	return TableNameSysOperLog
}