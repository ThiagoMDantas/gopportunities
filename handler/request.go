package handler

import "fmt"

func errParamIsRequeed(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type Opening struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *Opening) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("malformed request body or is empty")
	}
	if r.Role == "" {
		return errParamIsRequeed("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequeed("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequeed("location", "string")
	}
	if r.Remote == nil {
		return errParamIsRequeed("remote", "boolean")
	}
	if r.Link == "" {
		return errParamIsRequeed("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequeed("salary", "int64")
	}
	return nil
}

type UpdateOpening struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpening) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
