package okex

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountCurrencies(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountCurrencies()

	fmt.Printf("%+v, %+v", ac, err)

	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

// 币种列表
func TestOKExCurrencies(t *testing.T) {

	resp, err := NewOKExClient().GetAccountCurrencies()

	if err != nil {
		t.Error(err)
	}

	x, err := GetBytes(resp)

	if err != nil {
		t.Error(err)
	}

	var currencist []Currency
	err = json.Unmarshal(x, &currencist)

	if err != nil {
		t.Error(err)
	}

	for _, v := range currencist {
		fmt.Printf("Name : %v \n", v.Name)
		fmt.Printf("Currency : %v \n", v.Currency)
		fmt.Printf("CanWithdraw : %v \n", v.CanWithdraw)
		fmt.Printf("CanDeposit : %v \n", v.CanDeposit)
		fmt.Printf("MinWithdrawal : %v \n\n\n\n", v.MinWithdrawal)
	}

}

func TestGetAccountWallet(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWallet()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountWithdrawalFeeByCurrency(t *testing.T) {
	c := NewTestClient()
	currency := "btc"
	ac, err := c.GetAccountWithdrawalFeeByCurrency(&currency)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	ac, err = c.GetAccountWithdrawalFeeByCurrency(nil)
	assert.True(t, err == nil)
	jstr, _ = Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountWithdrawalHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWithdrawalHistory()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestClient_GetAccountWithdrawalHistoryByCurrency(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWithdrawalHistoryByCurrency("btc")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountDepositAddress(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountDepositAddress("btc")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountDepositHistory(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountDepositHistory()
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountDepositHistoryByCurrency(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountDepositHistoryByCurrency("btc")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestGetAccountLeger(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountLeger(nil)
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

	optionals := NewParams()
	optionals["type"] = "37"

	ac, err = c.GetAccountLeger(&optionals)
	assert.True(t, err == nil)
	jstr, _ = Struct2JsonString(ac)
	println(jstr)

}

func TestGetAccountWalletByCurrency(t *testing.T) {
	c := NewTestClient()
	ac, err := c.GetAccountWalletByCurrency("btc")
	assert.True(t, err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)
}

func TestPostAccountWithdrawal(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostAccountWithdrawal("btc", "17DKe3kkkkiiiiTvAKKi2vMPbm1Bz3CMKw", "123456",
		4, 1, 0.0005)
	assert.True(t, ac != nil && err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

}

func TestPostAccountTransfer(t *testing.T) {
	c := NewTestClient()
	ac, err := c.PostAccountTransfer("eos", 6, 5, 0.0001, nil)
	assert.True(t, ac != nil && err == nil)
	jstr, _ := Struct2JsonString(ac)
	println(jstr)

}
