package controllers

import (
	"context"
)

func HashPassword(password string) string {

}
func VerifyPassword(userPassword string, givenPassword string)(bool, string) {

}
func Signup() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}
		validationERR := validateStruct(user)
		if validtationERR != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			retrun
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err ! nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err })
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone is already in use"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updatet_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		User.ID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details  = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if insererr !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully signed in!")

	}
}
func Login() gin.HandlerFunc {
	return func (c *gin.COntext){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Seconds)
		defer cancel()
		
		var user models.User 
		if err := c.BindJSON(&user); err!=nill {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	}

}

func ProductViewerAdmin() gin.HandlerFunc {

}
func searchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}
