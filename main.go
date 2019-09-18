package main

import (
  "errors"
  "fmt"
  "os"
  "os/exec"
  "runtime"
  "path/filepath"
)

func main() {
  path, err := getAppPath()

  if err != nil {
    fmt.Printf("%s\n", err)
    os.Exit(1)
  }

  args := getArgs()

  if len(os.Args) < 2 {
    execApp(path, args)
  } else {
    args := getPathArgs(args, os.Args[1])
    execApp(path, args)
  }
}

func getAppPath() (string, error) {
  switch runtime.GOOS {
  case "windows":
    return os.Getenv("LOCALAPPDATA") + "\\gitkraken\\Update.exe", nil
  case "darwing":
    return "/Applications/GitKraken.app/Contents/MacOS/GitKraken", nil
  default:
    return "", errors.New("This os is not currently supported!")
  }
}

func getArgs() ([]string) {
  switch runtime.GOOS {
  case "windows":
    return []string{"--processStart", "gitkraken.exe"}
  default:
    return []string{}
  }
}

func getPathArgs(base []string, path string) ([]string) {
  path = resolvePath(path)

  switch runtime.GOOS {
  case "windows":
    return append(base, "--process-start-args", "--path \"" + path + "\"")
  default:
    return append(base, "--path", path)
  }
}

func resolvePath(path string) string {
  p, _ := filepath.Abs(path)
  return p
}

func execApp(path string, args []string) {
  cwd, _ := os.Getwd()

  cmd := exec.Command(path, args...)
  cmd.Dir = cwd
  cmd.Run()
}
