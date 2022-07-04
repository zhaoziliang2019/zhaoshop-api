package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
	"time"
	"zhaoshop-api/user-web/forms"
	"zhaoshop-api/user-web/global"
	"zhaoshop-api/user-web/global/response"
	"zhaoshop-api/user-web/middlewares"
	"zhaoshop-api/user-web/models"
	"zhaoshop-api/user-web/proto"
	"zhaoshop-api/user-web/utils"
)

//将grpc的code转换为http状态码
func HandleGrapcErrorToHttp(err error, context2 *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				context2.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				context2.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				context2.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				context2.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				context2.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
			return
		}
	}
}

//获取用户列表
func GetUserList(ctx *gin.Context) {

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("pSize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{Pn: uint32(pnInt), PSize: uint32(pSizeInt)})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询用户列表失败")
		HandleGrapcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			//BirthDay: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			BirthDay: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Moblie:   value.Mobile,
			Gender:   value.Gender,
		}
		result = append(result, user)
		//data := make(map[string]interface{})
		//data["id"] = value.Id
		//data["gender"] = value.Gender
		//data["name"] = value.NickName
		//data["moblie"] = value.Mobile
		//data["birthday"] = value.BirthDay
		//result = append(result, data)
	}
	ctx.JSON(http.StatusOK, result)
}

//登录
func PassWordLogin(ctx *gin.Context) {
	passWordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passWordLoginForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	//图片验证码验证
	if !store.Verify(passWordLoginForm.CaptchaId, passWordLoginForm.Captcha, true) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}
	//从注册中心获取用户服务信息
	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: passWordLoginForm.Mobile}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "登录失败",
				})
			}
			return
		}
	} else {
		//只是查询到用户，并没有坚持密码
		if passRsp, err := global.UserSrvClient.CheckPassWord(context.Background(), &proto.PassWordCheckInfo{
			Password:          passWordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "登录失败",
			})
		} else {
			if passRsp.Success {
				//生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + 60*60*24, //一天
						Issuer:    "zhaoshop-api",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": time.Now().Unix() + 60*60*24,
				})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "账号密码错误！",
				})
			}
		}
	}
}

//用户注册
func RegisterUser(ctx *gin.Context) {
	registerForm := forms.RegisterForm{}
	if err := ctx.ShouldBind(&registerForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	//验证短信验证码
	rdb := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port)})
	rSms, err := rdb.Get(registerForm.Mobile).Result()
	if err != redis.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	} else {
		if registerForm.Code != rSms {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证码错误",
			})
			return
		}
	}
	//注册中心获取服务
	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		zap.S().Errorw("[Register] 查询 【新建用户失败】失败：%s", err.Error())
		HandleGrapcErrorToHttp(err, ctx)
		return
	}
	//生成token
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24, //一天
			Issuer:    "zhaoshop-api",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": time.Now().Unix() + 60*60*24,
	})
}

//获取用户详细信息
func GetUserDetail(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.Id)
	rsp, err := global.UserSrvClient.GetUserById(context.Background(), &proto.IdRequset{Id: int32(currentUser.ID)})
	if err != nil {
		HandleGrapcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":     rsp.NickName,
		"birthday": time.Unix(int64(rsp.BirthDay), 0).Format("2006-01-02"),
		"gender":   rsp.Gender,
		"mobile":   rsp.Mobile,
	})
}

func UpdateUser(ctx *gin.Context) {
	updateUserForm := forms.UpdateUserForm{}
	err := ctx.ShouldBind(&updateUserForm)
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.Id)
	//将前端传递过来的日期格式转换成int
	location, _ := time.LoadLocation("Local")
	birthDay, _ := time.ParseInLocation("2006-01-02", updateUserForm.Birthday, location)
	_, err = global.UserSrvClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
		Id:       int32(currentUser.ID),
		NickName: updateUserForm.Name,
		Gender:   updateUserForm.Gender,
		BirthDay: uint64(birthDay.Unix()),
	})
	if err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
