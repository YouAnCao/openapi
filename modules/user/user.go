package user

// LoginParam 用户信息
type LoginParam struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password"`
	VerfiyCode string `json:"verfiyCode"`
}

// SignParam 用户注册参数
type SignParam struct {
	UserName string `json:"username" binding:"required,min=6,max=12"`
	Password string `json:"password" binding:"required,min=8,max=21"`
}

// UserInfo 系统用户信息
type UserInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UserID   uint64 `json:"user_id"`
}

// UserToken 用户登录凭证
type UserToken struct {
	UserID uint64 `json:"uid"`
	Token  string `json:"token"`
}
