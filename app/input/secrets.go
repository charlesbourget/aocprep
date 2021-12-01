package input

import (
	"bufio"
	"fmt"
	"github.com/keybase/go-keychain"
	"os"
	"strings"
)

const service = "aocprep"
const accessGroup = "github.com/charlesbourget/aocprep"
const account = "Advent of Code"

func Token() (token string, err error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(service)
	query.SetAccount(account)
	query.SetAccessGroup(accessGroup)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return token, err
	}
	if len(results) != 1 {
		// Token not found
		token, _ = addToken()
	} else {
		fmt.Println("Found key")
		token = string(results[0].Data)
	}

	return token, nil
}

func addToken() (string, error) {
	session, err := queryToken()
	if err != nil {
		return "", err
	}

	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetLabel("Advent of Code session key")
	item.SetAccessGroup(accessGroup)
	item.SetData([]byte(session))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	err = keychain.AddItem(item)

	if err == keychain.ErrorDuplicateItem {
		return "", err
	}

	return session, nil
}

func queryToken() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Session key not set please enter: ")
	session, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.Trim(session, "\n"), nil
}
