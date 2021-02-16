package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"

	db "main/db"
	models "main/models"
	tools "main/tools"
)

// GetWallpaper handles 100% of the current api cases.
func GetWallpaper(ctx *gin.Context) {

	url := ctx.Request.URL.Query().Get("url")
	db := db.InitDb()

	var wallpaper models.Wallpaper

	//fullpath, thumbPath := "", ""

	db.Where("url = ?", url).First(&wallpaper)
	if wallpaper.ID == 0 {

		fullpath, err := tools.DownloadFile("./download", url)

		if err != nil {
			panic(err)
		}

		thumbPath := tools.ResizeImage(fullpath)

		wallpaper := models.Wallpaper{Url: url, Path: fullpath, ThumbPath: thumbPath}

		db.Create(&wallpaper)

		ctx.JSON(200, gin.H{
			"full":  ("/img" + fullpath[strings.LastIndex(fullpath, "/"):]),
			"thumb": ("/img" + thumbPath[strings.LastIndex(thumbPath, "/"):]),
		})

	} else {
		fullpath := wallpaper.Path
		thumbPath := wallpaper.ThumbPath

		ctx.JSON(200, gin.H{
			"full":  ("/img" + fullpath[strings.LastIndex(fullpath, "/"):]),
			"thumb": ("/img" + thumbPath[strings.LastIndex(thumbPath, "/"):]),
		})
	}

}
