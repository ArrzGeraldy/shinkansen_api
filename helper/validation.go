package helper

import "strconv"

func UserValidation(username string, password string) bool {
	const minLength = 3
	const maxLength = 20

	usernameLen := len(username)
	passwordLen := len(password)

	return usernameLen >= minLength && usernameLen <= maxLength &&
		passwordLen >= minLength && passwordLen <= maxLength
}

func PageQueryParamsValidation(pageNumberStr string) (int,error) {
	var pageNumber int;
	var err error = nil;

	if pageNumberStr == "" {
		pageNumber = 1
	} else {
		pageNumber, err = strconv.Atoi(pageNumberStr);
	}

	return pageNumber,err;
}