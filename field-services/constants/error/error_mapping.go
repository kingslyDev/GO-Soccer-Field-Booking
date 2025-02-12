package error

import (
	errField "field-services/constants/error/field"
	errFieldSchedule "field-services/constants/error/fieldschedule"
	errTime "field-services/constants/error/time"
)

func ErrMapping(err error) bool {
	var (
		GeneralErrors       = GeneralErrors
		FieldErrors         = errField.FieldErrors
		FieldScheduleErrors = errFieldSchedule.FieldScheduleErrors
		TimeErrors          = errTime.TimeErrors
	)

	allErrors := make([]error, 0)
	allErrors = append(allErrors, GeneralErrors...)
	allErrors = append(allErrors, FieldErrors...)
	allErrors = append(allErrors, FieldScheduleErrors...)
	allErrors = append(allErrors, TimeErrors...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
