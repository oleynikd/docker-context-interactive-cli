package main

import (
  "os"
  "fmt"
  "flag"
  "os/exec"
  "syscall"
  "strings"
  "os/signal"
  "encoding/json"
  "github.com/jwalton/gchalk"
  "github.com/AlecAivazis/survey/v2"
)

type Context struct {
  Current bool
  Description string
  DockerEndpoint string
  KubernetesEndpoint string
  ContextType string
  Name string
  StackOrchestrator string
}

func end() {
  exit(gchalk.WithYellow().Bold("Canceled"), 0)
}

func die(msg string) {
  if msg == "" {
    msg = "Error"
  }
  exit(gchalk.WithBgRed().Bold(msg), 1)
}

func exit(msg string, code int) {
  fmt.Printf(msg)
  fmt.Println()
  os.Exit(code)
}

func SetupCloseHandler() {
  c := make(chan os.Signal)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() {
    <-c
    fmt.Println("\r- Ctrl+C pressed in Terminal")
    os.Exit(0)
  }()
}

func getNamesAndCurrent(ctxs []Context) ([]string, string) {
  var names []string
  var current string
  for _, c := range ctxs {
    names = append(names, c.Name)
    if c.Current {
      current = c.Name
    }
  }
  return names, current
}

func findContextByName(ctxs []Context, name string) Context {
  for _, c := range ctxs {
    if c.Name == name {
      return c
    }
  }
  return Context{}
}

func main() {

  // Read flags
  ssh := flag.Bool("s", false, "use to ssh to docker host")
  flag.Parse()

  // Setup Ctrl+C handler
  SetupCloseHandler()

  jsonStr, err := exec.Command("docker", "context", "list", "--format", "json").Output()
  if err != nil {
    die(err.Error())
  }

  var ctxs []Context
  json.Unmarshal([]byte(jsonStr), &ctxs)

  selectedContextName := ""
  list, current := getNamesAndCurrent(ctxs)
  groupPrompt := &survey.Select{
    Message: "",
    Options: list,
    PageSize: 15,
    Default: current,
  }
  survey.AskOne(groupPrompt, &selectedContextName)

  if selectedContextName == "" {
    end()
  }

  selectedContext := findContextByName(ctxs, selectedContextName)

  if (*ssh == false) {
    exec.Command("docker", "context", "use", selectedContext.Name).Run()
  } else {
    if (strings.HasPrefix(selectedContext.DockerEndpoint, "ssh://")) {
      cmd := exec.Command("ssh", selectedContext.DockerEndpoint[6:len(selectedContext.DockerEndpoint)])
      cmd.Stdout = os.Stdout
      cmd.Stdin = os.Stdin
      cmd.Stderr = os.Stderr
      cmd.Run()
    } else {
      die("Not a remote host")
    }
  }

}