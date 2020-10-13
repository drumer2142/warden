<img alt="chi" src="https://lh3.googleusercontent.com/proxy/CwX3vKIh-CatbUrA9vi_fjqg-p_25ceAPX6HRonYZu-2niiOiNNdfe3zJh9-lD6mEBHRTENs_6ErwyVoIj-51VzgtRAX5lLOWKK_SloMVj3Xxt8KfcfjxIhGaw8cFO2FKvydSE15m4hXB7OIy_A" width="150" />

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
