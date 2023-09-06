package handler

import "fmt"

type Jobs struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Remote      *bool  `json:"isRemote"`
	Link        string `json:"link"`
	Salary      int64  `json:"salary"`
}

func validateParamFormat(name, typ string) error {
	return fmt.Errorf("parameter %s (type: %s) is required", name, typ)
}

func (r *Jobs) Validate() error {
	if r.Role == "" {
		return validateParamFormat("role", "string")
	}
	if r.Company == "" {
		return validateParamFormat("company", "string")
	}
	if r.Location == "" {
		return validateParamFormat("location", "string")
	}
	if r.Description == "" {
		return validateParamFormat("description", "string")
	}
	if r.Remote == nil {
		return validateParamFormat("isRemote", "bool")
	}
	if r.Salary <= 0 {
		return validateParamFormat("salary", "int64")
	}
	return nil
}
