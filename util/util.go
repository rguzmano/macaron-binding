package util

import "fmt"

func SetMessageByClassification(classification, field string) string {
	switch classification {
	case "RequiredError":
		return fmt.Sprintf("Missing required field %s", field)
	case "EmailError":
		return fmt.Sprintf("field %s has an invalid format", field)
	case "RangeError":
		return fmt.Sprintf("field %s is not in a valid range", field)
	case "GmailValidation":
		return fmt.Sprintf("field %s is not a gmail", field)
	default:
		return fmt.Sprintf("Bad value in field %s", field)
	}
}