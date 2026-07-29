package domain

import (
	"time"
	"://github.com"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID \"gorm:\"type:uuid;primaryKey;default:gen_random_uuid()\"\"
	Email string \"gorm:\"uniqueIndex;not null;type:varchar(255)\"\"
	Password string \"gorm:\"not null;type:varchar(255)\"\"
	Role string \"gorm:\"not null;default:'user';type:varchar(50)\"\"
	CreatedAt, UpdatedAt time.Time
	DeletedAt gorm.DeletedAt \"gorm:\"index\"\"
	Wallet Wallet \"gorm:\"foreignKey:UserID;constraint:OnDelete:CASCADE\"\"
}
