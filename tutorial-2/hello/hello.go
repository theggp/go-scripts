package main

import (
    "fmt"
    "log"

    "example/greetings"
    //"strconv"
)

func main() {
  // Set properties of the predefined Logger, including
  // the log entry prefix and a flag to disable printing
  // the time, source file, and line number.
  log.SetPrefix("greetings: ")
  log.SetFlags(0)


  names := []string{"Gladys", "Samantha", "Darrin"}

  // Request a greeting message.
  messages, err := greetings.Hellos(names)
  // If an error was returned, print it to the console and
  // exit the program.
  if err != nil {
      log.Fatal(err)
  }

  // If no error was returned, print the returned map of
  // messages to the console.
  fmt.Println(messages)
  
  // testing a simple sum function
  //sum := greetings.SimpleSum(3,5)
  //sum_message := fmt.Sprintf("The sum result is %v.", strconv.Itoa(sum))
  //fmt.Println(sum_message)
}
