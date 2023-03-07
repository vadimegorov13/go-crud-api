package images

import (
	"gorm.io/gorm"
)

/*
The image handlers/route will be based on Pointer receivers, for
that, it is defined as struct. This struct will receive the database
information later, so whenever image handler/route is called gives
the access to GORM.
*/

type handler struct {
	DB *gorm.DB
}
