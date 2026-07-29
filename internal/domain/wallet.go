package domain

import (
	"time"
	"://github.com"
)

type Wallet struct {
	ID uuid.UUID \"gorm:\"type:uuid;primaryKey;default:gen_random_uuid()\"\"
	UserID uuid.UUID \"gorm:\"type:uuid;uniqueIndex;not null\"\"
	Balance float64 \"gorm:\"type:numeric(15,2);not null;default:0.00\"\"
	Currency string \"gorm:\"type:varchar(10);not null;default:'USD'\"\"
	CreatedAt, UpdatedAt time.Time
}
