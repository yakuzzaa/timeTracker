package serializers

import (
	"fmt"
	"strconv"
	"strings"
)

type Passport string
type PassportNumber string
type PassportSeries string

func (p *Passport) Validate() error {
	parts := strings.SplitN(string(*p), " ", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid passport number format")
	}

	if len(parts[0]) != 4 && len(parts[1]) != 6 {
		return fmt.Errorf("invalid passport number format")
	}

	if _, err := strconv.Atoi(parts[0]); err != nil {
		return fmt.Errorf("passport series must contain only digits")
	}

	if _, err := strconv.Atoi(parts[1]); err != nil {
		return fmt.Errorf("passport number must contain only digits")
	}

	return nil
}

func (p *Passport) Series() (string, error) {
	parts := strings.SplitN(string(*p), " ", 2)
	if len(parts[0]) != 4 {
		return "", fmt.Errorf("invalid passport series format")
	}
	return parts[0], nil
}

func (p *Passport) Number() (string, error) {
	parts := strings.SplitN(string(*p), " ", 2)
	if len(parts[1]) != 6 {
		return "", fmt.Errorf("invalid passport number format")
	}
	return parts[1], nil
}

func (p *PassportNumber) Validate() error {
	if len(*p) != 6 {
		return fmt.Errorf("invalid passport number format")
	}

	if _, err := strconv.Atoi(string(*p)); err != nil {
		return fmt.Errorf("passport number must contain only digits")
	}
	return nil
}

func (p *PassportSeries) Validate() error {
	if len(*p) != 4 {
		return fmt.Errorf("invalid passport series format")
	}

	if _, err := strconv.Atoi(string(*p)); err != nil {
		return fmt.Errorf("passport series must contain only digits")
	}

	return nil
}
