# Envy-go

`envy-go` is a very simple package that loads one or more .env files.

## Usage

Given `.env` of:

```
TEST=hello
```

**Load:**

```go
fmt.Println(os.Getenv("TEST")) // Prints nothing
envy.Load(".env", ".env.local")
fmt.Println(os.Getenv("TEST")) // Prints "hello"
```

**MustLoad**

MustLoad fails if any files are missing.