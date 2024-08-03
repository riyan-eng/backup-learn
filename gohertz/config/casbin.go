package config

import (
	"fmt"
	"os"

	"gohertz/infrastructure"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, err := sqladapter.NewAdapter(infrastructure.SqlDB, "postgres", "permissions")
	if err != nil {
		fmt.Printf("casbin: failed to initialize adapter - %v \n", err)
		os.Exit(1)
	}
	enforce, err := casbin.NewEnforcer("./casbin.conf", adapter)
	if err != nil {
		fmt.Printf("casbin: failed to create enforcer - %v \n", err)
		os.Exit(1)
	}

	policies := [][]string{
		{"ADMIN", "/example/*", "(GET)|(POST)|(PATCH)|(PUT)|(DELETE)"},
		{"MASYARAKAT", "/example/*", "(GET)"},
	}

	enforce.RemovePolicies(policies)
	_, err = enforce.AddPoliciesEx(policies)
	if err != nil {
		fmt.Printf("casbin: failed to add policies - %v \n", err)
		os.Exit(1)
	}
	enforce.LoadPolicy()

	return enforce
}
