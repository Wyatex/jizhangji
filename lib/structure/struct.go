// @Author:  alienyan
// @Date:    2020-8-23 16:59
// Software: GoLand
package structure

type Ledger struct {
	Id 			int
	LedgerName	string
}

// session中的用户信息
type SessionUser struct {
	Uid			int		`form:"uid" json:"uid"`
	Passport	string	`form:"passport" json:"passport"`
	Nickname	string	`form:"nickname" json:"nickname"`
	LedgerNum	int		`form:"ledgernum" json:"ledgernum"`
	Ledger		interface{} `form:"ledger" json:"ledger"`
}