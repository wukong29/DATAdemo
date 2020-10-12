package controllers
import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"time"
)
type UploadController struct{
	beego.Controller
}

func (r *UploadController) Get(){
	r.TplName = "demo.html"
}

func (r *UploadController) Post() {

	f, h, _ := r.GetFile("myfile") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		r.Ctx.WriteString("后缀名不符合上传要求")
		return
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		r.Ctx.WriteString(fmt.Sprintf("%v", err))
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ))

	fileName := fmt.Sprintf("%x", hashName) + ext
	//this.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = r.SaveToFile("myfile", fpath)
	if err != nil {
		r.Ctx.WriteString(fmt.Sprintf("%v", err))
	}
	r.Ctx.WriteString("上传成功~！！！！！！！")
}