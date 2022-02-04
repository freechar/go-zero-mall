| `api` 服务 | 端口：8000 | `rpc` 服务 | 端口：9000 |
| -------- | ------- | -------- | ------- |
| login    | 用户登录接口  | login    | 用户登录接口  |
| register | 用户注册接口  | register | 用户注册接口  |
| userinfo | 用户信息接口  | userinfo | 用户信息接口  |

# api

## 用户登录

```go
    LoginRequest {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	LoginResponse {
		AccessToken  string `json:"accessToken"`
		AccessExpire int64  `json:"accessExpire"`
	}


```

## 用户注册

```go
    RegisterRequest {
		Name     string `json:"ame"`
		Gender   int64  `json:"gender"`
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	RegisterResponse {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Gender int64  `json:"gender"`
		Mobile string `json:"mobile"`
	}


```

## 用户信息

```go
serInfoResponse {
		Id     int64  `json:"id"`
		Name   string `json:"name"`
		Gender int64  `json:"gender"`
		Mobile string `json:"mobile"`
	}

```


