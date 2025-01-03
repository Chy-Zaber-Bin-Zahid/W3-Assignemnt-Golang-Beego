package routers

import (
	"catapi/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/catImages", &controllers.MainController{}, "get:FetchCatImages")
	beego.Router("/add-to-favourites", &controllers.MainController{}, "post:AddToFavourites")
}
