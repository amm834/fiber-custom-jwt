# How to use this reusable JWT component

This is a simple example of how to use JWT in a golang web application with custom claims.

## Defining secret key

Define custom `SECRET_KEY` in `.env` file.

```bash
SECRET_KEY=seret
```

## Generating JWT

And then you can generate a token with the following function which will return a token string and error if any you have
to handle error in your code.

```go
    token, err := services.CreateToken(user.Id)
if err != nil {
return c.JSON(http.StatusInternalServerError, map[string]string{
"message": "error in creating token",
})
}
```

## Getting decoded JWT

`Claims()` will verify your jwt algorithm and return decoded JWT claims in map interface. You can access your custom claims 
by indexing the map which is you defined in `CreateToken()` function.

```go
    claims, err := services.Claims(token)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{
            "message": "unauthorized",
        })
    }
```
