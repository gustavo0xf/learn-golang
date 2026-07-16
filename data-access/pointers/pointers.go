package main

import "fmt"

type Wallet struct {
	btcAmount int
	user      string
}

// definição de um método (golang não é orientada a objetos)
func (w *Wallet) depositBTC(qtd int) {
	w.btcAmount += qtd
}

// w Wallet = copia da struct, w *Wallet = ponteiro pra struct
func (w *Wallet) getBalance() int {
	//fmt.Printf("[!] btcAmount addr in getBalance() context: %p\n", &w.btcAmount)
	return w.btcAmount
}

func (w *Wallet) withdrawBTC(src *Wallet, dst *Wallet, qtd int) {
	if src == nil || dst == nil {
		fmt.Println("[!] one of these accounts don't exist!")
	}

	srcBalance := src.getBalance()
	if srcBalance >= qtd && src != dst {
		src.btcAmount -= qtd
		dst.btcAmount += qtd
		fmt.Printf("[TRANSFER] %s -> %s (%d BTC)\n", src.user, dst.user, qtd)
	} else {
		fmt.Println("[!] you cannot complete the transaction")
	}
}

func main() {
	ghuWallet := Wallet{0, "ghu"}
	bcraffWallet := Wallet{0, "bcraff"}

	ghuWallet.depositBTC(1000)
	bcraffWallet.depositBTC(1000)
	ghuWallet.withdrawBTC(&ghuWallet, &bcraffWallet, 900)
	fmt.Printf("[*] balances:\n[+] ghu: %d\n[+] bcraff: %d\n", ghuWallet.getBalance(), bcraffWallet.getBalance())
}
