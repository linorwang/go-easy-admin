/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package models

import "gorm.io/gorm"

// 菜单

type Menu struct {
	gorm.Model
	Name     string `gorm:"comment:'菜单名称';size:64" json:"name"`
	Icon     string `gorm:"comment:'菜单图标';size:64" json:"icon"`
	Path     string `gorm:"comment:'菜单访问路径';size:64" json:"path"`
	Sort     int    `gorm:"default:0;type:int(3);comment:'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)'" json:"sort"`
	ParentId uint   `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parent_id"`
	Creator  string `gorm:"comment:'创建人';size:64" json:"creator"`
	Children []Menu `gorm:"-" json:"children"`
	Roles    []Role `gorm:"many2many:relation_role_menu;" json:"roles"`
}

func (*Menu) TableName() string {
	return "menu"
}
