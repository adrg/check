package check

// ValidateFunc represents a validation function.
type ValidateFunc func() error

// Run executes a list of validation functions and checks if any of them fail.
// Returns the first error it encounters.
func Run(vfs ...ValidateFunc) error {
	for _, vf := range vfs {
		if err := vf(); err != nil {
			return err
		}
	}

	return nil
}
