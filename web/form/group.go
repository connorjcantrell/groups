package form

type NoGroupNameErr struct{}

func (e NoGroupNameErr) Error() string {
	return "no group name given"
}

type GroupForm struct {
	Name        string
	Description string
}

func (f GroupForm) Validate() (bool, error) {
	if f.Name == "" {
		return false, NoGroupNameErr{}
	}
	return true, nil
}
