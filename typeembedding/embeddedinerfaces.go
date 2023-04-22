package main

// type embedding is a way to reuse the code of an existing type
// without having to explicitly define all its methods
// in the new type

type Whisperer interface {
	Whisper() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whisperer
	Yeller
}

func (a Account) Whisper() string {
	return "I am whispering"
}

func (a Account) Yell() string {
	return "I am yelling"
}

func talk(t Talker) {
	println(t.Whisper(), t.Yell())
}

// Embedded structs below

type Account struct {
	accountId int
	balance   int
	name      string
}

type ManagerAccount struct {
	Account
}

func main() {
	// Embedded interfaces
	mgrAcct := ManagerAccount{Account{1, 100, "John"}}
	println(mgrAcct.accountId, mgrAcct.balance, mgrAcct.name)

	talk(mgrAcct)

}
