package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func init() {
	Log = log.New(os.Stderr, "[env] ", 0)
}

var (
	Log    *log.Logger
	envs   []*Env
	envMap = make(map[string]*Env)
	mu     = &sync.Mutex{}
)

type Env struct {
	Name     string
	Describe string
}

func (e *Env) String() string {
	return os.Getenv(e.Name)
}

func (e *Env) Bool() bool {
	v := strings.ToLower(os.Getenv(e.Name))
	for _, s := range []string{"true", "1"} {
		if v == s {
			return true
		}
	}
	return false
}

func (e *Env) Int() (int, error) {
	return strconv.Atoi(os.Getenv(e.Name))
}
func (e *Env) MustInt() int {
	x, err := e.Int()
	if err != nil {
		Log.Println("env", e.Name, "must be int")
	}
	return x
}

func (e *Env) Int64() (int64, error) {
	return strconv.ParseInt(os.Getenv(e.Name), 10, 64)
}
func (e *Env) MustInt64() int64 {
	x, err := e.Int64()
	if err != nil {
		Log.Println("env", e.Name, "must be int64")
	}
	return x
}

func (e *Env) Float64() (float64, error) {
	return strconv.ParseFloat(os.Getenv(e.Name), 64)
}
func (e *Env) MustFloat64() float64 {
	x, err := e.Float64()
	if err != nil {
		Log.Println("env", e.Name, "must be float64")
	}
	return x
}

func (e *Env) Duration() (time.Duration, error) {
	return time.ParseDuration(os.Getenv(e.Name))
}
func (e *Env) MustDuration() time.Duration {
	x, err := e.Duration()
	if err != nil {
		Log.Println("env", e.Name, "must be duration")
	}
	return x
}

func New(name string, desc ...string) *Env {
	if len(desc) > 0 {
		return &Env{Name: name, Describe: desc[0]}
	}
	return &Env{Name: name}
}

func Add(e *Env) {
	mu.Lock()
	defer mu.Unlock()
	if val, ok := envMap[e.Name]; !ok {
		envMap[e.Name] = e
		envs = append(envs, e)
	} else {
		Log.Printf("env defined already %s; [exist]%s; [add]%s", e.Name, val.Describe, e.Describe)
	}
}

func Get(name string) *Env {
	mu.Lock()
	defer mu.Unlock()
	return envMap[name]
}

func PrintDefaults() {
	var s string
	for _, v := range envs {
		s += fmt.Sprintf("  - %s (current \"%s\")\n", v.Name, os.Getenv(v.Name))
		if len(v.Describe) > 0 {
			s += fmt.Sprintf("    \t%s\n", v.Describe)
		}
	}
	fmt.Printf(s)
}
