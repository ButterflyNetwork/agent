package main

import (
  "fmt"
  "log"
  "github.com/buildboxhq/buildbox-agent/buildbox"
)

func main() {
  // This callback will get called every second with the
  // entire output of the command.
  callback := func(process buildbox.Process) {
    fmt.Println(process)
    fmt.Println(process.Output)
  }

  // Define the ENV variables that should be used for
  // the script
  env := []string{"BUILDBOX_COMIMT=lollies", "BUILDBOX_REPO=repotime"}

  // Run the script
  err := buildbox.RunScript("test/script.sh", env, callback)
  if err != nil {
    log.Fatal(err)
  }
}