# dost

dost is a CLI password manager written in Go.

Inspired by (Pass)[https://www.passwordstore.org/]

## Features

- Generate random passwords of configurable length
- Copy generated passwords to clipboard automatically 
- Skip using symbols

## Usage

```bash
> go build -o dost main.go
```

Generating password:
```
> ./dost generate email/vema@example.com
Generated Password: );XE,7-Dv?)Aa+&<{V-|pKuq5
```

Generating password with specified length (default is 25):
```
> ./dost generate email/vema@example.com 12
Generated Password: si<yJ=5/lEb3
```

Copy generated password to clipboard without printing:
```
> ./dost generate -c email/vema@example.com 
Copied to clipboard! ✅
```

Avoid symbols for generating passwords:
```
> ./dost generate -n email/vema@example.com 
Generated Password: E2UST}^{Ac[Fb&D|cD%;Eij>H
```

### Under development
- Insert new password manually
- Show existing password
- View all entries
- Password storage 
- GPG Key based encryption

### License

MIT