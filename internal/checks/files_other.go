//go:build !darwin

package checks

import (
	"github.com/onyz107/OnyDetect/internal/utils"
)

// CheckSystemFiles verifies the presence and validity of specific system files
// by delegating the check to the utility function utils.CheckFiles.
// It returns a boolean value indicating whether the files meet the required criteria.
func CheckSystemFiles() bool {
	return utils.CheckFiles()
}
