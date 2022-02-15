# loan
**loan** is a blockchain that was built using Cosmos SDK and Tendermint and scaffolded with [Starport](https://starport.com).

The module consists of:
- An `id`
- The `amount` being lent
- A `fee` for using the loan
- A `collateral` provided by the borrower to request a loan
- The `deadline` that is set on creation after which the loan can be liquidated
- The `state` of the loan which describes the status as one of the following:
  -  requested
  -  approved
  -  paid
  -  cancelled
  -  liquidated

This module requires two accounts to participate. For any given loan, there must be a `borrower` and a `lender` account, but there is nothing preventing a `borrower` for one loan from also being a `lender` for another.

A `borrower` posts a loan request with the information as follows:
- `amount`
- `fee`
- `collateral`
- `deadline`

The borrowers collateral is stored in the loan module and can be liquidated by the lender if the borrower does not pay the loan amount and fee to the lender by the deadline. The lender is the party that must approve a request from a borrower.
- After lender approves a loan, amount is transferred from lender account to borrower account
- If borrower is unable to pay back loan, lender can liquidate loan
- Liquidating the loan transfers collateral from the loan module to the lender account


## How to try it out

First you'll have to make sure that you have Starport and its dependencies [installed](https://docs.starport.network/guide/install.html)

After cloning this repo to your local machine, you can use:

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development. Adding the `-r` flag lets you reset the chain.

Your blockchain in development can be configured with `config.yml`. The config in this repo provides two user accounts with tokens to do some testing. To learn more, see the [Starport docs](https://docs.starport.com).

## Request a loan

In another terminal window while your first terminal is running `starport chain serve`, try making your first loan request from Alice's account:

```
loand tx loan request-loan 100token 2token 200token 500 --from alice 
```

If you are met with an error saying that loand doesn't exist, try changing your go PATH with the following command:

```
export PATH=$PATH:$(go env GOPATH)/bin
```

You can query for you loan to see if you were successful. 

```
loand query loan list-loan
```

Approve the loan from Bob's account and notice the state change of the loan to `approved`:
```
loand tx loan approve-loan 0 --from bob
```

Query their balances to see the loan in effect:
```
loand query bank balances <alice_or_bob_address>
```

Repay the loan from Alice's account
```
loand tx loan repay-loan 0 --from alice
```

If you don't repay the loan, you can try using the liquidate feature from Bob's account:
```
loand tx loan liquidate-loan <loan_id> -from bob
```

A loan requester can also cancel a loan request if the loan has not yet been approved by another account like in the following example:
Reset your chain
```
starport chain serve -r
```
Create a new loan request
```
loand tx loan request-loan 100token 2token 200token 100 --from bob
```
Query that your loan was requsted
```
loand query loan list-loan
```
Cancel your loan
```
loand tx loan cancel-loan 0 --from bob
```
Query to view the state change from `requested` to  `cancelled`
```
loand query loan list-loan
```

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.com/nodemadic/loan@latest! | sudo bash
```
`nodemadic/loan` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)
