package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/silva4dev/microservices-codepix-project/domain/model"

	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Lucas"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}
