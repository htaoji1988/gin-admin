package models

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(30);unique"`
	Email     string `gorm:"type:varchar(100);not null;default:''"`
	Name      string `gorm:"type:varchar(30);not null;default:''"`
	Password  string `gorm:"type:varchar(64);not null;default:''"`
	Nickname  string `gorm:"type:varchar(30)"`
	Sex       string `gorm:"type:varchar(2);"`
	Active    bool   `gorm:"not null;default:'1'"`
	Superuser bool   `gorm:"not null;default:'1'"`
	Role      Role   `gorm:"foreignkey:RoleID;association_foreignkey:ID"`
	RoleID    uint   `json:"role_id,omitempty"`
}

func (User) TableName() string {
	return "usermanage_users"
}

type Role struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(30);"`
	Permission []Urls `gorm:"many2many:usermanage_role_urls;"`
}

func (Role) TableName() string {
	return "usermanage_roles"
}
