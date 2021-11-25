package v0

type Result struct {
	Passed bool
	Reason string
}

type Rule func(password string) Result

var VerifyPassword = func(password string, rules []Rule) []string {
	errs := []string{}
	for _, rule := range rules {
		result := rule(password)
		if !result.Passed {
			errs = append(errs, "error "+result.Reason)
		}
	}
	return errs
}