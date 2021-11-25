package v0

type Rule func(password string) error

var VerifyPassword = func(password string, rules []Rule) []error {
	errs := []error{}
	for _, rule := range rules {
		if err := rule(password); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}