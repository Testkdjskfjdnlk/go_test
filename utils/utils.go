package utils

import(
  "log"
  "os/user"
  "errors"
)

func GetUserFromOS() (*user.User, error){
  user, err := user.Current(); if err != nil {
    return nil, errors.New("Unable to get current user from os")
  }
  return user, nil
}

func GetUsernameFromOS() (string) {
  var userName string

  user, err := GetUserFromOS(); if err != nil {
    log.Fatal(err)
    userName = "Stranger"
  } else {
    userName = user.Name
  }
  return userName
}

