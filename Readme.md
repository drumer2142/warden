<img alt="chi" src="https://www.clipartkey.com/mpngs/m/17-172929_handcuffs-clipart-png-image-handcuffs-png.png" width="220" />

# Warden

Warden is a JWT microservice. It serves via http call as a token creator and validator.

## Routes

SignIn
```
/sign-in
```
With parameters
```json
{
    "username": "user",
    "password": "password"
}
```

SignIn
```
/auth-route
```
With parameters
```json
{
    "name": "token",
    "value": "JWT_TOKEN",
    "expires": "2020-09-09T18:07:22.67909+03:00"
}
```
