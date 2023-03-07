package users

import (
	"gorm.io/gorm"
)

/*
The user handlers/route will be based on Pointer receivers, for
that, it is defined as struct. This struct will receive the database
information later, so whenever user handler/route is called gives
the access to GORM.
*/

type handler struct {
	DB *gorm.DB
}
