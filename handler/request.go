package handler

import "fmt"

// create oppening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamRequired(name, typ string) error {
	return fmt.Errorf("param: %s type %s is required", name, typ)
}

func (r *CreateOpeningRequest) Validade() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary == 0 && r.Link == "" {
		return fmt.Errorf("request body is empty")
	}
	if r.Role == "" {
		return errParamRequired("role", "string")

	}
	if r.Company == "" {
		return errParamRequired("company", "string")

	}
	if r.Location == "" {
		return errParamRequired("location", "string")

	}
	if r.Link == "" {
		return errParamRequired("link", "string")

	}
	if r.Remote == nil {
		return errParamRequired("remote", "bool")

	}
	if r.Salary <= 0 {
		return errParamRequired("salary", "int64")
	}
	return nil
}

// updateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validade() error {
	// if any field is provied, validation is true
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 || r.Link != "" {
		return nil
	}

	// if none of the fields provied, return false
	return fmt.Errorf("at least one valid must be provied")

}
