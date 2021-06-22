<<<<<<< HEAD
package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

// DocumentVersion !
type DocumentVersion struct {
	uadmin.Model
	Document   Document
	DocumentID uint
	File       string `uadmin:"file"`
	Number     int    `uadmin:"help:version number"`
	Date       time.Time
	Format     Format
}

// Returns the version number
func (d DocumentVersion) String() string {
	return fmt.Sprint(d.Number)
}

// HideInDashboard !
func (DocumentVersion) HideInDashboard() bool {
	return true
}
=======
package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

// DocumentVersion !
type DocumentVersion struct {
	uadmin.Model
	Document   Document
	DocumentID uint
	File       string `uadmin:"file"`
	Number     int    `uadmin:"help:version number"`
	Date       time.Time
	Format     Format
}

// Returns the version number
func (d DocumentVersion) String() string {
	return fmt.Sprint(d.Number)
}

// HideInDashboard !
func (DocumentVersion) HideInDashboard() bool {
	return true
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
