package formaterror

import (
	"errors"
)

func FormatError(err string) error {
  switch err {
      case "username":
        return errors.New("User name already exist")
  }

  return errors.New("Incorrect user details")
}
